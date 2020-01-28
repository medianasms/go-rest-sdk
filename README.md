# Mediana SMS api SDK

This repository contains open source Go client for `mediana_sms` api. Documentation can be found at: <http://docs.medianasms.com>.

[![Build Status](https://travis-ci.org/medianasms/go-rest-sdk.svg?branch=master)](https://travis-ci.org/medianasms/go-rest-sdk) [![GoDoc](https://godoc.org/github.com/medianasms/go-rest-sdk?status.svg)](https://godoc.org/github.com/medianasms/go-rest-sdk)

## Installation

If you are using go modules, just install it with `go mod install` or `go build .`, Otherwise you can use `go get ./...`

```bash
go get github.com/medianasms/go-rest-sdk
```

## Examples

After installing medianasms sdk with above methods, you can import it in your project like this:

```go
import "github.com/medianasms/go-rest-sdk"
```

For using sdk, after importing, you have to create a client instance that gives you available methods on API

```go
// you api key that generated from panel
apiKey := "api-key"

// create client instance
sms := medianasms.New(apiKey)

...
```

### Credit check

```go
// return float64 type credit amount
credit, err := sms.GetCredit()
if err != nil {
    t.Error("error occurred ", err)
}
```

### Send one to many

For sending sms, obviously you need `originator` number, `recipients` and `message`.

```go
bulkID, err := sms.Send("+9810001", []string{"98912xxxxxxx"}, "mediana is awesome")
if err != nil {
    t.Error("error occurred ", err)
}
```

If send is successful, a unique tracking code returned and you can track your message status with that.

### Get message summery

```go
bulkID := "message-tracking-code"

message, err := sms.GetMessage(bulkID)
if err != nil {
    t.Error("error occurred ", err)
}

fmt.Println(message.Status) // get message status
fmt.Println(message.Cost)   // get message cost
fmt.Println(message.Payack) // get message payback
```

### Get message delivery statuses

```go
bulkID := "message-tracking-code"
// pagination params for defining fetch size and offset
paginationParams := medianasms.ListParams{Page: 0, Limit: 10}

statuses, paginationInfo, err := sms.FetchStatuses(bulkID, paginationParams)
if err != nil {
    t.Error("error occurred ", err)
}

// you can loop in messages statuses list
for _, status := range statuses {
    fmt.Printf("Recipient: %s, Status: %s", status.Recipient, status.Status)
}

fmt.Println("Total: ", paginationInfo.Total)
```

### Inbox fetch

fetch inbox messages

```go
paginationParams := medianasms.ListParams{Page: 0, Limit: 10}

messages, paginationInfo, err := sms.FetchInbox(paginationParams)
if err != nil {
    t.Error("error occurred ", err)
}
```

### Pattern create

For sending messages with predefined pattern(e.g. verification codes, ...), you hav to create a pattern. a pattern at least have a parameter. parameters defined with `%param_name%`.

```go
pattern, err := sms.CreatePattern("%name% is awesome", false)
if err != nil {
    t.Error("error occurred ", err)
}
```

### Send with pattern

```go
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
```

### Error checking

```go
_, err := sms.Send("9810001", []string{"98912xxxxx"}, "mediana is awesome")
if err != nil {
    // check that error is type of medianasms standard error
    if e, ok := err.(Error); ok {
        // after casting, you have access to error code and error message
        switch e.Code {
        // its special type of error, occurred when POST form data validation failed 
        case ErrUnprocessableEntity:
            // for accessing field level errors you have to cast it to FieldErrors type
            fieldErrors := e.Message.(FieldErrs)
            for field, fieldError := range fieldErrors {
                t.Log(field, fieldError)
            }
        // in other error types, e.Message is string
        default:
            errMsg := e.Message.(string)
            t.Log(errMsg)
        }
    }
}
```
