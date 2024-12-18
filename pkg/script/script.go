// Code generated by w101-client-go. DO NOT EDIT.
package script

import (
	"bytes"
	"encoding/binary"
	"github.com/cedws/w101-client-go/codegen"
	"github.com/cedws/w101-client-go/proto"
)

type service interface {
	ADDBREAKPOINT(ADDBREAKPOINT)
	ADDMESSAGE(ADDMESSAGE)
	ADDPROCESS(ADDPROCESS)
	ADDWATCH(ADDWATCH)
	ATTACHPROCESS(ATTACHPROCESS)
	BREAKPROCESS(BREAKPROCESS)
	DELBREAKPOINT(DELBREAKPOINT)
	DELWATCH(DELWATCH)
	DETACHPROCESS(DETACHPROCESS)
	ENUMPROCESS(ENUMPROCESS)
	KILLPROCESS(KILLPROCESS)
	PROCESSSTATE(PROCESSSTATE)
	PROCESSSTATUS(PROCESSSTATUS)
	REMOVEPROCESS(REMOVEPROCESS)
	RUNPROCESS(RUNPROCESS)
	SENDSTATE(SENDSTATE)
	SETVARIABLE(SETVARIABLE)
	STARTPROCESS(STARTPROCESS)
	STEPOVERPROCESS(STEPOVERPROCESS)
	STEPPROCESS(STEPPROCESS)
}

func (Service) ADDBREAKPOINT(ADDBREAKPOINT)     {}
func (Service) ADDMESSAGE(ADDMESSAGE)           {}
func (Service) ADDPROCESS(ADDPROCESS)           {}
func (Service) ADDWATCH(ADDWATCH)               {}
func (Service) ATTACHPROCESS(ATTACHPROCESS)     {}
func (Service) BREAKPROCESS(BREAKPROCESS)       {}
func (Service) DELBREAKPOINT(DELBREAKPOINT)     {}
func (Service) DELWATCH(DELWATCH)               {}
func (Service) DETACHPROCESS(DETACHPROCESS)     {}
func (Service) ENUMPROCESS(ENUMPROCESS)         {}
func (Service) KILLPROCESS(KILLPROCESS)         {}
func (Service) PROCESSSTATE(PROCESSSTATE)       {}
func (Service) PROCESSSTATUS(PROCESSSTATUS)     {}
func (Service) REMOVEPROCESS(REMOVEPROCESS)     {}
func (Service) RUNPROCESS(RUNPROCESS)           {}
func (Service) SENDSTATE(SENDSTATE)             {}
func (Service) SETVARIABLE(SETVARIABLE)         {}
func (Service) STARTPROCESS(STARTPROCESS)       {}
func (Service) STEPOVERPROCESS(STEPOVERPROCESS) {}
func (Service) STEPPROCESS(STEPPROCESS)         {}

func RegisterService(r *proto.MessageRouter, s service) {
	proto.RegisterMessageHandler(r, 10, 1, s.ADDBREAKPOINT)
	proto.RegisterMessageHandler(r, 10, 2, s.ADDMESSAGE)
	proto.RegisterMessageHandler(r, 10, 3, s.ADDPROCESS)
	proto.RegisterMessageHandler(r, 10, 4, s.ADDWATCH)
	proto.RegisterMessageHandler(r, 10, 5, s.ATTACHPROCESS)
	proto.RegisterMessageHandler(r, 10, 6, s.BREAKPROCESS)
	proto.RegisterMessageHandler(r, 10, 7, s.DELBREAKPOINT)
	proto.RegisterMessageHandler(r, 10, 8, s.DELWATCH)
	proto.RegisterMessageHandler(r, 10, 9, s.DETACHPROCESS)
	proto.RegisterMessageHandler(r, 10, 10, s.ENUMPROCESS)
	proto.RegisterMessageHandler(r, 10, 11, s.KILLPROCESS)
	proto.RegisterMessageHandler(r, 10, 12, s.PROCESSSTATE)
	proto.RegisterMessageHandler(r, 10, 13, s.PROCESSSTATUS)
	proto.RegisterMessageHandler(r, 10, 14, s.REMOVEPROCESS)
	proto.RegisterMessageHandler(r, 10, 15, s.RUNPROCESS)
	proto.RegisterMessageHandler(r, 10, 16, s.SENDSTATE)
	proto.RegisterMessageHandler(r, 10, 17, s.SETVARIABLE)
	proto.RegisterMessageHandler(r, 10, 18, s.STARTPROCESS)
	proto.RegisterMessageHandler(r, 10, 19, s.STEPOVERPROCESS)
	proto.RegisterMessageHandler(r, 10, 20, s.STEPPROCESS)
}

func NewClient(c *proto.Client) Client {
	return Client{c}
}

func (c Client) ADDBREAKPOINT(m *ADDBREAKPOINT) error {
	return c.c.WriteMessage(10, 1, m)
}

func (c Client) ADDMESSAGE(m *ADDMESSAGE) error {
	return c.c.WriteMessage(10, 2, m)
}

func (c Client) ADDPROCESS(m *ADDPROCESS) error {
	return c.c.WriteMessage(10, 3, m)
}

func (c Client) ADDWATCH(m *ADDWATCH) error {
	return c.c.WriteMessage(10, 4, m)
}

func (c Client) ATTACHPROCESS(m *ATTACHPROCESS) error {
	return c.c.WriteMessage(10, 5, m)
}

func (c Client) BREAKPROCESS(m *BREAKPROCESS) error {
	return c.c.WriteMessage(10, 6, m)
}

func (c Client) DELBREAKPOINT(m *DELBREAKPOINT) error {
	return c.c.WriteMessage(10, 7, m)
}

func (c Client) DELWATCH(m *DELWATCH) error {
	return c.c.WriteMessage(10, 8, m)
}

func (c Client) DETACHPROCESS(m *DETACHPROCESS) error {
	return c.c.WriteMessage(10, 9, m)
}

func (c Client) ENUMPROCESS(m *ENUMPROCESS) error {
	return c.c.WriteMessage(10, 10, m)
}

func (c Client) KILLPROCESS(m *KILLPROCESS) error {
	return c.c.WriteMessage(10, 11, m)
}

func (c Client) PROCESSSTATE(m *PROCESSSTATE) error {
	return c.c.WriteMessage(10, 12, m)
}

func (c Client) PROCESSSTATUS(m *PROCESSSTATUS) error {
	return c.c.WriteMessage(10, 13, m)
}

func (c Client) REMOVEPROCESS(m *REMOVEPROCESS) error {
	return c.c.WriteMessage(10, 14, m)
}

func (c Client) RUNPROCESS(m *RUNPROCESS) error {
	return c.c.WriteMessage(10, 15, m)
}

func (c Client) SENDSTATE(m *SENDSTATE) error {
	return c.c.WriteMessage(10, 16, m)
}

func (c Client) SETVARIABLE(m *SETVARIABLE) error {
	return c.c.WriteMessage(10, 17, m)
}

func (c Client) STARTPROCESS(m *STARTPROCESS) error {
	return c.c.WriteMessage(10, 18, m)
}

func (c Client) STEPOVERPROCESS(m *STEPOVERPROCESS) error {
	return c.c.WriteMessage(10, 19, m)
}

func (c Client) STEPPROCESS(m *STEPPROCESS) error {
	return c.c.WriteMessage(10, 20, m)
}

type Service struct {
	service
}

type Client struct {
	c *proto.Client
}
type ADDBREAKPOINT struct {
	File string
	Line uint32
}

func (s *ADDBREAKPOINT) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.File)))
	binary.Write(b, binary.LittleEndian, s.File)
	return b.Bytes()
}

func (s *ADDBREAKPOINT) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.File, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Line); err != nil {
		return err
	}
	return nil
}

type ADDMESSAGE struct {
	Message string
}

func (s *ADDMESSAGE) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.Message)))
	binary.Write(b, binary.LittleEndian, s.Message)
	return b.Bytes()
}

func (s *ADDMESSAGE) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Message, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type ADDPROCESS struct {
	Status string
	Name   string
	Zone   string
	Root   string
	PID    uint32
}

func (s *ADDPROCESS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 12+len(s.Root)+len(s.Zone)+len(s.Name)+len(s.Status)))
	binary.Write(b, binary.LittleEndian, s.Root)
	binary.Write(b, binary.LittleEndian, s.Zone)
	binary.Write(b, binary.LittleEndian, s.Name)
	binary.Write(b, binary.LittleEndian, s.Status)
	return b.Bytes()
}

func (s *ADDPROCESS) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.PID); err != nil {
		return err
	}
	if s.Root, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Zone, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Name, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Status, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type ADDWATCH struct {
	Variable string
}

func (s *ADDWATCH) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.Variable)))
	binary.Write(b, binary.LittleEndian, s.Variable)
	return b.Bytes()
}

func (s *ADDWATCH) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Variable, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type ATTACHPROCESS struct {
	PID uint32
}

func (s *ATTACHPROCESS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	return b.Bytes()
}

func (s *ATTACHPROCESS) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.PID); err != nil {
		return err
	}
	return nil
}

type BREAKPROCESS struct {
}

func (s *BREAKPROCESS) Marshal() []byte {
	return []byte{}
}

func (s *BREAKPROCESS) Unmarshal(data []byte) error {
	return nil
}

type DELBREAKPOINT struct {
	File string
	Line uint32
}

func (s *DELBREAKPOINT) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.File)))
	binary.Write(b, binary.LittleEndian, s.File)
	return b.Bytes()
}

func (s *DELBREAKPOINT) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.File, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Line); err != nil {
		return err
	}
	return nil
}

type DELWATCH struct {
	Variable string
}

func (s *DELWATCH) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.Variable)))
	binary.Write(b, binary.LittleEndian, s.Variable)
	return b.Bytes()
}

func (s *DELWATCH) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Variable, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type DETACHPROCESS struct {
	PID uint32
}

func (s *DETACHPROCESS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	return b.Bytes()
}

func (s *DETACHPROCESS) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.PID); err != nil {
		return err
	}
	return nil
}

type ENUMPROCESS struct {
}

func (s *ENUMPROCESS) Marshal() []byte {
	return []byte{}
}

func (s *ENUMPROCESS) Unmarshal(data []byte) error {
	return nil
}

type KILLPROCESS struct {
	PID uint32
}

func (s *KILLPROCESS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	return b.Bytes()
}

func (s *KILLPROCESS) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.PID); err != nil {
		return err
	}
	return nil
}

type PROCESSSTATE struct {
	State string
	PID   uint32
}

func (s *PROCESSSTATE) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.State)))
	binary.Write(b, binary.LittleEndian, s.State)
	return b.Bytes()
}

func (s *PROCESSSTATE) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.PID); err != nil {
		return err
	}
	if s.State, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type PROCESSSTATUS struct {
	Status string
	PID    uint32
}

func (s *PROCESSSTATUS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.Status)))
	binary.Write(b, binary.LittleEndian, s.Status)
	return b.Bytes()
}

func (s *PROCESSSTATUS) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.PID); err != nil {
		return err
	}
	if s.Status, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type REMOVEPROCESS struct {
	Status string
	Name   string
	PID    uint32
}

func (s *REMOVEPROCESS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 8+len(s.Name)+len(s.Status)))
	binary.Write(b, binary.LittleEndian, s.Name)
	binary.Write(b, binary.LittleEndian, s.Status)
	return b.Bytes()
}

func (s *REMOVEPROCESS) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.PID); err != nil {
		return err
	}
	if s.Name, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Status, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type RUNPROCESS struct {
}

func (s *RUNPROCESS) Marshal() []byte {
	return []byte{}
}

func (s *RUNPROCESS) Unmarshal(data []byte) error {
	return nil
}

type SENDSTATE struct {
}

func (s *SENDSTATE) Marshal() []byte {
	return []byte{}
}

func (s *SENDSTATE) Unmarshal(data []byte) error {
	return nil
}

type SETVARIABLE struct {
	Value    string
	Variable string
}

func (s *SETVARIABLE) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4+len(s.Variable)+len(s.Value)))
	binary.Write(b, binary.LittleEndian, s.Variable)
	binary.Write(b, binary.LittleEndian, s.Value)
	return b.Bytes()
}

func (s *SETVARIABLE) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Variable, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Value, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type STARTPROCESS struct {
	Script        string
	Zone          string
	DataRoot      string
	StartAttached int32
	StartRunning  int32
}

func (s *STARTPROCESS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 14+len(s.DataRoot)+len(s.Zone)+len(s.Script)))
	binary.Write(b, binary.LittleEndian, s.DataRoot)
	binary.Write(b, binary.LittleEndian, s.Zone)
	binary.Write(b, binary.LittleEndian, s.Script)
	return b.Bytes()
}

func (s *STARTPROCESS) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.DataRoot, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Zone, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Script, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.StartAttached); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.StartRunning); err != nil {
		return err
	}
	return nil
}

type STEPOVERPROCESS struct {
}

func (s *STEPOVERPROCESS) Marshal() []byte {
	return []byte{}
}

func (s *STEPOVERPROCESS) Unmarshal(data []byte) error {
	return nil
}

type STEPPROCESS struct {
}

func (s *STEPPROCESS) Marshal() []byte {
	return []byte{}
}

func (s *STEPPROCESS) Unmarshal(data []byte) error {
	return nil
}
