package apodWorker

import (
	"encoding/json"
	"fmt"
	"github.com/go-co-op/gocron"
	"io/ioutil"
	"net/http"
	"tesk-task-betera/models"
	"time"
)

type ApodWorker struct {
	periodicitySeconds int
	apiKey             string
	url                string
	ch                 chan models.ApodChanelType
}

func NewApodWorker(periodicity int, apiKey string, url string, ch chan models.ApodChanelType) *ApodWorker {
	return &ApodWorker{
		periodicitySeconds: periodicity,
		apiKey:             apiKey,
		url:                url,
		ch:                 ch,
	}
}

func (worker *ApodWorker) GetApodData() error {
	s := gocron.NewScheduler(time.UTC)

	_, err := s.Every(worker.periodicitySeconds).Seconds().Do(func() {
		apodData, err := worker.getNasaInformation()
		if err != nil {
			worker.ch <- models.ApodChanelType{ApodData: nil, Error: err}
		}

		worker.ch <- models.ApodChanelType{ApodData: apodData}
	})
	if err != nil {
		return err
	}

	s.StartBlocking()

	return nil
}

func (worker *ApodWorker) getNasaInformation() (*models.ApodData, error) {
	fullURL := fmt.Sprintf("%s?api_key=%s", worker.url, worker.apiKey)

	response, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET request failed: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK status code received: %d", response.StatusCode)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	var apodData models.ApodData

	err = json.Unmarshal(responseBody, &apodData)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshal error: %s", err)
	}

	return &apodData, nil
}
