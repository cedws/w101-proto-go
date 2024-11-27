// Code generated by w101-client-go. DO NOT EDIT.
package mg2

import (
	"bytes"
	"encoding/binary"
	"github.com/cedws/w101-client-go/proto"
	"unsafe"
)

type mg2Service interface {
	MG2Connect(MG2Connect)
	MG2Moved(MG2Moved)
	MG2Rewards(MG2Rewards)
}

type Mg2Service struct {
	mg2Service
}

type Mg2Client struct {
	c *proto.Client
}

func (l *Mg2Service) MG2Connect(_ MG2Connect) {}
func (l *Mg2Service) MG2Moved(_ MG2Moved)     {}
func (l *Mg2Service) MG2Rewards(_ MG2Rewards) {}

func RegisterMg2Service(r *proto.MessageRouter, s mg2Service) {
	proto.RegisterMessageHandler(r, 43, 1, s.MG2Connect)
	proto.RegisterMessageHandler(r, 43, 2, s.MG2Moved)
	proto.RegisterMessageHandler(r, 43, 3, s.MG2Rewards)
}

func NewMg2Client(c *proto.Client) Mg2Client {
	return Mg2Client{c}
}

func (c Mg2Client) MG2Connect(m *MG2Connect) error {
	return c.c.WriteMessage(43, 1, m)
}

func (c Mg2Client) MG2Moved(m *MG2Moved) error {
	return c.c.WriteMessage(43, 2, m)
}

func (c Mg2Client) MG2Rewards(m *MG2Rewards) error {
	return c.c.WriteMessage(43, 3, m)
}

type MG2Connect struct {
}

func (s *MG2Connect) Marshal() []byte {
	return []byte{}
}

func (s *MG2Connect) Unmarshal(data []byte) error {
	return nil
}

type MG2Moved struct {
}

func (s *MG2Moved) Marshal() []byte {
	return []byte{}
}

func (s *MG2Moved) Unmarshal(data []byte) error {
	return nil
}

type MG2Rewards struct {
	Score    int32
	GameName string
}

func (s *MG2Rewards) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.GameName)))
	binary.Write(b, binary.LittleEndian, s.Score)
	writeString_43(b, s.GameName)
	return b.Bytes()
}

func (s *MG2Rewards) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Score); err != nil {
		return err
	}
	if s.GameName, err = readString_43(b); err != nil {
		return err
	}
	return nil
}

func writeString_43(b *bytes.Buffer, v string) {
	binary.Write(b, binary.LittleEndian, uint16(len(v)))
	b.WriteString(v)
}

func readString_43(buf *bytes.Reader) (string, error) {
	var length uint16
	if err := binary.Read(buf, binary.LittleEndian, &length); err != nil {
		return "", err
	}
	data := make([]byte, length)
	if _, err := buf.Read(data); err != nil {
		return "", err
	}
	return *(*string)(unsafe.Pointer(&data)), nil
}