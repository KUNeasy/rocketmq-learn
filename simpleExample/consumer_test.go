package simpleExample

import "testing"

func TestSendSyncMessage(t *testing.T) {
	SendSyncMessage()
}

func TestSendAsyncMessage(t *testing.T) {
	SendAsyncMessage()
}

func TestSendOneWayMessage(t *testing.T) {
	SendOneWayMessage()
}

func TestConsumeMessage(t *testing.T) {
	ConsumeMessage()
}
