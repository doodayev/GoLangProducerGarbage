package service

import (
	"bytes"
	"encoding/json"
	v1 "go-producer-consumer-v1-main/Producer/contracts"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ProducerService struct {
	Collector Collector
	Router    *mux.Router
	StopChan  chan bool
	Consumer  string
}

func (prd *ProducerService) Initialize() {
	prd.Collector = NewCollector()
	http.HandleFunc("/producer/start", prd.InitializeDispatcher)
	http.HandleFunc("/producer/stop", prd.StopProducer)
	http.HandleFunc("/tasks/produce", prd.Collector.RequestCollector)
}

// Function to capture the task-requests in buffered channels and produce to Consumer
func (prd *ProducerService) InitializeDispatcher(_ http.ResponseWriter, _ *http.Request) {
	log.Print("Producer Started")
	go prd.StartProducer()
}
//Implement Start/Stop producer
/*StartProducer should listen to TaskChan and StopChan. If a task is received on TaskChan, send the task data to the Consumer
* by using prd.ConsumerClient(). If a signal is received on StopChan, break out of the loop.
*/
func (prd *ProducerService) StartProducer() {

}
//StopProducer creates an anonymous function on a goroutine that sends a boolean value true through prd.StopChan.
func (prd *ProducerService) StopProducer(_ http.ResponseWriter, _ *http.Request) {

}

//Encodes the task data and does an HTTP POST request on the consumer endpoint, allowing the consumer to GET the task data
//The consumer will be listening on /tasks/consume and will create a task with the POST data
func (prd *ProducerService) ConsumerClient(task v1.Task) (string, error) {
	client := http.Client{}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(task)
	consumer := prd.Consumer
	consumerEndPoint := "/tasks/consume"
	request, err := http.NewRequest("POST", consumer+consumerEndPoint, bytes.NewBuffer(reqBodyBytes.Bytes()))
	if err != nil {
		log.Fatal("Unable to POST task to consumer")
		return "", err
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Unexpected response from consumer")
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Unexpected body from consumer")
		return "", err
	}
	return string(body), nil
}
