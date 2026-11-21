package services

import "log"

type ObsService struct{}

func NewObsService() *ObsService {
	return &ObsService{}
}

func (o *ObsService) SendMessage(eventName, msg string) error {
	log.Println("Sending message: ", eventName, msg)
	return nil
}
