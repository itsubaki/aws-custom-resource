package main

import "encoding/json"

type Request struct {
	RequestType           string          `json:"RequestType"` // Create, Update, Delete
	ResponseURL           string          `json:"ResponseURL"`
	StackId               string          `json:"StackId"`
	RequestId             string          `json:"RequestId"`
	ResourceType          string          `json:"ResourceType"`
	LogicalResourceId     string          `json:"LogicalResourceId"`
	PhysicalResourceId    string          `json:"PhysicalResourceId"`
	ResourceProperties    json.RawMessage `json:"ResourceProperties"`
	OldResourceProperties json.RawMessage `json:"OldResourceProperties"`
	ServiceToken          string          `json:"ServiceToken"`
}
