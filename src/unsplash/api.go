package unsplash

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/application/args"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/unsplash/models"
	"io/ioutil"
	"net/http"
	"os"
)

type Api struct {
	apiKey string
}

func New() (Api, error) {
	key := os.Getenv("UNSPLASH_API_ACCESS_KEY")
	if len(key) == 0 {
		return Api{}, errors.New("missing unsplash api access key environment variable")
	}
	return Api{
		apiKey: key,
	}, nil
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

func formatErrorMsg(body []byte) string {
	var data models.Error
	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Sprintf("can't unmarshal body: %v", err)
	}
	var message string
	for i, e := range data.Errors {
		message += e
		if i != len(data.Errors)-1 {
			message += " "
		}
	}
	return fmt.Sprintf("%v", message)
}

func (api Api) GetImages(args args.ProgArg) (models.SearchImageQuery, error) {
	reqUrl := fmt.Sprintf("https://api.unsplash.com/search/photos?query=%v&page=%v", args.Query, args.Page)
	request, err := api.prepareRequest(reqUrl)
	if err != nil {
		return models.SearchImageQuery{}, err
	}
	client := &http.Client{}
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		return models.SearchImageQuery{}, err
	}
	var body []byte
	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return models.SearchImageQuery{}, err
	}
	if response.StatusCode != http.StatusOK {
		return models.SearchImageQuery{}, errors.New(fmt.Sprintf("invalid status code (%v), error: '%v'", response.StatusCode, formatErrorMsg(body)))
	}

	var data models.SearchImageQuery
	if err = json.Unmarshal(body, &data); err != nil {
		return models.SearchImageQuery{}, err
	}
	return data, nil
}
