package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"slices"
	"time"

	"github.com/cedws/w101-client-go/proto"
	"github.com/cedws/w101-client-go/wad"
	"github.com/cedws/w101-proto-go/pkg/patch"
)

const patchServer = "patch.us.wizard101.com:12500"

var messageFiles = []string{
	"GameMessages.xml",
	"PatchMessages.xml",
	"LoginMessages.xml",
}

type patchHandler struct {
	patch.PatchService
	fileListCh chan patch.LatestFileListV2
}

func (p patchHandler) LatestFileListV2(m patch.LatestFileListV2) {
	p.fileListCh <- m
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := downloadMessageFiles(ctx); err != nil {
		log.Fatal(err)
	}

	slog.Info("Downloaded message files")
}

func downloadMessageFiles(ctx context.Context) error {
	if err := downloadRootWAD(ctx); err != nil {
		return err
	}

	wadArchive, err := wad.Open("Root.wad")
	if err != nil {
		return err
	}
	defer wadArchive.Close()

	return extractMessages(wadArchive)
}

func extractMessages(wadArchive *wad.Archive) error {
	for entry := range wadArchive.Entries() {
		if !slices.Contains(messageFiles, entry.Path) {
			continue
		}

		slog.Info("Extracting message file", "path", entry.Path)

		entryReader, err := wadArchive.Entry(entry)
		if err != nil {
			return err
		}

		msgFile, err := os.Create(entry.Path)
		if err != nil {
			return err
		}
		defer msgFile.Close()

		if _, err := io.Copy(msgFile, entryReader); err != nil {
			return err
		}
	}

	return nil
}

func downloadRootWAD(ctx context.Context) error {
	wadReader, err := latestRootWAD(ctx)
	if err != nil {
		return err
	}
	defer wadReader.Close()

	wadFile, err := os.Create("Root.wad")
	if err != nil {
		return err
	}
	defer wadFile.Close()

	slog.Info("Downloading Root.wad")

	if _, err := io.Copy(wadFile, wadReader); err != nil {
		return err
	}

	return nil
}

func latestRootWAD(ctx context.Context) (io.ReadCloser, error) {
	fileListCh := make(chan patch.LatestFileListV2)

	r := proto.NewMessageRouter()
	patch.RegisterPatchService(r, &patchHandler{fileListCh: fileListCh})

	client, err := proto.Dial(ctx, patchServer, r)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	slog.Info("Connected to patch server", "server", patchServer)

	c := patch.NewPatchClient(client)
	if err := c.LatestFileListV2(&patch.LatestFileListV2{}); err != nil {
		return nil, err
	}

	select {
	case fileList := <-fileListCh:
		slog.Info("Received latest file list", "url", fileList.ListFileURL)
		return requestRootWAD(ctx, fileList)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func requestRootWAD(ctx context.Context, fileList patch.LatestFileListV2) (io.ReadCloser, error) {
	wadURL, err := url.JoinPath(fileList.URLPrefix, "Data/GameData/Root.wad")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, wadURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp.Body, err
}
