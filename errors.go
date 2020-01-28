package medianasms

import (
	"encoding/json"
	"errors"
)

// ResponseCode api response code error type
type ResponseCode string

const (
	// ErrCredential error when executing repository query
	ErrCredential ResponseCode = "10001"
	// ErrMessageBodyIsEmpty message body is empty
	ErrMessageBodyIsEmpty ResponseCode = "10002"
	// ErrUserLimitted user is limited
	ErrUserLimitted ResponseCode = "10003"
	// ErrLineNotAssignedToYou line not assigned to you
	ErrLineNotAssignedToYou ResponseCode = "10004"
	// ErrRecipientsEmpty recipients is empty
	ErrRecipientsEmpty ResponseCode = "10005"
	// ErrCreditNotEnough credit not enough
	ErrCreditNotEnough ResponseCode = "10006"
	// ErrLineNotProfitForBulkSend line not profit for bulk send
	ErrLineNotProfitForBulkSend ResponseCode = "10007"
	// ErrLineDeactiveTemp line deactivated temporally
	ErrLineDeactiveTemp ResponseCode = "10008"
	// ErrMaximumRecipientExceeded maximum recipients number exceeded
	ErrMaximumRecipientExceeded ResponseCode = "10009"
	// ErrOperatorOffline operator is offline
	ErrOperatorOffline ResponseCode = "10010"
	// ErrNoPricing pricing not defined for user
	ErrNoPricing ResponseCode = "10011"
	// ErrTicketIsInvalid ticket is invalid
	ErrTicketIsInvalid ResponseCode = "10012"
	// ErrAccessDenied access denied
	ErrAccessDenied ResponseCode = "10013"
	// ErrPatternIsInvalid pattern is invalid
	ErrPatternIsInvalid ResponseCode = "10014"
	// ErrPatternParamettersInvalid pattern parameters is invalid
	ErrPatternParamettersInvalid ResponseCode = "10015"
	// ErrPatternIsInactive pattern is inactive
	ErrPatternIsInactive ResponseCode = "10016"
	// ErrPatternRecipientInvalid pattern recipient invalid
	ErrPatternRecipientInvalid ResponseCode = "10017"
	// ErrPatternUnAuthorizedSend unauthorized send with pattern
	ErrPatternUnAuthorizedSend ResponseCode = "10018"
	// ErrItsTimeToSleep send time is 8-23
	ErrItsTimeToSleep ResponseCode = "10019"
	// ErrCreditCardNotProvided credit card not provided
	ErrCreditCardNotProvided ResponseCode = "10020"
	// ErrDocumentsNotApproved one/all of users documents not approved
	ErrDocumentsNotApproved ResponseCode = "10021"
	// ErrInternal internal error
	ErrInternal ResponseCode = "10022"
	// ErrEntityNotFound internal error
	ErrEntityNotFound ResponseCode = "10023"
	// ErrForbidden internal error
	ErrForbidden ResponseCode = "10024"
	// ErrUnprocessableEntity inputs have some problems
	ErrUnprocessableEntity ResponseCode = "422"
	// ErrUnauthorized unauthorized
	ErrUnauthorized ResponseCode = "1401"
	// ErrKeyNotValid api key is not valid
	ErrKeyNotValid ResponseCode = "1402"
	// ErrKeyRevoked api key revoked
	ErrKeyRevoked ResponseCode = "1403"
)

// Error general service error
type Error struct {
	Code    ResponseCode
	Message interface{}
}

// FieldErr input field level error
type FieldErr struct {
	Code string `json:"code"`
	Err  string `json:"err"`
}

// FieldErrs input field level errors
type FieldErrs map[string][]FieldErr

// Error implement error interface
func (e Error) Error() string {
	switch e.Message.(type) {
	case string:
		return e.Message.(string)
	case FieldErrs:
		m, _ := json.Marshal(e.Message)
		return string(m)
	}

	return string(e.Code)
}

// fieldErrsRes field errors response
type fieldErrsRes struct {
	Errors FieldErrs `json:"error"`
}

// defaultErrsRes default template for errors body
type defaultErrsRes struct {
	Errors string `json:"error"`
}

// ParseErrors ...
func ParseErrors(res *BaseResponse) error {
	var err error
	e := Error{Code: res.Code}

	// TODO: improve logic
	switch res.Code {
	case ErrUnprocessableEntity:
		message := fieldErrsRes{}
		err = json.Unmarshal(res.Data, &message)
		e.Message = message.Errors
	default:
		message := defaultErrsRes{}
		err = json.Unmarshal(res.Data, &message)
		e.Message = message.Errors
	}

	if err != nil {
		return errors.New("cant marshal errors into standard template")
	}

	return e
}
