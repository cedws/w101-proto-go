// Code generated by w101-client-go. DO NOT EDIT.
package script

import (
	"bytes"
	"encoding/binary"
	"github.com/cedws/w101-client-go/proto"
	"unsafe"
)

type scriptService interface {
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

type ScriptService struct {
	scriptService
}

type ScriptClient struct {
	c *proto.Client
}

func (l *ScriptService) ADDBREAKPOINT(_ ADDBREAKPOINT)     {}
func (l *ScriptService) ADDMESSAGE(_ ADDMESSAGE)           {}
func (l *ScriptService) ADDPROCESS(_ ADDPROCESS)           {}
func (l *ScriptService) ADDWATCH(_ ADDWATCH)               {}
func (l *ScriptService) ATTACHPROCESS(_ ATTACHPROCESS)     {}
func (l *ScriptService) BREAKPROCESS(_ BREAKPROCESS)       {}
func (l *ScriptService) DELBREAKPOINT(_ DELBREAKPOINT)     {}
func (l *ScriptService) DELWATCH(_ DELWATCH)               {}
func (l *ScriptService) DETACHPROCESS(_ DETACHPROCESS)     {}
func (l *ScriptService) ENUMPROCESS(_ ENUMPROCESS)         {}
func (l *ScriptService) KILLPROCESS(_ KILLPROCESS)         {}
func (l *ScriptService) PROCESSSTATE(_ PROCESSSTATE)       {}
func (l *ScriptService) PROCESSSTATUS(_ PROCESSSTATUS)     {}
func (l *ScriptService) REMOVEPROCESS(_ REMOVEPROCESS)     {}
func (l *ScriptService) RUNPROCESS(_ RUNPROCESS)           {}
func (l *ScriptService) SENDSTATE(_ SENDSTATE)             {}
func (l *ScriptService) SETVARIABLE(_ SETVARIABLE)         {}
func (l *ScriptService) STARTPROCESS(_ STARTPROCESS)       {}
func (l *ScriptService) STEPOVERPROCESS(_ STEPOVERPROCESS) {}
func (l *ScriptService) STEPPROCESS(_ STEPPROCESS)         {}

func RegisterScriptService(r *proto.MessageRouter, s scriptService) {
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

func NewScriptClient(c *proto.Client) ScriptClient {
	return ScriptClient{c}
}

func (c ScriptClient) ADDBREAKPOINT(m *ADDBREAKPOINT) error {
	return c.c.WriteMessage(10, 1, m)
}

func (c ScriptClient) ADDMESSAGE(m *ADDMESSAGE) error {
	return c.c.WriteMessage(10, 2, m)
}

func (c ScriptClient) ADDPROCESS(m *ADDPROCESS) error {
	return c.c.WriteMessage(10, 3, m)
}

func (c ScriptClient) ADDWATCH(m *ADDWATCH) error {
	return c.c.WriteMessage(10, 4, m)
}

func (c ScriptClient) ATTACHPROCESS(m *ATTACHPROCESS) error {
	return c.c.WriteMessage(10, 5, m)
}

func (c ScriptClient) BREAKPROCESS(m *BREAKPROCESS) error {
	return c.c.WriteMessage(10, 6, m)
}

func (c ScriptClient) DELBREAKPOINT(m *DELBREAKPOINT) error {
	return c.c.WriteMessage(10, 7, m)
}

func (c ScriptClient) DELWATCH(m *DELWATCH) error {
	return c.c.WriteMessage(10, 8, m)
}

func (c ScriptClient) DETACHPROCESS(m *DETACHPROCESS) error {
	return c.c.WriteMessage(10, 9, m)
}

func (c ScriptClient) ENUMPROCESS(m *ENUMPROCESS) error {
	return c.c.WriteMessage(10, 10, m)
}

func (c ScriptClient) KILLPROCESS(m *KILLPROCESS) error {
	return c.c.WriteMessage(10, 11, m)
}

func (c ScriptClient) PROCESSSTATE(m *PROCESSSTATE) error {
	return c.c.WriteMessage(10, 12, m)
}

func (c ScriptClient) PROCESSSTATUS(m *PROCESSSTATUS) error {
	return c.c.WriteMessage(10, 13, m)
}

func (c ScriptClient) REMOVEPROCESS(m *REMOVEPROCESS) error {
	return c.c.WriteMessage(10, 14, m)
}

func (c ScriptClient) RUNPROCESS(m *RUNPROCESS) error {
	return c.c.WriteMessage(10, 15, m)
}

func (c ScriptClient) SENDSTATE(m *SENDSTATE) error {
	return c.c.WriteMessage(10, 16, m)
}

func (c ScriptClient) SETVARIABLE(m *SETVARIABLE) error {
	return c.c.WriteMessage(10, 17, m)
}

func (c ScriptClient) STARTPROCESS(m *STARTPROCESS) error {
	return c.c.WriteMessage(10, 18, m)
}

func (c ScriptClient) STEPOVERPROCESS(m *STEPOVERPROCESS) error {
	return c.c.WriteMessage(10, 19, m)
}

func (c ScriptClient) STEPPROCESS(m *STEPPROCESS) error {
	return c.c.WriteMessage(10, 20, m)
}

type ADDBREAKPOINT struct {
	File string
	Line uint32
}

func (s *ADDBREAKPOINT) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.File)))
	writeString_10(b, s.File)
	binary.Write(b, binary.LittleEndian, s.Line)
	return b.Bytes()
}

func (s *ADDBREAKPOINT) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.File, err = readString_10(b); err != nil {
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
	writeString_10(b, s.Message)
	return b.Bytes()
}

func (s *ADDMESSAGE) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Message, err = readString_10(b); err != nil {
		return err
	}
	return nil
}

type ADDPROCESS struct {
	PID    uint32
	Root   string
	Zone   string
	Name   string
	Status string
}

func (s *ADDPROCESS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 12+len(s.Root)+len(s.Zone)+len(s.Name)+len(s.Status)))
	binary.Write(b, binary.LittleEndian, s.PID)
	writeString_10(b, s.Root)
	writeString_10(b, s.Zone)
	writeString_10(b, s.Name)
	writeString_10(b, s.Status)
	return b.Bytes()
}

func (s *ADDPROCESS) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.PID); err != nil {
		return err
	}
	if s.Root, err = readString_10(b); err != nil {
		return err
	}
	if s.Zone, err = readString_10(b); err != nil {
		return err
	}
	if s.Name, err = readString_10(b); err != nil {
		return err
	}
	if s.Status, err = readString_10(b); err != nil {
		return err
	}
	return nil
}

type ADDWATCH struct {
	Variable string
}

func (s *ADDWATCH) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.Variable)))
	writeString_10(b, s.Variable)
	return b.Bytes()
}

func (s *ADDWATCH) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Variable, err = readString_10(b); err != nil {
		return err
	}
	return nil
}

type ATTACHPROCESS struct {
	PID uint32
}

func (s *ATTACHPROCESS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	binary.Write(b, binary.LittleEndian, s.PID)
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
	writeString_10(b, s.File)
	binary.Write(b, binary.LittleEndian, s.Line)
	return b.Bytes()
}

func (s *DELBREAKPOINT) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.File, err = readString_10(b); err != nil {
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
	writeString_10(b, s.Variable)
	return b.Bytes()
}

func (s *DELWATCH) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Variable, err = readString_10(b); err != nil {
		return err
	}
	return nil
}

type DETACHPROCESS struct {
	PID uint32
}

func (s *DETACHPROCESS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	binary.Write(b, binary.LittleEndian, s.PID)
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
	binary.Write(b, binary.LittleEndian, s.PID)
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
	PID   uint32
	State string
}

func (s *PROCESSSTATE) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.State)))
	binary.Write(b, binary.LittleEndian, s.PID)
	writeString_10(b, s.State)
	return b.Bytes()
}

func (s *PROCESSSTATE) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.PID); err != nil {
		return err
	}
	if s.State, err = readString_10(b); err != nil {
		return err
	}
	return nil
}

type PROCESSSTATUS struct {
	PID    uint32
	Status string
}

func (s *PROCESSSTATUS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.Status)))
	binary.Write(b, binary.LittleEndian, s.PID)
	writeString_10(b, s.Status)
	return b.Bytes()
}

func (s *PROCESSSTATUS) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.PID); err != nil {
		return err
	}
	if s.Status, err = readString_10(b); err != nil {
		return err
	}
	return nil
}

type REMOVEPROCESS struct {
	PID    uint32
	Name   string
	Status string
}

func (s *REMOVEPROCESS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 8+len(s.Name)+len(s.Status)))
	binary.Write(b, binary.LittleEndian, s.PID)
	writeString_10(b, s.Name)
	writeString_10(b, s.Status)
	return b.Bytes()
}

func (s *REMOVEPROCESS) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.PID); err != nil {
		return err
	}
	if s.Name, err = readString_10(b); err != nil {
		return err
	}
	if s.Status, err = readString_10(b); err != nil {
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
	Variable string
	Value    string
}

func (s *SETVARIABLE) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4+len(s.Variable)+len(s.Value)))
	writeString_10(b, s.Variable)
	writeString_10(b, s.Value)
	return b.Bytes()
}

func (s *SETVARIABLE) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Variable, err = readString_10(b); err != nil {
		return err
	}
	if s.Value, err = readString_10(b); err != nil {
		return err
	}
	return nil
}

type STARTPROCESS struct {
	DataRoot      string
	Zone          string
	Script        string
	StartAttached int32
	StartRunning  int32
}

func (s *STARTPROCESS) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 14+len(s.DataRoot)+len(s.Zone)+len(s.Script)))
	writeString_10(b, s.DataRoot)
	writeString_10(b, s.Zone)
	writeString_10(b, s.Script)
	binary.Write(b, binary.LittleEndian, s.StartAttached)
	binary.Write(b, binary.LittleEndian, s.StartRunning)
	return b.Bytes()
}

func (s *STARTPROCESS) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.DataRoot, err = readString_10(b); err != nil {
		return err
	}
	if s.Zone, err = readString_10(b); err != nil {
		return err
	}
	if s.Script, err = readString_10(b); err != nil {
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

func writeString_10(b *bytes.Buffer, v string) {
	binary.Write(b, binary.LittleEndian, uint16(len(v)))
	b.WriteString(v)
}

func readString_10(buf *bytes.Reader) (string, error) {
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