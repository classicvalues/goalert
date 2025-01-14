package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/target/goalert/config"
	"github.com/target/goalert/notification"
)

type Sender struct{}

// POSTDataAlert represents fields in outgoing alert notification.
type POSTDataAlert struct {
	Type    string
	AlertID int
	Summary string
	Details string
}

// POSTDataAlertBundle represents fields in outgoing alert bundle notification.
type POSTDataAlertBundle struct {
	Type        string
	ServiceID   string
	ServiceName string
	Count       int
}

// POSTDataAlertStatus represents fields in outgoing alert status notification.
type POSTDataAlertStatus struct {
	Type     string
	AlertID  int
	LogEntry string
}

// POSTDataAlertStatusBundle represents fields in outgoing alert status bundle notification.
type POSTDataAlertStatusBundle struct {
	Type     string
	AlertID  int
	LogEntry string
	Count    int
}

// POSTDataVerification represents fields in outgoing verification notification.
type POSTDataVerification struct {
	Type string
	Code string
}

// POSTDataTest represents fields in outgoing test notification.
type POSTDataTest struct {
	Type string
}

func NewSender(ctx context.Context) *Sender {
	return &Sender{}
}

// Send will send an alert for the provided message type
func (s *Sender) Send(ctx context.Context, msg notification.Message) (string, *notification.Status, error) {
	var payload interface{}
	switch m := msg.(type) {
	case notification.Test:
		payload = POSTDataTest{
			Type: "Test",
		}
	case notification.Verification:
		payload = POSTDataVerification{
			Type: "Verification",
			Code: strconv.Itoa(m.Code),
		}
	case notification.Alert:
		payload = POSTDataAlert{
			Type:    "Alert",
			Details: m.Details,
			AlertID: m.AlertID,
			Summary: m.Summary,
		}
	case notification.AlertBundle:
		payload = POSTDataAlertBundle{
			Type:        "AlertBundle",
			ServiceID:   m.ServiceID,
			ServiceName: m.ServiceName,
			Count:       m.Count,
		}
	case notification.AlertStatus:
		payload = POSTDataAlertStatus{
			Type:     "AlertStatus",
			AlertID:  m.AlertID,
			LogEntry: m.LogEntry,
		}
	default:
		return "", nil, fmt.Errorf("message type '%s' not supported", m.Type().String())
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	cfg := config.FromContext(ctx)
	if !cfg.ValidWebhookURL(msg.Destination().Value) {
		// fail permanently if the URL is not currently valid/allowed
		return "", &notification.Status{
			State:   notification.StateFailedPerm,
			Details: "invalid or not allowed URL",
		}, nil
	}

	req, err := http.NewRequestWithContext(ctx, "POST", msg.Destination().Value, bytes.NewReader(data))
	if err != nil {
		return "", nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return "", nil, err
	}

	return "", &notification.Status{State: notification.StateSent}, nil
}
