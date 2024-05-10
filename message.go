package camunda_client_go

import "log/slog"

// Message a client for Message API
type Message struct {
	client *Client
}

// ReqMessage a request to send a message
type ReqMessage struct {
	MessageName       string               `json:"messageName"`
	BusinessKey       string               `json:"businessKey,omitempty"`
	ProcessInstanceId *string              `json:"processInstanceId,omitempty"`
	ProcessVariables  *map[string]Variable `json:"processVariables,omitempty"`
	CorrelationKeys   *map[string]Variable `json:"correlationKeys,omitempty"`
}

// SendMessage sends message to a process
func (m *Message) SendMessage(query *ReqMessage) error {
	res, err := m.client.doPostJson("/message", map[string]string{}, query)
	if res != nil {
		if closeErr := res.Body.Close(); closeErr != nil {
			slog.Default().Error("Failed to close response body", "error", closeErr)
		}
	}
	return err
}
