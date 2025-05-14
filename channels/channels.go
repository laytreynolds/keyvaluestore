// Package channels provides request channels for concurrent access of the Key Value Store
package channels

import (
	"kvstore/store"
)

var (
	GetChannel    = make(chan Request) // Create unbuffered GET channel
	AddChannel    = make(chan Request)
	GetAllChannel = make(chan Request)
	ExistChannel  = make(chan Request)
	CountChannel  = make(chan Request)
	ClearChannel  = make(chan Request)
	DeleteChannel = make(chan Request)
	UpdateChannel = make(chan Request)
	UpsertChannel = make(chan Request)
)

type Request struct {
	Key      string
	Value    []byte
	Response chan Response
}

type Response struct {
	Value any
	Error error
}

func Requests() {
	for {
		select {
		case req := <-GetChannel:
			value, err := store.Store.Get(req.Key)
			req.Response <- Response{value, err}
			close(req.Response)
		case req := <-AddChannel:
			value, err := store.Store.Add(req.Key, req.Value)
			req.Response <- Response{value, err}
			close(req.Response)
		case req := <-GetAllChannel:
			value, err := store.Store.GetAll()
			req.Response <- Response{value, err}
			close(req.Response)
		case req := <-ExistChannel:
			value, err := store.Store.Exists(req.Key)
			req.Response <- Response{value, err}
			close(req.Response)
		case req := <-CountChannel:
			value, err := store.Store.Count()
			req.Response <- Response{value, err}
			close(req.Response)
		case req := <-ClearChannel:
			value, err := store.Store.Clear()
			req.Response <- Response{value, err}
			close(req.Response)
		case req := <-DeleteChannel:
			err := store.Store.Delete(req.Key)
			req.Response <- Response{nil, err}
			close(req.Response)
		case req := <-UpdateChannel:
			value, err := store.Store.Update(req.Key, req.Value)
			req.Response <- Response{value, err}
			close(req.Response)
		case req := <-UpsertChannel:
			value, err := store.Store.Upsert(req.Key, req.Value)
			req.Response <- Response{value, err}
			close(req.Response)
		}
	}
}

func GetRequest(key string) (response Response) {
	responseCh := make(chan Response)
	GetChannel <- Request{Key: key, Response: responseCh}
	response = <-responseCh
	return response
}

func AddRequest(key string, value []byte) (response Response) {
	responseCh := make(chan Response)
	AddChannel <- Request{Key: key, Value: value, Response: responseCh}
	response = <-responseCh
	return response
}

func GetAllRequest() (response Response) {
	responseCh := make(chan Response)
	GetAllChannel <- Request{Response: responseCh}
	response = <-responseCh
	return response
}
func ExistsRequest(key string) (response Response) {
	responseCh := make(chan Response)
	ExistChannel <- Request{Key: key, Response: responseCh}
	response = <-responseCh
	return response
}

func CountRequest() (response Response) {
	responseCh := make(chan Response)
	CountChannel <- Request{Response: responseCh}
	response = <-responseCh
	return response
}

func ClearRequest() (response Response) {
	responseCh := make(chan Response)
	ClearChannel <- Request{Response: responseCh}
	response = <-responseCh
	return response
}

func DeleteRequest(key string) (response Response) {
	responseCh := make(chan Response)
	DeleteChannel <- Request{Key: key, Response: responseCh}
	response = <-responseCh
	return response
}

func UpdateRequest(key string, value []byte) (response Response) {
	responseCh := make(chan Response)
	UpdateChannel <- Request{Key: key, Value: value, Response: responseCh}
	response = <-responseCh
	return response
}

func UpsertRequest(key string, value []byte) (response Response) {
	responseCh := make(chan Response)
	UpsertChannel <- Request{Key: key, Value: value, Response: responseCh}
	response = <-responseCh
	return response
}
