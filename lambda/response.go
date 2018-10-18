package main

type Response struct {
	StackId            string `json:"StackId"`
	RequestId          string `json:"RequestId"`
	LogicalResourceId  string `json:"LogicalResourceId"`
	PhysicalResourceId string `json:"PhysicalResourceId"`
	Status             string `json:"Status"`
	Reason             string `json:"Reason"`
	Data               struct {
		Value string `json:"Value"`
	}
}
