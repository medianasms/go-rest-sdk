package medianasms

import (
	"testing"
)

// TestSend test send sms
func TestSend(t *testing.T) {
	sms := New("API-KEY")

	bulkID, err := sms.Send("+9810001", []string{"98912xxxxxxx"}, "mediana is awesome")
	if err != nil {
		t.Error("error occurred ", err)
	}

	t.Log(bulkID)
}

// TestErrors test api errors handling
func TestErrors(t *testing.T) {
	sms := New("API_KEY")

	_, err := sms.Send("9810001", []string{"98912xxxxx"}, "mediana is awesome")
	if err != nil {
		if e, ok := err.(Error); ok {
			switch e.Code {
			case ErrUnprocessableEntity:
				fieldErrors := e.Message.(FieldErrs)
				for field, fieldError := range fieldErrors {
					t.Log(field, fieldError)
				}
			default:
				errMsg := e.Message.(string)
				t.Log(errMsg)
			}
		}
	}
}

// TestGetMessage tests getMessage method
func TestGetMessage(t *testing.T) {
	sms := New("API-KEY")

	// 73301196
	message, err := sms.GetMessage(73301196)
	if err != nil {
		t.Error("error occurred ", err)
	}

	t.Log(message)
}

// TestFetchStatuses test fetch message recipients status
func TestFetchStatuses(t *testing.T) {
	sms := New("API-KEY")

	// 73301196
	statuses, paginationInfo, err := sms.FetchStatuses(73301196, ListParams{Page: 0, Limit: 10})
	if err != nil {
		t.Error("error occurred ", err)
	}

	t.Log(statuses, paginationInfo)
}

// TestFetchInbox fetch inbox test
func TestFetchInbox(t *testing.T) {
	sms := New("API-KEY")

	messages, paginationInfo, err := sms.FetchInbox(ListParams{Page: 1, Limit: 10})
	if err != nil {
		t.Error("error occurred ", err)
	}

	t.Log(messages, paginationInfo)
}

// TestCreatePattern test create pattern
func TestCreatePattern(t *testing.T) {
	sms := New("API-KEY")

	pattern, err := sms.CreatePattern("%name% is awesome", true)
	if err != nil {
		t.Error("error occurred ", err)
	}

	t.Log(pattern)
}

// TestSendPattern test send with pattern
func TestSendPattern(t *testing.T) {
	sms := New("API-KEY")
	patternValues := map[string]string{
		"name": "Mediana",
	}

	bulkID, err := sms.SendPattern(
		"t2cfmnyo0c",   // pattern code
		"+9810001",     // originator
		"98912xxxxxxx", // recipient
		patternValues,  // pattern values
	)
	if err != nil {
		t.Error("error occurred ", err)
	}

	t.Log(bulkID)
}

func TestGetCredit(t *testing.T) {
	sms := New("fyhwuRWBn6cf-6eMO90JjDmIrq2KarEnmva1Vs-TcZE=")

	credit, err := sms.GetCredit()
	if err != nil {
		t.Error("error occurred ", err)
	}

	t.Log(credit)
}
