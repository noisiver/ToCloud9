package events

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type ChatServiceProducer interface {
	IncomingWhisper(payload *ChatEventIncomingWhisperPayload) error
}

type chatServiceProducerNatsJSON struct {
	conn           *nats.Conn
	ver            string
	loadBalancerID string
}

func NewChatServiceProducerNatsJSON(conn *nats.Conn, ver string, loadBalancerID string) ChatServiceProducer {
	return &chatServiceProducerNatsJSON{
		conn:           conn,
		ver:            ver,
		loadBalancerID: loadBalancerID,
	}
}

func (c *chatServiceProducerNatsJSON) IncomingWhisper(payload *ChatEventIncomingWhisperPayload) error {
	return c.publish(ChatEventIncomingWhisper, payload)
}

func (c *chatServiceProducerNatsJSON) publish(e ChatServiceEvent, payload interface{}) error {
	msg := EventToSendGenericPayload{
		Version:   c.ver,
		EventType: int(e),
		Payload:   payload,
	}

	d, err := json.Marshal(&msg)
	if err != nil {
		return err
	}

	return c.conn.Publish(e.SubjectName(c.loadBalancerID), d)
}
