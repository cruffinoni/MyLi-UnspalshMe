package unsplash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Api struct {
	apiKey string
}

func New() Api {
	key := os.Getenv("UNSPLASH_API_ACCESS_KEY")
	if len(key) == 0 {
		log.Fatalf("Missing access key")
	}
	return Api{
		apiKey: key,
	}
}

func (api Api) prepareRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Err while creating new generic get http request: %v\n", err.Error())
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Client-ID %v", api.apiKey))
	return req, nil
}

func (api Api) GetImages(url string) (SearchQueryModel, error) {
	request, err := api.prepareRequest(url)
	if err != nil {
		return SearchQueryModel{}, err
	}
	client := &http.Client{}
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		log.Printf("Err during http request: %v\n", err)
		return SearchQueryModel{}, err
	}
	if response.StatusCode != http.StatusOK {
		log.Printf("Code is not 200: %v\n", response.StatusCode)
	}
	var body []byte
	if body, err = ioutil.ReadAll(response.Body); err != nil {
		log.Printf("Cant read all: %v\n", err)
		return SearchQueryModel{}, err
	}

	var data SearchQueryModel
	if err = json.Unmarshal(body, &data); err != nil {
		log.Printf("JSON unmarshal: %v\n", err)
		return SearchQueryModel{}, err
	}
	return data, nil
}
