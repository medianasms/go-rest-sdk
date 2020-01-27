package medianasms

import (
	"encoding/json"
	"fmt"
	"time"
)

// MessageStatus message status
type MessageStatus string

// TODO: Add other message status codes
const (
	// MessageStatusActive ...
	MessageStatusActive MessageStatus = "active"
)

// MessageType message type
type MessageType string

// TODO: Add other message types
const (
	// MessageTypeNormal normal message
	MessageTypeNormal MessageType = "normal"
)

// MessageConfirmState message confirm state
type MessageConfirmState string

const (
	// MessageConfirmeStatePending pending
	MessageConfirmeStatePending MessageConfirmState = "pending"
	// MessageConfirmeStateConfirmed confirmed
	MessageConfirmeStateConfirmed MessageConfirmState = "confirmed"
	// MessageConfirmeStateRejected rejected
	MessageConfirmeStateRejected MessageConfirmState = "rejected"
)

// PatternStatus ...
type PatternStatus string

const (
	// PatternStatusActive active
	PatternStatusActive PatternStatus = "active"
	// PatternStatusInactive inactive
	PatternStatusInactive PatternStatus = "inactive"
	// PatternStatusPending pending
	PatternStatusPending PatternStatus = "pending"
)

// Message message model
type Message struct {
	BulkID               int64               `json:"bulk_id"`
	Number               string              `json:"number"`
	Message              string              `json:"message"`
	Status               MessageStatus       `json:"status"`
	Type                 MessageType         `json:"type"`
	ConfirmState         MessageConfirmState `json:"confirm_state"`
	CreatedAt            time.Time           `json:"created_at"`
	SentAt               time.Time           `json:"sent_at"`
	RecipientsCount      int64               `json:"recipients_count"`
	ValidRecipientsCount int64               `json:"valid_recipients_count"`
	Page                 int64               `json:"page"`
	Cost                 float64             `json:"cost"`
	PaybackCost          float64             `json:"payback_cost"`
	Description          string              `json:"description"`
}

// MessageRecipient message recipient status
type MessageRecipient struct {
	Recipient string `json:"recipient"`
	Status    string `json:"status"`
}

// InboxMessage inbox message
type InboxMessage struct {
	Number    string    `json:"number"`
	Message   string    `json:"message"`
	Sender    string    `json:"sender"`
	CreatedAt time.Time `json:"time"`
	Type      string    `json:"type"`
}

// Pattern pattern
type Pattern struct {
	Code     string        `json:"code"`
	Status   PatternStatus `json:"status"`
	Message  string        `json:"message"`
	IsShared bool          `json:"is_shared"`
}

// sendSMSReqType request type for send sms
type sendSMSReqType struct {
	Originator string   `json:"originator"`
	Recipients []string `json:"recipients"`
	Message    string   `json:"message"`
}

// sendResType response type for send sms
type sendResType struct {
	BulkID int64 `json:"bulk_id"`
}

// getMessageResType get message by bulk id response template
type getMessageResType struct {
	Message *Message `json:"message"`
}

// fetchMessageStatusesResType get message statuses response template
type fetchMessageStatusesResType struct {
	Statuses []MessageRecipient `json:"recipients"`
}

// fetchInboxResType fetch inbox response template
type fetchInboxResType struct {
	Messages []InboxMessage `json:"messages"`
}

// createPatternReqType create pattern request type
type createPatternReqType struct {
	Pattern  string `json:"pattern"`
	IsShared bool   `json:"is_shared"`
}

// sendPatternReqType send sms with pattern request template
type sendPatternReqType struct {
	PatternCode string            `json:"pattern_code"`
	Originator  string            `json:"originator"`
	Recipient   string            `json:"recipient"`
	Values      map[string]string `json:"values"`
}

// Send send a message
func (sms *MedianaSMS) Send(originator string, recipients []string, message string) (int64, error) {
	data := sendSMSReqType{
		Originator: originator,
		Recipients: recipients,
		Message:    message,
	}

	_res, err := sms.post("/messages", "application/json", data)
	if err != nil {
		return 0, err
	}

	res := sendResType{}
	if err = json.Unmarshal(_res.Data, &res); err != nil {
		return 0, err
	}

	return res.BulkID, nil
}

// GetMessage get a message by bulk_id
func (sms *MedianaSMS) GetMessage(bulkID int64) (*Message, error) {
	_res, err := sms.get(fmt.Sprintf("/messages/%d", bulkID), nil)
	if err != nil {
		return nil, err
	}

	res := &getMessageResType{}
	if err = json.Unmarshal(_res.Data, res); err != nil {
		return nil, err
	}

	return res.Message, nil
}

// FetchStatuses get message recipients statuses
func (sms *MedianaSMS) FetchStatuses(bulkID int64, pp ListParams) ([]MessageRecipient, *PaginationInfo, error) {
	_res, err := sms.get(fmt.Sprintf("/messages/%d/recipients", bulkID), map[string]string{
		"page":  fmt.Sprintf("%d", pp.Page),
		"limit": fmt.Sprintf("%d", pp.Limit),
	})
	if err != nil {
		return nil, nil, err
	}

	res := &fetchMessageStatusesResType{}
	if err = json.Unmarshal(_res.Data, res); err != nil {
		return nil, nil, err
	}

	return res.Statuses, _res.Meta, nil
}

// FetchInbox fetch inbox messages list
func (sms *MedianaSMS) FetchInbox(pp ListParams) ([]InboxMessage, *PaginationInfo, error) {
	_res, err := sms.get("/messages/inbox", map[string]string{
		"page":  fmt.Sprintf("%d", pp.Page),
		"limit": fmt.Sprintf("%d", pp.Limit),
	})
	if err != nil {
		return nil, nil, err
	}

	res := &fetchInboxResType{}
	if err = json.Unmarshal(_res.Data, res); err != nil {
		return nil, nil, err
	}

	return res.Messages, _res.Meta, nil
}

// CreatePattern create new pattern
func (sms *MedianaSMS) CreatePattern(pattern string, isShared bool) (*Pattern, error) {
	data := createPatternReqType{
		Pattern:  pattern,
		IsShared: isShared,
	}

	_res, err := sms.post("/messages/patterns", "application/json", data)
	if err != nil {
		return nil, err
	}

	res := &Pattern{}
	if err = json.Unmarshal(_res.Data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SendPattern send a message with pattern
func (sms *MedianaSMS) SendPattern(patternCode string, originator string, recipient string, values map[string]string) (int64, error) {
	data := sendPatternReqType{
		PatternCode: patternCode,
		Originator:  originator,
		Recipient:   recipient,
		Values:      values,
	}

	_res, err := sms.post("/messages/patterns/send", "application/json", data)
	if err != nil {
		return 0, err
	}

	res := sendResType{}
	if err = json.Unmarshal(_res.Data, &res); err != nil {
		return 0, err
	}

	return res.BulkID, nil
}
