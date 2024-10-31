package main

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cedws/w101-client-go/proto"
	"github.com/cedws/w101-proto-go/patch"
)

const patchServer = "patch.us.wizard101.com:12500"

type patchHandler struct {
	patch.PatchService
	fileListCh chan patch.LatestFileListV2
}

func (p patchHandler) LatestFileListV2(m patch.LatestFileListV2) {
	p.fileListCh <- m
}

func main() {
	Patch(context.Background())
}

func Patch(ctx context.Context) error {
	fileListCh := make(chan patch.LatestFileListV2)

	r := proto.NewMessageRouter()
	patch.RegisterPatchService(r, &patchHandler{fileListCh: fileListCh})

	client, err := proto.Dial(ctx, patchServer, r)
	if err != nil {
		return err
	}
	defer client.Close()

	c := patch.NewPatchClient(client)
	if err := c.LatestFileListV2(&patch.LatestFileListV2{}); err != nil {
		return err
	}

	select {
	case fileList := <-fileListCh:
		return downloadFileList(ctx, fileList)
	case <-ctx.Done():
		return ctx.Err()
	}
}

func downloadFileList(ctx context.Context, fileList patch.LatestFileListV2) error {
	wadURL, err := url.JoinPath(fileList.URLPrefix, "Data/GameData/Root.wad")
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, wadURL, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return resp.Body.Close()
}
