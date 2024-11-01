// Code generated by w101-client-go. DO NOT EDIT.
package login

import (
	"bytes"
	"encoding/binary"
	"github.com/cedws/w101-client-go/proto"
	"unsafe"
)

type loginService interface {
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

type LoginService struct {
	loginService
}

type LoginClient struct {
	c *proto.Client
}

func (l *LoginService) CharacterInfo(_ CharacterInfo)                         {}
func (l *LoginService) CharacterList(_ CharacterList)                         {}
func (l *LoginService) CharacterSelected(_ CharacterSelected)                 {}
func (l *LoginService) CreateCharacter(_ CreateCharacter)                     {}
func (l *LoginService) CreateCharacterResponse(_ CreateCharacterResponse)     {}
func (l *LoginService) DeleteCharacter(_ DeleteCharacter)                     {}
func (l *LoginService) DeleteCharacterResponse(_ DeleteCharacterResponse)     {}
func (l *LoginService) RequestCharacterList(_ RequestCharacterList)           {}
func (l *LoginService) RequestServerList(_ RequestServerList)                 {}
func (l *LoginService) SelectCharacter(_ SelectCharacter)                     {}
func (l *LoginService) ServerList(_ ServerList)                               {}
func (l *LoginService) StartCharacterList(_ StartCharacterList)               {}
func (l *LoginService) UserAuthen(_ UserAuthen)                               {}
func (l *LoginService) UserAuthenRsp(_ UserAuthenRsp)                         {}
func (l *LoginService) UserValidate(_ UserValidate)                           {}
func (l *LoginService) UserValidateRsp(_ UserValidateRsp)                     {}
func (l *LoginService) DisconnectLoginAfk(_ DisconnectLoginAfk)               {}
func (l *LoginService) LoginNotAfk(_ LoginNotAfk)                             {}
func (l *LoginService) ServerShutdown(_ ServerShutdown)                       {}
func (l *LoginService) UserAdmitInd(_ UserAdmitInd)                           {}
func (l *LoginService) WebCharacterInfo(_ WebCharacterInfo)                   {}
func (l *LoginService) UserAuthenV2(_ UserAuthenV2)                           {}
func (l *LoginService) SaveCharacter(_ SaveCharacter)                         {}
func (l *LoginService) WebAuthen(_ WebAuthen)                                 {}
func (l *LoginService) WebValidate(_ WebValidate)                             {}
func (l *LoginService) ChangeCharacterName(_ ChangeCharacterName)             {}
func (l *LoginService) UserAuthenV3(_ UserAuthenV3)                           {}
func (l *LoginService) LoginLogCharacterCreation(_ LoginLogCharacterCreation) {}

func RegisterLoginService(r *proto.MessageRouter, s loginService) {
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

func NewLoginClient(c *proto.Client) LoginClient {
	return LoginClient{c}
}

func (c LoginClient) CharacterInfo(m *CharacterInfo) error {
	return c.c.WriteMessage(7, 1, m)
}

func (c LoginClient) CharacterList(m *CharacterList) error {
	return c.c.WriteMessage(7, 2, m)
}

func (c LoginClient) CharacterSelected(m *CharacterSelected) error {
	return c.c.WriteMessage(7, 3, m)
}

func (c LoginClient) CreateCharacter(m *CreateCharacter) error {
	return c.c.WriteMessage(7, 4, m)
}

func (c LoginClient) CreateCharacterResponse(m *CreateCharacterResponse) error {
	return c.c.WriteMessage(7, 5, m)
}

func (c LoginClient) DeleteCharacter(m *DeleteCharacter) error {
	return c.c.WriteMessage(7, 6, m)
}

func (c LoginClient) DeleteCharacterResponse(m *DeleteCharacterResponse) error {
	return c.c.WriteMessage(7, 7, m)
}

func (c LoginClient) RequestCharacterList(m *RequestCharacterList) error {
	return c.c.WriteMessage(7, 8, m)
}

func (c LoginClient) RequestServerList(m *RequestServerList) error {
	return c.c.WriteMessage(7, 9, m)
}

func (c LoginClient) SelectCharacter(m *SelectCharacter) error {
	return c.c.WriteMessage(7, 10, m)
}

func (c LoginClient) ServerList(m *ServerList) error {
	return c.c.WriteMessage(7, 11, m)
}

func (c LoginClient) StartCharacterList(m *StartCharacterList) error {
	return c.c.WriteMessage(7, 12, m)
}

func (c LoginClient) UserAuthen(m *UserAuthen) error {
	return c.c.WriteMessage(7, 13, m)
}

func (c LoginClient) UserAuthenRsp(m *UserAuthenRsp) error {
	return c.c.WriteMessage(7, 14, m)
}

func (c LoginClient) UserValidate(m *UserValidate) error {
	return c.c.WriteMessage(7, 15, m)
}

func (c LoginClient) UserValidateRsp(m *UserValidateRsp) error {
	return c.c.WriteMessage(7, 16, m)
}

func (c LoginClient) DisconnectLoginAfk(m *DisconnectLoginAfk) error {
	return c.c.WriteMessage(7, 17, m)
}

func (c LoginClient) LoginNotAfk(m *LoginNotAfk) error {
	return c.c.WriteMessage(7, 18, m)
}

func (c LoginClient) ServerShutdown(m *ServerShutdown) error {
	return c.c.WriteMessage(7, 19, m)
}

func (c LoginClient) UserAdmitInd(m *UserAdmitInd) error {
	return c.c.WriteMessage(7, 20, m)
}

func (c LoginClient) WebCharacterInfo(m *WebCharacterInfo) error {
	return c.c.WriteMessage(7, 21, m)
}

func (c LoginClient) UserAuthenV2(m *UserAuthenV2) error {
	return c.c.WriteMessage(7, 22, m)
}

func (c LoginClient) SaveCharacter(m *SaveCharacter) error {
	return c.c.WriteMessage(7, 23, m)
}

func (c LoginClient) WebAuthen(m *WebAuthen) error {
	return c.c.WriteMessage(7, 24, m)
}

func (c LoginClient) WebValidate(m *WebValidate) error {
	return c.c.WriteMessage(7, 25, m)
}

func (c LoginClient) ChangeCharacterName(m *ChangeCharacterName) error {
	return c.c.WriteMessage(7, 26, m)
}

func (c LoginClient) UserAuthenV3(m *UserAuthenV3) error {
	return c.c.WriteMessage(7, 27, m)
}

func (c LoginClient) LoginLogCharacterCreation(m *LoginLogCharacterCreation) error {
	return c.c.WriteMessage(7, 28, m)
}

type CharacterInfo struct {
	CharacterInfo string
}

func (s *CharacterInfo) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.CharacterInfo)))
	writeString_7(b, s.CharacterInfo)
	return b.Bytes(), nil
}

func (s *CharacterInfo) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.CharacterInfo, err = readString_7(b); err != nil {
		return err
	}
	return nil
}

type CharacterList struct {
	Error uint32
}

func (s *CharacterList) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	binary.Write(b, binary.LittleEndian, s.Error)
	return b.Bytes(), nil
}

func (s *CharacterList) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Error); err != nil {
		return err
	}
	return nil
}

type CharacterSelected struct {
	IP          string
	TCPPort     int32
	UDPPort     int32
	Key         string
	UserID      uint64
	CharID      uint64
	ZoneID      uint64
	ZoneName    string
	Location    string
	Slot        int32
	PrepPhase   int32
	Error       int32
	LoginServer string
}

func (s *CharacterSelected) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 54+len(s.IP)+len(s.Key)+len(s.ZoneName)+len(s.Location)+len(s.LoginServer)))
	writeString_7(b, s.IP)
	binary.Write(b, binary.LittleEndian, s.TCPPort)
	binary.Write(b, binary.LittleEndian, s.UDPPort)
	writeString_7(b, s.Key)
	binary.Write(b, binary.LittleEndian, s.UserID)
	binary.Write(b, binary.LittleEndian, s.CharID)
	binary.Write(b, binary.LittleEndian, s.ZoneID)
	writeString_7(b, s.ZoneName)
	writeString_7(b, s.Location)
	binary.Write(b, binary.LittleEndian, s.Slot)
	binary.Write(b, binary.LittleEndian, s.PrepPhase)
	binary.Write(b, binary.LittleEndian, s.Error)
	writeString_7(b, s.LoginServer)
	return b.Bytes(), nil
}

func (s *CharacterSelected) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.IP, err = readString_7(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.TCPPort); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.UDPPort); err != nil {
		return err
	}
	if s.Key, err = readString_7(b); err != nil {
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
	if s.ZoneName, err = readString_7(b); err != nil {
		return err
	}
	if s.Location, err = readString_7(b); err != nil {
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
	if s.LoginServer, err = readString_7(b); err != nil {
		return err
	}
	return nil
}

type CreateCharacter struct {
	CreationInfo string
}

func (s *CreateCharacter) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.CreationInfo)))
	writeString_7(b, s.CreationInfo)
	return b.Bytes(), nil
}

func (s *CreateCharacter) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.CreationInfo, err = readString_7(b); err != nil {
		return err
	}
	return nil
}

type CreateCharacterResponse struct {
	ErrorCode int32
}

func (s *CreateCharacterResponse) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	binary.Write(b, binary.LittleEndian, s.ErrorCode)
	return b.Bytes(), nil
}

func (s *CreateCharacterResponse) UnmarshalBinary(data []byte) error {
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

func (s *DeleteCharacter) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 8))
	binary.Write(b, binary.LittleEndian, s.CharID)
	return b.Bytes(), nil
}

func (s *DeleteCharacter) UnmarshalBinary(data []byte) error {
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

func (s *DeleteCharacterResponse) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	binary.Write(b, binary.LittleEndian, s.ErrorCode)
	return b.Bytes(), nil
}

func (s *DeleteCharacterResponse) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.ErrorCode); err != nil {
		return err
	}
	return nil
}

type RequestCharacterList struct {
}

func (s *RequestCharacterList) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (s *RequestCharacterList) UnmarshalBinary(data []byte) error {
	return nil
}

type RequestServerList struct {
}

func (s *RequestServerList) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (s *RequestServerList) UnmarshalBinary(data []byte) error {
	return nil
}

type SelectCharacter struct {
	CharID     uint64
	ServerName string
}

func (s *SelectCharacter) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 10+len(s.ServerName)))
	binary.Write(b, binary.LittleEndian, s.CharID)
	writeString_7(b, s.ServerName)
	return b.Bytes(), nil
}

func (s *SelectCharacter) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.CharID); err != nil {
		return err
	}
	if s.ServerName, err = readString_7(b); err != nil {
		return err
	}
	return nil
}

type ServerList struct {
}

func (s *ServerList) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (s *ServerList) UnmarshalBinary(data []byte) error {
	return nil
}

type StartCharacterList struct {
	LoginServer             string
	PurchasedCharacterSlots int32
}

func (s *StartCharacterList) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 6+len(s.LoginServer)))
	writeString_7(b, s.LoginServer)
	binary.Write(b, binary.LittleEndian, s.PurchasedCharacterSlots)
	return b.Bytes(), nil
}

func (s *StartCharacterList) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.LoginServer, err = readString_7(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.PurchasedCharacterSlots); err != nil {
		return err
	}
	return nil
}

type UserAuthen struct {
	Rec1          string
	Version       string
	Revision      string
	DataRevision  string
	CRC           string
	MachineID     uint64
	PatchClientID string
}

func (s *UserAuthen) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 20+len(s.Rec1)+len(s.Version)+len(s.Revision)+len(s.DataRevision)+len(s.CRC)+len(s.PatchClientID)))
	writeString_7(b, s.Rec1)
	writeString_7(b, s.Version)
	writeString_7(b, s.Revision)
	writeString_7(b, s.DataRevision)
	writeString_7(b, s.CRC)
	binary.Write(b, binary.LittleEndian, s.MachineID)
	writeString_7(b, s.PatchClientID)
	return b.Bytes(), nil
}

func (s *UserAuthen) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Rec1, err = readString_7(b); err != nil {
		return err
	}
	if s.Version, err = readString_7(b); err != nil {
		return err
	}
	if s.Revision, err = readString_7(b); err != nil {
		return err
	}
	if s.DataRevision, err = readString_7(b); err != nil {
		return err
	}
	if s.CRC, err = readString_7(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	if s.PatchClientID, err = readString_7(b); err != nil {
		return err
	}
	return nil
}

type UserAuthenRsp struct {
	Error      int32
	UserID     uint64
	Rec1       string
	Reason     string
	TimeStamp  string
	PayingUser int32
	Flags      int32
}

func (s *UserAuthenRsp) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 26+len(s.Rec1)+len(s.Reason)+len(s.TimeStamp)))
	binary.Write(b, binary.LittleEndian, s.Error)
	binary.Write(b, binary.LittleEndian, s.UserID)
	writeString_7(b, s.Rec1)
	writeString_7(b, s.Reason)
	writeString_7(b, s.TimeStamp)
	binary.Write(b, binary.LittleEndian, s.PayingUser)
	binary.Write(b, binary.LittleEndian, s.Flags)
	return b.Bytes(), nil
}

func (s *UserAuthenRsp) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Error); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.UserID); err != nil {
		return err
	}
	if s.Rec1, err = readString_7(b); err != nil {
		return err
	}
	if s.Reason, err = readString_7(b); err != nil {
		return err
	}
	if s.TimeStamp, err = readString_7(b); err != nil {
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
	UserID        uint64
	PassKey3      string
	MachineID     uint64
	Locale        string
	PatchClientID string
}

func (s *UserValidate) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 22+len(s.PassKey3)+len(s.Locale)+len(s.PatchClientID)))
	binary.Write(b, binary.LittleEndian, s.UserID)
	writeString_7(b, s.PassKey3)
	binary.Write(b, binary.LittleEndian, s.MachineID)
	writeString_7(b, s.Locale)
	writeString_7(b, s.PatchClientID)
	return b.Bytes(), nil
}

func (s *UserValidate) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.UserID); err != nil {
		return err
	}
	if s.PassKey3, err = readString_7(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	if s.Locale, err = readString_7(b); err != nil {
		return err
	}
	if s.PatchClientID, err = readString_7(b); err != nil {
		return err
	}
	return nil
}

type UserValidateRsp struct {
	Error      int32
	Reason     string
	UserID     uint64
	TimeStamp  string
	PayingUser int32
	Flags      int32
}

func (s *UserValidateRsp) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 24+len(s.Reason)+len(s.TimeStamp)))
	binary.Write(b, binary.LittleEndian, s.Error)
	writeString_7(b, s.Reason)
	binary.Write(b, binary.LittleEndian, s.UserID)
	writeString_7(b, s.TimeStamp)
	binary.Write(b, binary.LittleEndian, s.PayingUser)
	binary.Write(b, binary.LittleEndian, s.Flags)
	return b.Bytes(), nil
}

func (s *UserValidateRsp) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.Error); err != nil {
		return err
	}
	if s.Reason, err = readString_7(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.UserID); err != nil {
		return err
	}
	if s.TimeStamp, err = readString_7(b); err != nil {
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

func (s *DisconnectLoginAfk) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 1))
	binary.Write(b, binary.LittleEndian, s.Warning)
	return b.Bytes(), nil
}

func (s *DisconnectLoginAfk) UnmarshalBinary(data []byte) error {
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

func (s *LoginNotAfk) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	binary.Write(b, binary.LittleEndian, s.BadgeNameID)
	return b.Bytes(), nil
}

func (s *LoginNotAfk) UnmarshalBinary(data []byte) error {
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

func (s *ServerShutdown) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 4))
	binary.Write(b, binary.LittleEndian, s.Message)
	return b.Bytes(), nil
}

func (s *ServerShutdown) UnmarshalBinary(data []byte) error {
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

func (s *UserAdmitInd) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 8))
	binary.Write(b, binary.LittleEndian, s.Status)
	binary.Write(b, binary.LittleEndian, s.PositionInQueue)
	return b.Bytes(), nil
}

func (s *UserAdmitInd) UnmarshalBinary(data []byte) error {
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

func (s *WebCharacterInfo) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 12))
	binary.Write(b, binary.LittleEndian, s.Name)
	binary.Write(b, binary.LittleEndian, s.Gender)
	binary.Write(b, binary.LittleEndian, s.School)
	return b.Bytes(), nil
}

func (s *WebCharacterInfo) UnmarshalBinary(data []byte) error {
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
	Rec1          string
	Version       string
	Revision      string
	DataRevision  string
	CRC           string
	MachineID     uint64
	Locale        string
	PatchClientID string
}

func (s *UserAuthenV2) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 22+len(s.Rec1)+len(s.Version)+len(s.Revision)+len(s.DataRevision)+len(s.CRC)+len(s.Locale)+len(s.PatchClientID)))
	writeString_7(b, s.Rec1)
	writeString_7(b, s.Version)
	writeString_7(b, s.Revision)
	writeString_7(b, s.DataRevision)
	writeString_7(b, s.CRC)
	binary.Write(b, binary.LittleEndian, s.MachineID)
	writeString_7(b, s.Locale)
	writeString_7(b, s.PatchClientID)
	return b.Bytes(), nil
}

func (s *UserAuthenV2) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Rec1, err = readString_7(b); err != nil {
		return err
	}
	if s.Version, err = readString_7(b); err != nil {
		return err
	}
	if s.Revision, err = readString_7(b); err != nil {
		return err
	}
	if s.DataRevision, err = readString_7(b); err != nil {
		return err
	}
	if s.CRC, err = readString_7(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	if s.Locale, err = readString_7(b); err != nil {
		return err
	}
	if s.PatchClientID, err = readString_7(b); err != nil {
		return err
	}
	return nil
}

type SaveCharacter struct {
	CharID  uint64
	Success uint8
}

func (s *SaveCharacter) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 9))
	binary.Write(b, binary.LittleEndian, s.CharID)
	binary.Write(b, binary.LittleEndian, s.Success)
	return b.Bytes(), nil
}

func (s *SaveCharacter) UnmarshalBinary(data []byte) error {
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
	Rec1         string
	Version      string
	Revision     string
	DataRevision string
	CRC          string
	MachineID    uint64
}

func (s *WebAuthen) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 18+len(s.Rec1)+len(s.Version)+len(s.Revision)+len(s.DataRevision)+len(s.CRC)))
	writeString_7(b, s.Rec1)
	writeString_7(b, s.Version)
	writeString_7(b, s.Revision)
	writeString_7(b, s.DataRevision)
	writeString_7(b, s.CRC)
	binary.Write(b, binary.LittleEndian, s.MachineID)
	return b.Bytes(), nil
}

func (s *WebAuthen) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Rec1, err = readString_7(b); err != nil {
		return err
	}
	if s.Version, err = readString_7(b); err != nil {
		return err
	}
	if s.Revision, err = readString_7(b); err != nil {
		return err
	}
	if s.DataRevision, err = readString_7(b); err != nil {
		return err
	}
	if s.CRC, err = readString_7(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	return nil
}

type WebValidate struct {
	UserID    uint64
	PassKey3  string
	MachineID uint64
	Locale    string
}

func (s *WebValidate) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 20+len(s.PassKey3)+len(s.Locale)))
	binary.Write(b, binary.LittleEndian, s.UserID)
	writeString_7(b, s.PassKey3)
	binary.Write(b, binary.LittleEndian, s.MachineID)
	writeString_7(b, s.Locale)
	return b.Bytes(), nil
}

func (s *WebValidate) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.UserID); err != nil {
		return err
	}
	if s.PassKey3, err = readString_7(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	if s.Locale, err = readString_7(b); err != nil {
		return err
	}
	return nil
}

type ChangeCharacterName struct {
	CharID     uint64
	NewName    string
	ServerName string
}

func (s *ChangeCharacterName) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 12+len(s.NewName)+len(s.ServerName)))
	binary.Write(b, binary.LittleEndian, s.CharID)
	writeString_7(b, s.NewName)
	writeString_7(b, s.ServerName)
	return b.Bytes(), nil
}

func (s *ChangeCharacterName) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.CharID); err != nil {
		return err
	}
	if s.NewName, err = readString_7(b); err != nil {
		return err
	}
	if s.ServerName, err = readString_7(b); err != nil {
		return err
	}
	return nil
}

type UserAuthenV3 struct {
	Rec1           string
	Version        string
	Revision       string
	DataRevision   string
	CRC            string
	MachineID      uint64
	Locale         string
	PatchClientID  string
	IsSteamPatcher uint32
	ConsoleType    uint8
}

func (s *UserAuthenV3) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 27+len(s.Rec1)+len(s.Version)+len(s.Revision)+len(s.DataRevision)+len(s.CRC)+len(s.Locale)+len(s.PatchClientID)))
	writeString_7(b, s.Rec1)
	writeString_7(b, s.Version)
	writeString_7(b, s.Revision)
	writeString_7(b, s.DataRevision)
	writeString_7(b, s.CRC)
	binary.Write(b, binary.LittleEndian, s.MachineID)
	writeString_7(b, s.Locale)
	writeString_7(b, s.PatchClientID)
	binary.Write(b, binary.LittleEndian, s.IsSteamPatcher)
	binary.Write(b, binary.LittleEndian, s.ConsoleType)
	return b.Bytes(), nil
}

func (s *UserAuthenV3) UnmarshalBinary(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Rec1, err = readString_7(b); err != nil {
		return err
	}
	if s.Version, err = readString_7(b); err != nil {
		return err
	}
	if s.Revision, err = readString_7(b); err != nil {
		return err
	}
	if s.DataRevision, err = readString_7(b); err != nil {
		return err
	}
	if s.CRC, err = readString_7(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.MachineID); err != nil {
		return err
	}
	if s.Locale, err = readString_7(b); err != nil {
		return err
	}
	if s.PatchClientID, err = readString_7(b); err != nil {
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

func (s *LoginLogCharacterCreation) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 8))
	binary.Write(b, binary.LittleEndian, s.Stage)
	binary.Write(b, binary.LittleEndian, s.Parameter)
	return b.Bytes(), nil
}

func (s *LoginLogCharacterCreation) UnmarshalBinary(data []byte) error {
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

func writeString_7(b *bytes.Buffer, v string) {
	binary.Write(b, binary.LittleEndian, uint16(len(v)))
	b.WriteString(v)
}

func readString_7(buf *bytes.Reader) (string, error) {
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
