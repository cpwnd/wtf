package telraam

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/wtfutil/wtf/utils"
)

func GetCamera(cameraId int) (Camera, error) {
	var cameras []Camera

	resp, err := apiRequest("cameras/" + strconv.Itoa(cameraId))
	if err != nil {
		return cameras, err
	}

	err = utils.ParseJSON(&cameras, resp.Body)
	if err != nil {
		return cameras, err
	}

	return cameras[0], nil
}

/* -------------------- Unexported Functions -------------------- */

var (
	apiEndpoint = "https://telraam-api.net/v0/"
)

func apiRequest(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", apiEndpoint+path, nil)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}
