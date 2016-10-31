package main

import (
	"encoding/json"
	"fmt"
)

type MessageType int32

const (
	Sms_MESSAGE_TYPE_ALL    MessageType = iota
	Sms_MESSAGE_TYPE_INBOX  MessageType = iota
	Sms_MESSAGE_TYPE_SENT   MessageType = iota
	Sms_MESSAGE_TYPE_DRAFT  MessageType = iota
	Sms_MESSAGE_TYPE_OUTBOX MessageType = iota
	Sms_MESSAGE_TYPE_FAILED MessageType = iota
	Sms_MESSAGE_TYPE_QUEUED MessageType = iota
)

type Sms struct {
	Address      string
	Body         string
	DateReceived uint64
	DateSent     uint64
	Person       uint64 // Phone Number
	Type         MessageType
}

func (m *Sms) String() string {
	return fmt.Sprintf("Address: %s\nBody: %s\n", m.Address, m.Body)
}

func (m *Sms) Bytes() []byte {
	res, _ := json.Marshal(*m)
	return res
}

func NewSms(response []byte) (sms Sms) {
	json.Unmarshal(response, &sms)
	return
}
