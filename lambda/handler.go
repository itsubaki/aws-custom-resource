package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) handle(req Request) {
	log.Println("[start]")
	log.Printf("request type=%s", req.RequestType)

	res := Response{
		StackId:            req.StackId,
		RequestId:          req.RequestId,
		LogicalResourceId:  req.LogicalResourceId,
		PhysicalResourceId: uuid.New().String(),
		Status:             "SUCCESS",
		Reason:             "",
	}
	res.Data.Value = "Hello World"
	log.Printf("created response %#v", res)

	client := &http.Client{}
	bytea, _ := json.MarshalIndent(res, "", "    ")
	body := bytes.NewReader(bytea)

	resreq, _ := http.NewRequest("PUT", req.ResponseURL, body)
	resreq.ContentLength = int64(len(string(bytea)))
	resres, err := client.Do(resreq)
	if err != nil {
		log.Printf("put response to %s: %v", req.ResponseURL, err)
		os.Exit(1)
	}
	defer resres.Body.Close()

	contents, err := ioutil.ReadAll(resres.Body)
	if err != nil {
		log.Printf("read response from %s: %v", req.ResponseURL, err)
		os.Exit(1)
	}
	log.Printf("response status=%s", resres.Status)
	for k, v := range resres.Header {
		log.Printf("response header %s=%s", k, v)
	}
	log.Printf("response content=%v", string(contents))

	log.Println("[end] amilookup")
}
