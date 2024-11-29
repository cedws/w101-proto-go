// Code generated by w101-client-go. DO NOT EDIT.
package extendedbase

import (
	"bytes"
	"encoding/binary"
	"github.com/cedws/w101-client-go/codegen"
	"github.com/cedws/w101-client-go/proto"
)

type service interface {
	CustomDict(CustomDict)
	CustomRecord(CustomRecord)
	ForceDisconnect(ForceDisconnect)
	RawRecord(RawRecord)
	RawText(RawText)
	ServerMessage(ServerMessage)
}

func (Service) CustomDict(CustomDict)           {}
func (Service) CustomRecord(CustomRecord)       {}
func (Service) ForceDisconnect(ForceDisconnect) {}
func (Service) RawRecord(RawRecord)             {}
func (Service) RawText(RawText)                 {}
func (Service) ServerMessage(ServerMessage)     {}

func RegisterService(r *proto.MessageRouter, s service) {
	proto.RegisterMessageHandler(r, 2, 1, s.CustomDict)
	proto.RegisterMessageHandler(r, 2, 2, s.CustomRecord)
	proto.RegisterMessageHandler(r, 2, 3, s.ForceDisconnect)
	proto.RegisterMessageHandler(r, 2, 4, s.RawRecord)
	proto.RegisterMessageHandler(r, 2, 5, s.RawText)
	proto.RegisterMessageHandler(r, 2, 6, s.ServerMessage)
}

func NewClient(c *proto.Client) Client {
	return Client{c}
}

func (c Client) CustomDict(m *CustomDict) error {
	return c.c.WriteMessage(2, 1, m)
}

func (c Client) CustomRecord(m *CustomRecord) error {
	return c.c.WriteMessage(2, 2, m)
}

func (c Client) ForceDisconnect(m *ForceDisconnect) error {
	return c.c.WriteMessage(2, 3, m)
}

func (c Client) RawRecord(m *RawRecord) error {
	return c.c.WriteMessage(2, 4, m)
}

func (c Client) RawText(m *RawText) error {
	return c.c.WriteMessage(2, 5, m)
}

func (c Client) ServerMessage(m *ServerMessage) error {
	return c.c.WriteMessage(2, 6, m)
}

type Service struct {
	service
}

type Client struct {
	c *proto.Client
}
type CustomDict struct {
}

func (s *CustomDict) Marshal() []byte {
	return []byte{}
}

func (s *CustomDict) Unmarshal(data []byte) error {
	return nil
}

type CustomRecord struct {
}

func (s *CustomRecord) Marshal() []byte {
	return []byte{}
}

func (s *CustomRecord) Unmarshal(data []byte) error {
	return nil
}

type ForceDisconnect struct {
	Message   string
	TimeStamp string
	Type      uint32
}

func (s *ForceDisconnect) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 8+len(s.TimeStamp)+len(s.Message)))
	binary.Write(b, binary.LittleEndian, s.TimeStamp)
	binary.Write(b, binary.LittleEndian, s.Message)
	return b.Bytes()
}

func (s *ForceDisconnect) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Type); err != nil {
		return err
	}
	if s.TimeStamp, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Message, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type RawRecord struct {
}

func (s *RawRecord) Marshal() []byte {
	return []byte{}
}

func (s *RawRecord) Unmarshal(data []byte) error {
	return nil
}

type RawText struct {
	Message string
}

func (s *RawText) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.Message)))
	binary.Write(b, binary.LittleEndian, s.Message)
	return b.Bytes()
}

func (s *RawText) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Message, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type ServerMessage struct {
	Message string
	Modal   uint8
}

func (s *ServerMessage) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 3+len(s.Message)))
	binary.Write(b, binary.LittleEndian, s.Message)
	return b.Bytes()
}

func (s *ServerMessage) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Modal); err != nil {
		return err
	}
	if s.Message, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}
