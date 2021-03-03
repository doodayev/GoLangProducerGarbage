package service

import (
	"encoding/json"
	"errors"
	v1 "go-producer-consumer-v1-main/Producer/contracts"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	ErrIncorrectTaskFields = "task fields cannot be empty"
	ErrInvalidTimeFormat   = "invalid time format in request"
	ErrInvalidPeriodicity  = "invalid periodicity value"
)

type Collector struct {
	logger *log.Logger
}

func NewCollector() Collector {
	return Collector{}
}

// A buffered channel that captures tasks to be produced to Consumer
// Implement buffered channel of 10 here


// Function to capture the task-requests in buffered channels and produce to Consumer
func (ctr *Collector) RequestCollector(_ http.ResponseWriter, req *http.Request) {
	taskRequest := v1.Task{}
	requestBody := extractRequestBody(req)

	err := json.Unmarshal(requestBody, &taskRequest)
	if err != nil {
		ctr.logger.Fatal(
			"service", "Producer",
			"method", "ValidateRequest",
			"error", err)
		return
	}

	err = ctr.ValidateRequest(taskRequest)
	if err != nil {
		return
	}

	//Push the taskRequest to the buffered channel here

	return
}

func extractRequestBody(req *http.Request) []byte {
	body := ""
	if req.Body != nil {
		bytes, err := ioutil.ReadAll(req.Body)
		if err == nil {
			body = string(bytes)
		}
	}
	return []byte(body)
}

// Function to validate the request from the client
func (ctr *Collector) ValidateRequest(task v1.Task) error {
	if task.TaskName == "" || task.TaskType == "" || task.LastUpdateTime == "" ||
		task.ScheduledTime == "" {
		ctr.logger.Fatal(
			"service", "Producer",
			"method", "ValidateRequest",
			"error", "task fields cannot be empty")
		return errors.New(ErrIncorrectTaskFields)
	}

	if task.Periodicity < 0 {
		ctr.logger.Fatal(
			"service", "Producer",
			"method", "ValidateRequest",
			"error", "periodicity should be positive")
		return errors.New(ErrInvalidPeriodicity)
	}

	return nil
}

func (ctr *Collector) ValidateRequestTimeFields(req string, timerType string) error {
	_, err := time.Parse(time.RFC3339, req)
	if err != nil {
		ctr.logger.Fatal(
			"service", "Producer",
			"method", "ValidateRequest",
			"error", "Invalid format", timerType)
		return errors.New(ErrInvalidTimeFormat)
	}
	return nil
}
