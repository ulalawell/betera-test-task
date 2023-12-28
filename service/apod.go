package service

import (
	"context"
	"io/ioutil"
	"net/http"
	"tesk-task-betera/models"
	"tesk-task-betera/repository"
)

type Apod struct {
	rep *repository.Apod
}

func NewApodService(rep *repository.Apod) *Apod {
	return &Apod{rep: rep}
}

func (s *Apod) CreateApod(ctx context.Context, apod *models.ApodData) error {
	var apodDto = &models.ApodDto{}
	var data []byte
	var err error

	if apod.MediaType == "image" {
		data, err = downloadFile(apod.HDURL)
	} else {
		data, err = downloadFile(apod.URL)
	}

	if err != nil {
		return err
	}

	*apodDto = models.ApodDto{
		Copyright:      apod.Copyright,
		Date:           apod.Date,
		Explanation:    apod.Explanation,
		MediaType:      apod.MediaType,
		ServiceVersion: apod.ServiceVersion,
		Title:          apod.Title,
		URL:            apod.URL,
		HDURL:          apod.HDURL,
		Data:           data,
	}
	return s.rep.CreateApod(ctx, apodDto)
}

func (s *Apod) GetApods(ctx context.Context, getCouriersRequest *models.GetApodsRequest) (*models.GetApodsResponse, error) {
	return s.rep.GetApods(ctx, getCouriersRequest)
}

func (s *Apod) GetApodsByDate(ctx context.Context, getApodByDateRequest models.GetApodByDateRequest) (*models.GetApodByDateResponse, error) {
	return s.rep.GetApodsByDate(ctx, getApodByDateRequest)
}

func downloadFile(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	fileData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return fileData, nil
}
