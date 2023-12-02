package services

import (
	"beta/pkg/regulation"
	"time"
)

func (s *Service) ClearInfoForRequests(requestRegulator *regulation.Request) {

	ticker := time.NewTicker(time.Second * 1)
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				requestRegulator.ClearInfo(time.Now().Add(-1 * time.Second))
			case <-stop:
				return
			}
		}
	}()
}
