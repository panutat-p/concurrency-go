package http_request

import (
	"testing"
)

func TestSendSyncReq(t *testing.T) {
	SendSyncReq()
}

func TestSendAsyncReqChannel(t *testing.T) {
	SendAsyncReqChannel()
}

func TestSendAsyncReqRoutine(t *testing.T) {
	SendAsyncReqRoutine()
}
