// Code generated by w101-client-go. DO NOT EDIT.
package mg4

import (
	"bytes"
	"encoding/binary"
	"github.com/cedws/w101-client-go/codegen"
	"github.com/cedws/w101-client-go/proto"
)

type service interface {
	MG4Connect(MG4Connect)
	MG4Moved(MG4Moved)
	MG4Rewards(MG4Rewards)
}

func (Service) MG4Connect(MG4Connect) {}
func (Service) MG4Moved(MG4Moved)     {}
func (Service) MG4Rewards(MG4Rewards) {}

func RegisterService(r *proto.MessageRouter, s service) {
	proto.RegisterMessageHandler(r, 45, 1, s.MG4Connect)
	proto.RegisterMessageHandler(r, 45, 2, s.MG4Moved)
	proto.RegisterMessageHandler(r, 45, 3, s.MG4Rewards)
}

func NewClient(c *proto.Client) Client {
	return Client{c}
}

func (c Client) MG4Connect(m *MG4Connect) error {
	return c.c.WriteMessage(45, 1, m)
}

func (c Client) MG4Moved(m *MG4Moved) error {
	return c.c.WriteMessage(45, 2, m)
}

func (c Client) MG4Rewards(m *MG4Rewards) error {
	return c.c.WriteMessage(45, 3, m)
}

type Service struct {
	service
}

type Client struct {
	c *proto.Client
}
type MG4Connect struct {
}

func (s *MG4Connect) Marshal() []byte {
	return []byte{}
}

func (s *MG4Connect) Unmarshal(data []byte) error {
	return nil
}

type MG4Moved struct {
}

func (s *MG4Moved) Marshal() []byte {
	return []byte{}
}

func (s *MG4Moved) Unmarshal(data []byte) error {
	return nil
}

type MG4Rewards struct {
	GameName string
	Score    int32
}

func (s *MG4Rewards) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.GameName)))
	binary.Write(b, binary.LittleEndian, s.GameName)
	return b.Bytes()
}

func (s *MG4Rewards) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Score); err != nil {
		return err
	}
	if s.GameName, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}
