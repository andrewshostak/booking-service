package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type launchpadRepository struct {
	apiUrl     string
	httpClient *http.Client
}

func NewLaunchpadRepository(apiUrl string, httpClient *http.Client) LaunchpadRepository {
	return &launchpadRepository{apiUrl: apiUrl, httpClient: httpClient}
}

func (r *launchpadRepository) IsLaunchpadAvailable(launchpadId string, launchDate time.Time) (bool, error) {
	query := LaunchesQuery{
		Query: map[string]interface{}{
			"launchpad": launchpadId,
			"upcoming":  true,
			"date_utc": map[string]string{
				"$gte": launchDate.Format(time.RFC3339),
				"$lte": launchDate.Add(time.Hour * 24).Format(time.RFC3339),
			},
		},
	}

	response, err := r.sendLaunchesListRequest(query)
	if err != nil {
		return false, err
	}

	return response.TotalDocs == 0, nil
}

func (r *launchpadRepository) sendLaunchesListRequest(query LaunchesQuery) (*LaunchesResponse, error) {
	jsonBody, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v5/launches/query", r.apiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var launchesResponse LaunchesResponse
	err = json.Unmarshal(responseBody, &launchesResponse)
	if err != nil {
		return nil, err
	}

	return &launchesResponse, nil
}
