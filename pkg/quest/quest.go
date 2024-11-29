// Code generated by w101-client-go. DO NOT EDIT.
package quest

import (
	"bytes"
	"encoding/binary"
	"github.com/cedws/w101-client-go/codegen"
	"github.com/cedws/w101-client-go/proto"
)

type service interface {
	AcceptQuest(AcceptQuest)
	CanAcquireWorldElixir(CanAcquireWorldElixir)
	CompleteGoal(CompleteGoal)
	CompletePersonaGoal(CompletePersonaGoal)
	CompleteQuest(CompleteQuest)
	DebugQuestData(DebugQuestData)
	DeclineQuest(DeclineQuest)
	GetNextWorld(GetNextWorld)
	InteractNPC(InteractNPC)
	NPCInfo(NPCInfo)
	PersonaComplete(PersonaComplete)
	QuestOffer(QuestOffer)
	QuestReadyToTurnIn(QuestReadyToTurnIn)
	RemoveGoal(RemoveGoal)
	RemoveQuest(RemoveQuest)
	SendGoal(SendGoal)
	SendNPCOptions(SendNPCOptions)
	SendQuest(SendQuest)
	UpdateUserHiddenQuests(UpdateUserHiddenQuests)
}

func (Service) AcceptQuest(AcceptQuest)                       {}
func (Service) CanAcquireWorldElixir(CanAcquireWorldElixir)   {}
func (Service) CompleteGoal(CompleteGoal)                     {}
func (Service) CompletePersonaGoal(CompletePersonaGoal)       {}
func (Service) CompleteQuest(CompleteQuest)                   {}
func (Service) DebugQuestData(DebugQuestData)                 {}
func (Service) DeclineQuest(DeclineQuest)                     {}
func (Service) GetNextWorld(GetNextWorld)                     {}
func (Service) InteractNPC(InteractNPC)                       {}
func (Service) NPCInfo(NPCInfo)                               {}
func (Service) PersonaComplete(PersonaComplete)               {}
func (Service) QuestOffer(QuestOffer)                         {}
func (Service) QuestReadyToTurnIn(QuestReadyToTurnIn)         {}
func (Service) RemoveGoal(RemoveGoal)                         {}
func (Service) RemoveQuest(RemoveQuest)                       {}
func (Service) SendGoal(SendGoal)                             {}
func (Service) SendNPCOptions(SendNPCOptions)                 {}
func (Service) SendQuest(SendQuest)                           {}
func (Service) UpdateUserHiddenQuests(UpdateUserHiddenQuests) {}

func RegisterService(r *proto.MessageRouter, s service) {
	proto.RegisterMessageHandler(r, 52, 1, s.AcceptQuest)
	proto.RegisterMessageHandler(r, 52, 2, s.CanAcquireWorldElixir)
	proto.RegisterMessageHandler(r, 52, 3, s.CompleteGoal)
	proto.RegisterMessageHandler(r, 52, 4, s.CompletePersonaGoal)
	proto.RegisterMessageHandler(r, 52, 5, s.CompleteQuest)
	proto.RegisterMessageHandler(r, 52, 6, s.DebugQuestData)
	proto.RegisterMessageHandler(r, 52, 7, s.DeclineQuest)
	proto.RegisterMessageHandler(r, 52, 8, s.GetNextWorld)
	proto.RegisterMessageHandler(r, 52, 9, s.InteractNPC)
	proto.RegisterMessageHandler(r, 52, 10, s.NPCInfo)
	proto.RegisterMessageHandler(r, 52, 11, s.PersonaComplete)
	proto.RegisterMessageHandler(r, 52, 12, s.QuestOffer)
	proto.RegisterMessageHandler(r, 52, 13, s.QuestReadyToTurnIn)
	proto.RegisterMessageHandler(r, 52, 14, s.RemoveGoal)
	proto.RegisterMessageHandler(r, 52, 15, s.RemoveQuest)
	proto.RegisterMessageHandler(r, 52, 16, s.SendGoal)
	proto.RegisterMessageHandler(r, 52, 17, s.SendNPCOptions)
	proto.RegisterMessageHandler(r, 52, 18, s.SendQuest)
	proto.RegisterMessageHandler(r, 52, 19, s.UpdateUserHiddenQuests)
}

func NewClient(c *proto.Client) Client {
	return Client{c}
}

func (c Client) AcceptQuest(m *AcceptQuest) error {
	return c.c.WriteMessage(52, 1, m)
}

func (c Client) CanAcquireWorldElixir(m *CanAcquireWorldElixir) error {
	return c.c.WriteMessage(52, 2, m)
}

func (c Client) CompleteGoal(m *CompleteGoal) error {
	return c.c.WriteMessage(52, 3, m)
}

func (c Client) CompletePersonaGoal(m *CompletePersonaGoal) error {
	return c.c.WriteMessage(52, 4, m)
}

func (c Client) CompleteQuest(m *CompleteQuest) error {
	return c.c.WriteMessage(52, 5, m)
}

func (c Client) DebugQuestData(m *DebugQuestData) error {
	return c.c.WriteMessage(52, 6, m)
}

func (c Client) DeclineQuest(m *DeclineQuest) error {
	return c.c.WriteMessage(52, 7, m)
}

func (c Client) GetNextWorld(m *GetNextWorld) error {
	return c.c.WriteMessage(52, 8, m)
}

func (c Client) InteractNPC(m *InteractNPC) error {
	return c.c.WriteMessage(52, 9, m)
}

func (c Client) NPCInfo(m *NPCInfo) error {
	return c.c.WriteMessage(52, 10, m)
}

func (c Client) PersonaComplete(m *PersonaComplete) error {
	return c.c.WriteMessage(52, 11, m)
}

func (c Client) QuestOffer(m *QuestOffer) error {
	return c.c.WriteMessage(52, 12, m)
}

func (c Client) QuestReadyToTurnIn(m *QuestReadyToTurnIn) error {
	return c.c.WriteMessage(52, 13, m)
}

func (c Client) RemoveGoal(m *RemoveGoal) error {
	return c.c.WriteMessage(52, 14, m)
}

func (c Client) RemoveQuest(m *RemoveQuest) error {
	return c.c.WriteMessage(52, 15, m)
}

func (c Client) SendGoal(m *SendGoal) error {
	return c.c.WriteMessage(52, 16, m)
}

func (c Client) SendNPCOptions(m *SendNPCOptions) error {
	return c.c.WriteMessage(52, 17, m)
}

func (c Client) SendQuest(m *SendQuest) error {
	return c.c.WriteMessage(52, 18, m)
}

func (c Client) UpdateUserHiddenQuests(m *UpdateUserHiddenQuests) error {
	return c.c.WriteMessage(52, 19, m)
}

type Service struct {
	service
}

type Client struct {
	c *proto.Client
}
type AcceptQuest struct {
	QuestName string
	MobileID  uint64
}

func (s *AcceptQuest) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 10+len(s.QuestName)))
	binary.Write(b, binary.LittleEndian, s.QuestName)
	return b.Bytes()
}

func (s *AcceptQuest) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.MobileID); err != nil {
		return err
	}
	if s.QuestName, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type CanAcquireWorldElixir struct {
	ElixirTemplateGID uint64
	Result            int32
}

func (s *CanAcquireWorldElixir) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 12))
	return b.Bytes()
}

func (s *CanAcquireWorldElixir) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.ElixirTemplateGID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Result); err != nil {
		return err
	}
	return nil
}

type CompleteGoal struct {
	CompleteText string
	QuestID      uint64
	GoalID       uint64
}

func (s *CompleteGoal) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 18+len(s.CompleteText)))
	binary.Write(b, binary.LittleEndian, s.CompleteText)
	return b.Bytes()
}

func (s *CompleteGoal) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.QuestID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.GoalID); err != nil {
		return err
	}
	if s.CompleteText, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type CompletePersonaGoal struct {
	MobileID uint64
	QuestID  uint64
	GoalID   uint64
}

func (s *CompletePersonaGoal) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 24))
	return b.Bytes()
}

func (s *CompletePersonaGoal) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.MobileID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.QuestID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.GoalID); err != nil {
		return err
	}
	return nil
}

type CompleteQuest struct {
	CompleteText string
	QuestID      uint64
}

func (s *CompleteQuest) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 10+len(s.CompleteText)))
	binary.Write(b, binary.LittleEndian, s.CompleteText)
	return b.Bytes()
}

func (s *CompleteQuest) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.QuestID); err != nil {
		return err
	}
	if s.CompleteText, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type DebugQuestData struct {
	GoalName  string
	QuestName string
	QuestID   uint64
	GoalID    uint32
}

func (s *DebugQuestData) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 16+len(s.QuestName)+len(s.GoalName)))
	binary.Write(b, binary.LittleEndian, s.QuestName)
	binary.Write(b, binary.LittleEndian, s.GoalName)
	return b.Bytes()
}

func (s *DebugQuestData) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.QuestID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.GoalID); err != nil {
		return err
	}
	if s.QuestName, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.GoalName, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type DeclineQuest struct {
	QuestName string
}

func (s *DeclineQuest) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.QuestName)))
	binary.Write(b, binary.LittleEndian, s.QuestName)
	return b.Bytes()
}

func (s *DeclineQuest) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.QuestName, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type GetNextWorld struct {
	World string
}

func (s *GetNextWorld) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.World)))
	binary.Write(b, binary.LittleEndian, s.World)
	return b.Bytes()
}

func (s *GetNextWorld) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.World, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type InteractNPC struct {
	ServiceName        string
	GlobalID           uint64
	Reinteract         int32
	ServiceIndex       uint32
	RequestedSigilMode uint32
}

func (s *InteractNPC) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 22+len(s.ServiceName)))
	binary.Write(b, binary.LittleEndian, s.ServiceName)
	return b.Bytes()
}

func (s *InteractNPC) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.GlobalID); err != nil {
		return err
	}
	if s.ServiceName, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Reinteract); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ServiceIndex); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.RequestedSigilMode); err != nil {
		return err
	}
	return nil
}

type NPCInfo struct {
	PersonaMadlibs string
	Greeting       string
	Name           string
	MobileID       uint64
}

func (s *NPCInfo) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 14+len(s.Name)+len(s.Greeting)+len(s.PersonaMadlibs)))
	binary.Write(b, binary.LittleEndian, s.Name)
	binary.Write(b, binary.LittleEndian, s.Greeting)
	binary.Write(b, binary.LittleEndian, s.PersonaMadlibs)
	return b.Bytes()
}

func (s *NPCInfo) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.MobileID); err != nil {
		return err
	}
	if s.Name, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Greeting, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.PersonaMadlibs, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type PersonaComplete struct {
	GoalHyperlink string
	MobileID      uint64
	QuestID       uint64
	GoalID        uint64
}

func (s *PersonaComplete) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 26+len(s.GoalHyperlink)))
	binary.Write(b, binary.LittleEndian, s.GoalHyperlink)
	return b.Bytes()
}

func (s *PersonaComplete) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.MobileID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.QuestID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.GoalID); err != nil {
		return err
	}
	if s.GoalHyperlink, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}

type QuestOffer struct {
	GoalData   string
	Rewards    string
	QuestInfo  string
	QuestTitle string
	QuestName  string
	MobileID   uint64
	Level      int32
	Mainline   uint8
}

func (s *QuestOffer) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 23+len(s.QuestName)+len(s.QuestTitle)+len(s.QuestInfo)+len(s.Rewards)+len(s.GoalData)))
	binary.Write(b, binary.LittleEndian, s.QuestName)
	binary.Write(b, binary.LittleEndian, s.QuestTitle)
	binary.Write(b, binary.LittleEndian, s.QuestInfo)
	binary.Write(b, binary.LittleEndian, s.Rewards)
	binary.Write(b, binary.LittleEndian, s.GoalData)
	return b.Bytes()
}

func (s *QuestOffer) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.MobileID); err != nil {
		return err
	}
	if s.QuestName, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.QuestTitle, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.QuestInfo, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Level); err != nil {
		return err
	}
	if s.Rewards, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.GoalData, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Mainline); err != nil {
		return err
	}
	return nil
}

type QuestReadyToTurnIn struct {
	QuestID uint64
}

func (s *QuestReadyToTurnIn) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 8))
	return b.Bytes()
}

func (s *QuestReadyToTurnIn) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.QuestID); err != nil {
		return err
	}
	return nil
}

type RemoveGoal struct {
	QuestID uint64
	GoalID  uint64
}

func (s *RemoveGoal) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 16))
	return b.Bytes()
}

func (s *RemoveGoal) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.QuestID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.GoalID); err != nil {
		return err
	}
	return nil
}

type RemoveQuest struct {
	QuestID uint64
	NpcID   uint64
}

func (s *RemoveQuest) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 16))
	return b.Bytes()
}

func (s *RemoveQuest) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.QuestID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.NpcID); err != nil {
		return err
	}
	return nil
}

type SendGoal struct {
	PersonaName         string
	GoalImage2          string
	GoalMadlibs         string
	GoalDestinationZone string
	GoalLocation        string
	GoalTitle           string
	ClientTags          string
	PatronIcon          string
	TallyText           string
	GoalImage1          string
	QuestID             uint64
	GoalID              uint64
	SubscriberGoalTotal int32
	SendType            int32
	GoalTotal           int32
	GoalNameID          uint32
	GoalCount           int32
	GoalStatus          uint8
	UseTally            uint8
	GoalType            uint8
	NoQuestHelper       uint8
	PetOnlyQuest        uint8
	HasActiveResults    uint8
	HideGoalFloatyText  uint8
}

func (s *SendGoal) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 63+len(s.GoalTitle)+len(s.GoalLocation)+len(s.GoalDestinationZone)+len(s.GoalImage1)+len(s.GoalImage2)+len(s.PersonaName)+len(s.TallyText)+len(s.GoalMadlibs)+len(s.ClientTags)+len(s.PatronIcon)))
	binary.Write(b, binary.LittleEndian, s.GoalTitle)
	binary.Write(b, binary.LittleEndian, s.GoalLocation)
	binary.Write(b, binary.LittleEndian, s.GoalDestinationZone)
	binary.Write(b, binary.LittleEndian, s.GoalImage1)
	binary.Write(b, binary.LittleEndian, s.GoalImage2)
	binary.Write(b, binary.LittleEndian, s.PersonaName)
	binary.Write(b, binary.LittleEndian, s.TallyText)
	binary.Write(b, binary.LittleEndian, s.GoalMadlibs)
	binary.Write(b, binary.LittleEndian, s.ClientTags)
	binary.Write(b, binary.LittleEndian, s.PatronIcon)
	return b.Bytes()
}

func (s *SendGoal) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.QuestID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.GoalID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.GoalNameID); err != nil {
		return err
	}
	if s.GoalTitle, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.GoalLocation, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.GoalDestinationZone, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.GoalImage1, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.GoalImage2, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.PersonaName, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.GoalType); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.GoalStatus); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.GoalCount); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.GoalTotal); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.SubscriberGoalTotal); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.UseTally); err != nil {
		return err
	}
	if s.TallyText, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.SendType); err != nil {
		return err
	}
	if s.GoalMadlibs, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.ClientTags, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.PatronIcon, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.NoQuestHelper); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.PetOnlyQuest); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.HasActiveResults); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.HideGoalFloatyText); err != nil {
		return err
	}
	return nil
}

type SendNPCOptions struct {
	Options    string
	MobileID   uint64
	Reinteract int32
}

func (s *SendNPCOptions) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 14+len(s.Options)))
	binary.Write(b, binary.LittleEndian, s.Options)
	return b.Bytes()
}

func (s *SendNPCOptions) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.MobileID); err != nil {
		return err
	}
	if s.Options, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Reinteract); err != nil {
		return err
	}
	return nil
}

type SendQuest struct {
	QuestInfo        string
	QuestTitle       string
	Rewards          string
	ClientTags       string
	AssociatedWorlds string
	QuestMadlibs     string
	GoalData         string
	QuestID          uint64
	QuestLevel       int32
	QuestType        uint32
	QuestNameID      uint32
	New              uint8
	NoQuestHelper    uint8
	Mainline         uint8
	ReadyToTurnIn    uint8
	SkipQHAutoSelect uint8
	PetOnlyQuest     uint8
	ActivityType     uint8
}

func (s *SendQuest) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 41+len(s.QuestTitle)+len(s.QuestInfo)+len(s.QuestMadlibs)+len(s.GoalData)+len(s.Rewards)+len(s.ClientTags)+len(s.AssociatedWorlds)))
	binary.Write(b, binary.LittleEndian, s.QuestTitle)
	binary.Write(b, binary.LittleEndian, s.QuestInfo)
	binary.Write(b, binary.LittleEndian, s.QuestMadlibs)
	binary.Write(b, binary.LittleEndian, s.GoalData)
	binary.Write(b, binary.LittleEndian, s.Rewards)
	binary.Write(b, binary.LittleEndian, s.ClientTags)
	binary.Write(b, binary.LittleEndian, s.AssociatedWorlds)
	return b.Bytes()
}

func (s *SendQuest) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if err = binary.Read(b, binary.LittleEndian, &s.QuestID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.QuestNameID); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.QuestType); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.QuestLevel); err != nil {
		return err
	}
	if s.QuestTitle, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.QuestInfo, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.New); err != nil {
		return err
	}
	if s.QuestMadlibs, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.GoalData, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.Rewards, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.ClientTags, err = codegen.ReadString(b); err != nil {
		return err
	}
	if s.AssociatedWorlds, err = codegen.ReadString(b); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.NoQuestHelper); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.Mainline); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ReadyToTurnIn); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.SkipQHAutoSelect); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.PetOnlyQuest); err != nil {
		return err
	}
	if err = binary.Read(b, binary.LittleEndian, &s.ActivityType); err != nil {
		return err
	}
	return nil
}

type UpdateUserHiddenQuests struct {
	Data string
}

func (s *UpdateUserHiddenQuests) Marshal() []byte {
	b := bytes.NewBuffer(make([]byte, 0, 2+len(s.Data)))
	binary.Write(b, binary.LittleEndian, s.Data)
	return b.Bytes()
}

func (s *UpdateUserHiddenQuests) Unmarshal(data []byte) error {
	b := bytes.NewReader(data)
	var err error
	if s.Data, err = codegen.ReadString(b); err != nil {
		return err
	}
	return nil
}
