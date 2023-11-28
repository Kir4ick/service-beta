package client

import (
	"beta/internal/request"
	"bytes"
	"encoding/json"
	"net/http"
)

const voteStateUrl = "/voting-stats"

type GammaClient struct {
	gammaUrl    string
	contentType string
}

func NewGammaClient(gammaUrl string) *GammaClient {
	return &GammaClient{gammaUrl: gammaUrl + voteStateUrl}
}

func (client *GammaClient) Send(requestData *request.VoteGammaRequest) error {
	jsonData, err := json.Marshal(requestData)
	dataForSend := bytes.NewReader(jsonData)

	if err != nil {
		return err
	}

	_, err = http.Post(client.gammaUrl, client.contentType, dataForSend)

	if err != nil {
		return err
	}

	return nil
}
