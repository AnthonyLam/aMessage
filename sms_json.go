package main

import (
	"json"
)

const (
	MESSAGE_TYPE_ALL    = iota
	MESSAGE_TYPE_INBOX  = iota
	MESSAGE_TYPE_SENT   = iota
	MESSAGE_TYPE_DRAFT  = iota
	MESSAGE_TYPE_OUTBOX = iota
	MESSAGE_TYPE_FAILED = iota
	MESSAGE_TYPE_QUEUED = iota
)

type Message struct {
	address       string
	body          string
	date_received uint64
	date_sent     uint64
	person        uint64 // Phone Number
}
