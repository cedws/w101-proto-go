// Code generated by w101-client-go. DO NOT EDIT.
package mg1

import (
	"bytes"
	"encoding/binary"
	"github.com/cedws/w101-client-go/codegen"
	"github.com/cedws/w101-client-go/proto"
)

type service interface {
	MG1Connect(MG1Connect)
	MG1Moved(MG1Moved)
	MG1Rewards(MG1Rewards)
}

func (Service) MG1Connect(MG1Connect) {}
func (Service) MG1Moved(MG1Moved)     {}
func (Service) MG1Rewards(MG1Rewards) {}

func RegisterService(r *proto.MessageRouter, s service) {
	proto.RegisterMessageHandler(r, 42, 1, s.MG1Connect)
	proto.RegisterMessageHandler(r, 42, 2, s.MG1Moved)
	proto.RegisterMessageHandler(r, 42, 3, s.MG1Rewards)
}

func NewClient(c *proto.Client) Client {
	return Client{c}
}

func (c Client) MG1Connect(m *MG1Connect) error {
	return c.c.WriteMessage(42, 1, m)
}

func (c Client) MG1Moved(m *MG1Moved) error {
	return c.c.WriteMessage(42, 2, m)
}

func (c Client) MG1Rewards(m *MG1Rewards) error {
	return c.c.WriteMessage(42, 3, m)
}

type Service struct {
	service
}

type Client struct {
	c *proto.Client
}
type MG1Connect struct {
}

func (s *MG1Connect) Marshal() []byte {
	return []byte{}
}

func (s *MG1Connect) Unmarshal(data []byte) error {
	return nil
}

type MG1Moved struct {
}

func (s *MG1Moved) Marshal() []byte {
	return []byte{}
}

func (s *MG1Moved) Unmarshal(data []byte) error {
	return nil
}

type MG1Rewards struct {
	GameName string
	Score    int32
}

func (s *MG1Rewards) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.GameName)))
	binary.Write(b, binary.LittleEndian, s.GameName)
	return b.Bytes()
}

func (s *MG1Rewards) Unmarshal(data []byte) error {
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
