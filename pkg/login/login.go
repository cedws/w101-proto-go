// Code generated by w101-client-go. DO NOT EDIT.
package login

import (
	"bytes"
	"encoding/binary"
	"github.com/cedws/w101-client-go/codegen"
	"github.com/cedws/w101-client-go/proto"
)

type service interface {
	CharacterInfo(CharacterInfo)
	CharacterList(CharacterList)
	CharacterSelected(CharacterSelected)
	CreateCharacter(CreateCharacter)
	CreateCharacterResponse(CreateCharacterResponse)
	DeleteCharacter(DeleteCharacter)
	DeleteCharacterResponse(DeleteCharacterResponse)
	RequestCharacterList(RequestCharacterList)
	RequestServerList(RequestServerList)
	SelectCharacter(SelectCharacter)
	ServerList(ServerList)
	StartCharacterList(StartCharacterList)
	UserAuthen(UserAuthen)
	UserAuthenRsp(UserAuthenRsp)
	UserValidate(UserValidate)
	UserValidateRsp(UserValidateRsp)
	DisconnectLoginAfk(DisconnectLoginAfk)
	LoginNotAfk(LoginNotAfk)
	ServerShutdown(ServerShutdown)
	UserAdmitInd(UserAdmitInd)
	WebCharacterInfo(WebCharacterInfo)
	UserAuthenV2(UserAuthenV2)
	SaveCharacter(SaveCharacter)
	WebAuthen(WebAuthen)
	WebValidate(WebValidate)
	ChangeCharacterName(ChangeCharacterName)
	UserAuthenV3(UserAuthenV3)
	LoginLogCharacterCreation(LoginLogCharacterCreation)
}

func (Service) CharacterInfo(CharacterInfo)                         {}
func (Service) CharacterList(CharacterList)                         {}
func (Service) CharacterSelected(CharacterSelected)                 {}
func (Service) CreateCharacter(CreateCharacter)                     {}
func (Service) CreateCharacterResponse(CreateCharacterResponse)     {}
func (Service) DeleteCharacter(DeleteCharacter)                     {}
func (Service) DeleteCharacterResponse(DeleteCharacterResponse)     {}
func (Service) RequestCharacterList(RequestCharacterList)           {}
func (Service) RequestServerList(RequestServerList)                 {}
func (Service) SelectCharacter(SelectCharacter)                     {}
func (Service) ServerList(ServerList)                               {}
func (Service) StartCharacterList(StartCharacterList)               {}
func (Service) UserAuthen(UserAuthen)                               {}
func (Service) UserAuthenRsp(UserAuthenRsp)                         {}
func (Service) UserValidate(UserValidate)                           {}
func (Service) UserValidateRsp(UserValidateRsp)                     {}
func (Service) DisconnectLoginAfk(DisconnectLoginAfk)               {}
func (Service) LoginNotAfk(LoginNotAfk)                             {}
func (Service) ServerShutdown(ServerShutdown)                       {}
func (Service) UserAdmitInd(UserAdmitInd)                           {}
func (Service) WebCharacterInfo(WebCharacterInfo)                   {}
func (Service) UserAuthenV2(UserAuthenV2)                           {}
func (Service) SaveCharacter(SaveCharacter)                         {}
func (Service) WebAuthen(WebAuthen)                                 {}
func (Service) WebValidate(WebValidate)                             {}
func (Service) ChangeCharacterName(ChangeCharacterName)             {}
func (Service) UserAuthenV3(UserAuthenV3)                           {}
func (Service) LoginLogCharacterCreation(LoginLogCharacterCreation) {}

func RegisterService(r *proto.MessageRouter, s service) {
	proto.RegisterMessageHandler(r, 7, 1, s.CharacterInfo)
	proto.RegisterMessageHandler(r, 7, 2, s.CharacterList)
	proto.RegisterMessageHandler(r, 7, 3, s.CharacterSelected)
	proto.RegisterMessageHandler(r, 7, 4, s.CreateCharacter)
	proto.RegisterMessageHandler(r, 7, 5, s.CreateCharacterResponse)
	proto.RegisterMessageHandler(r, 7, 6, s.DeleteCharacter)
	proto.RegisterMessageHandler(r, 7, 7, s.DeleteCharacterResponse)
	proto.RegisterMessageHandler(r, 7, 8, s.RequestCharacterList)
	proto.RegisterMessageHandler(r, 7, 9, s.RequestServerList)
	proto.RegisterMessageHandler(r, 7, 10, s.SelectCharacter)
	proto.RegisterMessageHandler(r, 7, 11, s.ServerList)
	proto.RegisterMessageHandler(r, 7, 12, s.StartCharacterList)
	proto.RegisterMessageHandler(r, 7, 13, s.UserAuthen)
	proto.RegisterMessageHandler(r, 7, 14, s.UserAuthenRsp)
	proto.RegisterMessageHandler(r, 7, 15, s.UserValidate)
	proto.RegisterMessageHandler(r, 7, 16, s.UserValidateRsp)
	proto.RegisterMessageHandler(r, 7, 17, s.DisconnectLoginAfk)
	proto.RegisterMessageHandler(r, 7, 18, s.LoginNotAfk)
	proto.RegisterMessageHandler(r, 7, 19, s.ServerShutdown)
	proto.RegisterMessageHandler(r, 7, 20, s.UserAdmitInd)
	proto.RegisterMessageHandler(r, 7, 21, s.WebCharacterInfo)
	proto.RegisterMessageHandler(r, 7, 22, s.UserAuthenV2)
	proto.RegisterMessageHandler(r, 7, 23, s.SaveCharacter)
	proto.RegisterMessageHandler(r, 7, 24, s.WebAuthen)
	proto.RegisterMessageHandler(r, 7, 25, s.WebValidate)
	proto.RegisterMessageHandler(r, 7, 26, s.ChangeCharacterName)
	proto.RegisterMessageHandler(r, 7, 27, s.UserAuthenV3)
	proto.RegisterMessageHandler(r, 7, 28, s.LoginLogCharacterCreation)
}

func NewClient(c *proto.Client) Client {
	return Client{c}
}

func (c Client) CharacterInfo(m *CharacterInfo) error {
	return c.c.WriteMessage(7, 1, m)
}

func (c Client) CharacterList(m *CharacterList) error {
	return c.c.WriteMessage(7, 2, m)
}

func (c Client) CharacterSelected(m *CharacterSelected) error {
	return c.c.WriteMessage(7, 3, m)
}

func (c Client) CreateCharacter(m *CreateCharacter) error {
	return c.c.WriteMessage(7, 4, m)
}

func (c Client) CreateCharacterResponse(m *CreateCharacterResponse) error {
	return c.c.WriteMessage(7, 5, m)
}

func (c Client) DeleteCharacter(m *DeleteCharacter) error {
	return c.c.WriteMessage(7, 6, m)
}

func (c Client) DeleteCharacterResponse(m *DeleteCharacterResponse) error {
	return c.c.WriteMessage(7, 7, m)
}

func (c Client) RequestCharacterList(m *RequestCharacterList) error {
	return c.c.WriteMessage(7, 8, m)
}

func (c Client) RequestServerList(m *RequestServerList) error {
	return c.c.WriteMessage(7, 9, m)
}

func (c Client) SelectCharacter(m *SelectCharacter) error {
	return c.c.WriteMessage(7, 10, m)
}

func (c Client) ServerList(m *ServerList) error {
	return c.c.WriteMessage(7, 11, m)
}

func (c Client) StartCharacterList(m *StartCharacterList) error {
	return c.c.WriteMessage(7, 12, m)
}

func (c Client) UserAuthen(m *UserAuthen) error {
	return c.c.WriteMessage(7, 13, m)
}

func (c Client) UserAuthenRsp(m *UserAuthenRsp) error {
	return c.c.WriteMessage(7, 14, m)
}

func (c Client) UserValidate(m *UserValidate) error {
	return c.c.WriteMessage(7, 15, m)
}

func (c Client) UserValidateRsp(m *UserValidateRsp) error {
	return c.c.WriteMessage(7, 16, m)
}

func (c Client) DisconnectLoginAfk(m *DisconnectLoginAfk) error {
	return c.c.WriteMessage(7, 17, m)
}

func (c Client) LoginNotAfk(m *LoginNotAfk) error {
	return c.c.WriteMessage(7, 18, m)
}

func (c Client) ServerShutdown(m *ServerShutdown) error {
	return c.c.WriteMessage(7, 19, m)
}

func (c Client) UserAdmitInd(m *UserAdmitInd) error {
	return c.c.WriteMessage(7, 20, m)
}

func (c Client) WebCharacterInfo(m *WebCharacterInfo) error {
	return c.c.WriteMessage(7, 21, m)
}

func (c Client) UserAuthenV2(m *UserAuthenV2) error {
	return c.c.WriteMessage(7, 22, m)
}

func (c Client) SaveCharacter(m *SaveCharacter) error {
	return c.c.WriteMessage(7, 23, m)
}

func (c Client) WebAuthen(m *WebAuthen) error {
	return c.c.WriteMessage(7, 24, m)
}

func (c Client) WebValidate(m *WebValidate) error {
	return c.c.WriteMessage(7, 25, m)
}

func (c Client) ChangeCharacterName(m *ChangeCharacterName) error {
	return c.c.WriteMessage(7, 26, m)
}

func (c Client) UserAuthenV3(m *UserAuthenV3) error {
	return c.c.WriteMessage(7, 27, m)
}

func (c Client) LoginLogCharacterCreation(m *LoginLogCharacterCreation) error {
	return c.c.WriteMessage(7, 28, m)
}

type Service struct {
	service
}

type Client struct {
	c *proto.Client
}
type CharacterInfo struct {
	CharacterInfo string
}

func (s *CharacterInfo) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.CharacterInfo)))
	binary.Write(b, binary.LittleEndian, s.CharacterInfo)
	return b.Bytes()
}

func (s *CharacterInfo) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.CharacterInfo, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type CharacterList struct {
	Error uint32
}

func (s *CharacterList) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	return b.Bytes()
}

func (s *CharacterList) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Error); err != nil {
		return err
	}
	return nil
}

type CharacterSelected struct {
	ZoneName    string
	Key         string
	Location    string
	LoginServer string
	IP          string
	ZoneID      uint64
	CharID      uint64
	UserID      uint64
	UDPPort     int32
	Slot        int32
	PrepPhase   int32
	Error       int32
	TCPPort     int32
}

func (s *CharacterSelected) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 54+len(s.IP)+len(s.Key)+len(s.ZoneName)+len(s.Location)+len(s.LoginServer)))
	binary.Write(b, binary.LittleEndian, s.IP)
	binary.Write(b, binary.LittleEndian, s.Key)
	binary.Write(b, binary.LittleEndian, s.ZoneName)
	binary.Write(b, binary.LittleEndian, s.Location)
	binary.Write(b, binary.LittleEndian, s.LoginServer)
	return b.Bytes()
}

func (s *CharacterSelected) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.IP, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.TCPPort); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.UDPPort); err != nil {
		return err
	}
	if s.Key, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.UserID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.CharID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ZoneID); err != nil {
		return err
	}
	if s.ZoneName, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Location, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Slot); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.PrepPhase); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Error); err != nil {
		return err
	}
	if s.LoginServer, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type CreateCharacter struct {
	CreationInfo string
}

func (s *CreateCharacter) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.CreationInfo)))
	binary.Write(b, binary.LittleEndian, s.CreationInfo)
	return b.Bytes()
}

func (s *CreateCharacter) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.CreationInfo, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type CreateCharacterResponse struct {
	ErrorCode int32
}

func (s *CreateCharacterResponse) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	return b.Bytes()
}

func (s *CreateCharacterResponse) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.ErrorCode); err != nil {
		return err
	}
	return nil
}

type DeleteCharacter struct {
	CharID uint64
}

func (s *DeleteCharacter) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 8))
	return b.Bytes()
}

func (s *DeleteCharacter) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.CharID); err != nil {
		return err
	}
	return nil
}

type DeleteCharacterResponse struct {
	ErrorCode int32
}

func (s *DeleteCharacterResponse) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	return b.Bytes()
}

func (s *DeleteCharacterResponse) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.ErrorCode); err != nil {
		return err
	}
	return nil
}

type RequestCharacterList struct {
}

func (s *RequestCharacterList) Marshal() []byte {
	return []byte{}
}

func (s *RequestCharacterList) Unmarshal(data []byte) error {
	return nil
}

type RequestServerList struct {
}

func (s *RequestServerList) Marshal() []byte {
	return []byte{}
}

func (s *RequestServerList) Unmarshal(data []byte) error {
	return nil
}

type SelectCharacter struct {
	ServerName string
	CharID     uint64
}

func (s *SelectCharacter) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 10+len(s.ServerName)))
	binary.Write(b, binary.LittleEndian, s.ServerName)
	return b.Bytes()
}

func (s *SelectCharacter) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.CharID); err != nil {
		return err
	}
	if s.ServerName, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type ServerList struct {
}

func (s *ServerList) Marshal() []byte {
	return []byte{}
}

func (s *ServerList) Unmarshal(data []byte) error {
	return nil
}

type StartCharacterList struct {
	LoginServer             string
	PurchasedCharacterSlots int32
}

func (s *StartCharacterList) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.LoginServer)))
	binary.Write(b, binary.LittleEndian, s.LoginServer)
	return b.Bytes()
}

func (s *StartCharacterList) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.LoginServer, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.PurchasedCharacterSlots); err != nil {
		return err
	}
	return nil
}

type UserAuthen struct {
	PatchClientID string
	CRC           string
	DataRevision  string
	Revision      string
	Version       string
	Rec1          string
	MachineID     uint64
}

func (s *UserAuthen) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 20+len(s.Rec1)+len(s.Version)+len(s.Revision)+len(s.DataRevision)+len(s.CRC)+len(s.PatchClientID)))
	binary.Write(b, binary.LittleEndian, s.Rec1)
	binary.Write(b, binary.LittleEndian, s.Version)
	binary.Write(b, binary.LittleEndian, s.Revision)
	binary.Write(b, binary.LittleEndian, s.DataRevision)
	binary.Write(b, binary.LittleEndian, s.CRC)
	binary.Write(b, binary.LittleEndian, s.PatchClientID)
	return b.Bytes()
}

func (s *UserAuthen) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Rec1, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Version, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Revision, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.DataRevision, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.CRC, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	if s.PatchClientID, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type UserAuthenRsp struct {
	TimeStamp  string
	Reason     string
	Rec1       string
	UserID     uint64
	Error      int32
	PayingUser int32
	Flags      int32
}

func (s *UserAuthenRsp) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 26+len(s.Rec1)+len(s.Reason)+len(s.TimeStamp)))
	binary.Write(b, binary.LittleEndian, s.Rec1)
	binary.Write(b, binary.LittleEndian, s.Reason)
	binary.Write(b, binary.LittleEndian, s.TimeStamp)
	return b.Bytes()
}

func (s *UserAuthenRsp) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Error); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.UserID); err != nil {
		return err
	}
	if s.Rec1, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Reason, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.TimeStamp, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.PayingUser); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Flags); err != nil {
		return err
	}
	return nil
}

type UserValidate struct {
	PatchClientID string
	Locale        string
	PassKey3      string
	UserID        uint64
	MachineID     uint64
}

func (s *UserValidate) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 22+len(s.PassKey3)+len(s.Locale)+len(s.PatchClientID)))
	binary.Write(b, binary.LittleEndian, s.PassKey3)
	binary.Write(b, binary.LittleEndian, s.Locale)
	binary.Write(b, binary.LittleEndian, s.PatchClientID)
	return b.Bytes()
}

func (s *UserValidate) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.UserID); err != nil {
		return err
	}
	if s.PassKey3, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	if s.Locale, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.PatchClientID, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type UserValidateRsp struct {
	TimeStamp  string
	Reason     string
	UserID     uint64
	Error      int32
	PayingUser int32
	Flags      int32
}

func (s *UserValidateRsp) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 24+len(s.Reason)+len(s.TimeStamp)))
	binary.Write(b, binary.LittleEndian, s.Reason)
	binary.Write(b, binary.LittleEndian, s.TimeStamp)
	return b.Bytes()
}

func (s *UserValidateRsp) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Error); err != nil {
		return err
	}
	if s.Reason, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.UserID); err != nil {
		return err
	}
	if s.TimeStamp, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.PayingUser); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Flags); err != nil {
		return err
	}
	return nil
}

type DisconnectLoginAfk struct {
	Warning int8
}

func (s *DisconnectLoginAfk) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 1))
	return b.Bytes()
}

func (s *DisconnectLoginAfk) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Warning); err != nil {
		return err
	}
	return nil
}

type LoginNotAfk struct {
	BadgeNameID uint32
}

func (s *LoginNotAfk) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	return b.Bytes()
}

func (s *LoginNotAfk) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.BadgeNameID); err != nil {
		return err
	}
	return nil
}

type ServerShutdown struct {
	Message uint32
}

func (s *ServerShutdown) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	return b.Bytes()
}

func (s *ServerShutdown) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Message); err != nil {
		return err
	}
	return nil
}

type UserAdmitInd struct {
	Status          int32
	PositionInQueue uint32
}

func (s *UserAdmitInd) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 8))
	return b.Bytes()
}

func (s *UserAdmitInd) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Status); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.PositionInQueue); err != nil {
		return err
	}
	return nil
}

type WebCharacterInfo struct {
	Name   int32
	Gender int32
	School int32
}

func (s *WebCharacterInfo) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 12))
	return b.Bytes()
}

func (s *WebCharacterInfo) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Name); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Gender); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.School); err != nil {
		return err
	}
	return nil
}

type UserAuthenV2 struct {
	PatchClientID string
	Locale        string
	CRC           string
	DataRevision  string
	Revision      string
	Version       string
	Rec1          string
	MachineID     uint64
}

func (s *UserAuthenV2) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 22+len(s.Rec1)+len(s.Version)+len(s.Revision)+len(s.DataRevision)+len(s.CRC)+len(s.Locale)+len(s.PatchClientID)))
	binary.Write(b, binary.LittleEndian, s.Rec1)
	binary.Write(b, binary.LittleEndian, s.Version)
	binary.Write(b, binary.LittleEndian, s.Revision)
	binary.Write(b, binary.LittleEndian, s.DataRevision)
	binary.Write(b, binary.LittleEndian, s.CRC)
	binary.Write(b, binary.LittleEndian, s.Locale)
	binary.Write(b, binary.LittleEndian, s.PatchClientID)
	return b.Bytes()
}

func (s *UserAuthenV2) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Rec1, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Version, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Revision, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.DataRevision, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.CRC, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	if s.Locale, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.PatchClientID, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type SaveCharacter struct {
	CharID  uint64
	Success uint8
}

func (s *SaveCharacter) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 9))
	return b.Bytes()
}

func (s *SaveCharacter) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.CharID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Success); err != nil {
		return err
	}
	return nil
}

type WebAuthen struct {
	CRC          string
	DataRevision string
	Revision     string
	Version      string
	Rec1         string
	MachineID    uint64
}

func (s *WebAuthen) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 18+len(s.Rec1)+len(s.Version)+len(s.Revision)+len(s.DataRevision)+len(s.CRC)))
	binary.Write(b, binary.LittleEndian, s.Rec1)
	binary.Write(b, binary.LittleEndian, s.Version)
	binary.Write(b, binary.LittleEndian, s.Revision)
	binary.Write(b, binary.LittleEndian, s.DataRevision)
	binary.Write(b, binary.LittleEndian, s.CRC)
	return b.Bytes()
}

func (s *WebAuthen) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Rec1, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Version, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Revision, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.DataRevision, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.CRC, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	return nil
}

type WebValidate struct {
	Locale    string
	PassKey3  string
	UserID    uint64
	MachineID uint64
}

func (s *WebValidate) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 20+len(s.PassKey3)+len(s.Locale)))
	binary.Write(b, binary.LittleEndian, s.PassKey3)
	binary.Write(b, binary.LittleEndian, s.Locale)
	return b.Bytes()
}

func (s *WebValidate) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.UserID); err != nil {
		return err
	}
	if s.PassKey3, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	if s.Locale, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type ChangeCharacterName struct {
	ServerName string
	NewName    string
	CharID     uint64
}

func (s *ChangeCharacterName) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 12+len(s.NewName)+len(s.ServerName)))
	binary.Write(b, binary.LittleEndian, s.NewName)
	binary.Write(b, binary.LittleEndian, s.ServerName)
	return b.Bytes()
}

func (s *ChangeCharacterName) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.CharID); err != nil {
		return err
	}
	if s.NewName, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.ServerName, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type UserAuthenV3 struct {
	PatchClientID  string
	Locale         string
	CRC            string
	DataRevision   string
	Revision       string
	Version        string
	Rec1           string
	MachineID      uint64
	IsSteamPatcher uint32
	ConsoleType    uint8
}

func (s *UserAuthenV3) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 27+len(s.Rec1)+len(s.Version)+len(s.Revision)+len(s.DataRevision)+len(s.CRC)+len(s.Locale)+len(s.PatchClientID)))
	binary.Write(b, binary.LittleEndian, s.Rec1)
	binary.Write(b, binary.LittleEndian, s.Version)
	binary.Write(b, binary.LittleEndian, s.Revision)
	binary.Write(b, binary.LittleEndian, s.DataRevision)
	binary.Write(b, binary.LittleEndian, s.CRC)
	binary.Write(b, binary.LittleEndian, s.Locale)
	binary.Write(b, binary.LittleEndian, s.PatchClientID)
	return b.Bytes()
}

func (s *UserAuthenV3) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Rec1, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Version, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Revision, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.DataRevision, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.CRC, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	if s.Locale, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.PatchClientID, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.IsSteamPatcher); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ConsoleType); err != nil {
		return err
	}
	return nil
}

type LoginLogCharacterCreation struct {
	Stage     uint32
	Parameter uint32
}

func (s *LoginLogCharacterCreation) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 8))
	return b.Bytes()
}

func (s *LoginLogCharacterCreation) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Stage); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Parameter); err != nil {
		return err
	}
	return nil
}
