// Code generated by w101-client-go. DO NOT EDIT.
package login

import (
	"bytes"
	"encoding/binary"
	"github.com/cedws/w101-client-go/proto"
	"unsafe"
)

type patchService interface {
	LatestFileList(LatestFileList)
	LatestFileListV2(LatestFileListV2)
	NextVersion(NextVersion)
}

type PatchService struct {
	patchService
}

type PatchClient struct {
	c *proto.Client
}

func (l *PatchService) LatestFileList(_ LatestFileList)     {}
func (l *PatchService) LatestFileListV2(_ LatestFileListV2) {}
func (l *PatchService) NextVersion(_ NextVersion)           {}

func RegisterPatchService(r *proto.MessageRouter, s patchService) {
	proto.RegisterMessageHandler(r, 8, 1, s.LatestFileList)
	proto.RegisterMessageHandler(r, 8, 2, s.LatestFileListV2)
	proto.RegisterMessageHandler(r, 8, 3, s.NextVersion)
}

func NewPatchClient(c *proto.Client) PatchClient {
	return PatchClient{c}
}

func (c PatchClient) LatestFileList(m *LatestFileList) error {
	return c.c.WriteMessage(8, 1, m)
}

func (c PatchClient) LatestFileListV2(m *LatestFileListV2) error {
	return c.c.WriteMessage(8, 2, m)
}

func (c PatchClient) NextVersion(m *NextVersion) error {
	return c.c.WriteMessage(8, 3, m)
}

type LatestFileList struct {
	LatestVersion uint32
	ListFileName  string
	ListFileType  uint32
	ListFileTime  uint32
	ListFileSize  uint32
	ListFileCRC   uint32
	ListFileURL   string
	URLPrefix     string
	URLSuffix     string
}

func (s *LatestFileList) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	binary.Write(&b, binary.LittleEndian, s.LatestVersion)
	writeString_8(&b, s.ListFileName)
	binary.Write(&b, binary.LittleEndian, s.ListFileType)
	binary.Write(&b, binary.LittleEndian, s.ListFileTime)
	binary.Write(&b, binary.LittleEndian, s.ListFileSize)
	binary.Write(&b, binary.LittleEndian, s.ListFileCRC)
	writeString_8(&b, s.ListFileURL)
	writeString_8(&b, s.URLPrefix)
	writeString_8(&b, s.URLSuffix)
	return b.Bytes(), nil
}

func (s *LatestFileList) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.LatestVersion); err != nil {
		return err
	}
	if s.ListFileName, err = readString_8(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ListFileType); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ListFileTime); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ListFileSize); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ListFileCRC); err != nil {
		return err
	}
	if s.ListFileURL, err = readString_8(b); err != nil {
		return err
	}
	if s.URLPrefix, err = readString_8(b); err != nil {
		return err
	}
	if s.URLSuffix, err = readString_8(b); err != nil {
		return err
	}
	return nil
}

type LatestFileListV2 struct {
	LatestVersion uint32
	ListFileName  string
	ListFileType  uint32
	ListFileTime  uint32
	ListFileSize  uint32
	ListFileCRC   uint32
	ListFileURL   string
	URLPrefix     string
	URLSuffix     string
	Locale        string
}

func (s *LatestFileListV2) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	binary.Write(&b, binary.LittleEndian, s.LatestVersion)
	writeString_8(&b, s.ListFileName)
	binary.Write(&b, binary.LittleEndian, s.ListFileType)
	binary.Write(&b, binary.LittleEndian, s.ListFileTime)
	binary.Write(&b, binary.LittleEndian, s.ListFileSize)
	binary.Write(&b, binary.LittleEndian, s.ListFileCRC)
	writeString_8(&b, s.ListFileURL)
	writeString_8(&b, s.URLPrefix)
	writeString_8(&b, s.URLSuffix)
	writeString_8(&b, s.Locale)
	return b.Bytes(), nil
}

func (s *LatestFileListV2) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.LatestVersion); err != nil {
		return err
	}
	if s.ListFileName, err = readString_8(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ListFileType); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ListFileTime); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ListFileSize); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ListFileCRC); err != nil {
		return err
	}
	if s.ListFileURL, err = readString_8(b); err != nil {
		return err
	}
	if s.URLPrefix, err = readString_8(b); err != nil {
		return err
	}
	if s.URLSuffix, err = readString_8(b); err != nil {
		return err
	}
	if s.Locale, err = readString_8(b); err != nil {
		return err
	}
	return nil
}

type NextVersion struct {
	PkgName   string
	Version   int32
	URLPrefix string
	FileName  string
	FileType  int32
}

func (s *NextVersion) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	writeString_8(&b, s.PkgName)
	binary.Write(&b, binary.LittleEndian, s.Version)
	writeString_8(&b, s.URLPrefix)
	writeString_8(&b, s.FileName)
	binary.Write(&b, binary.LittleEndian, s.FileType)
	return b.Bytes(), nil
}

func (s *NextVersion) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.PkgName, err = readString_8(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Version); err != nil {
		return err
	}
	if s.URLPrefix, err = readString_8(b); err != nil {
		return err
	}
	if s.FileName, err = readString_8(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.FileType); err != nil {
		return err
	}
	return nil
}

func writeString_8(b *bytes.Buffer, v string) {
	binary.Write(b, binary.LittleEndian, uint16(len(v)))
	b.WriteString(v)
}

func readString_8(buf *bytes.Reader) (string, error) {
	var length uint16
	if err := binary.Read(buf, binary.LittleEndian, &length); err != nil {
		return ``, err
	}
	data := make([]byte, length)
	if _, err := buf.Read(data); err != nil {
		return ``, err
	}
	return *(*string)(unsafe.Pointer(&data)), nil
}
