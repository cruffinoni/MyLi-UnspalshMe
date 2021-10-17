package unsplash

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/application/args"
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
		return Api{}, errors.New("missing unsplash api access key")
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

func (api Api) GetImages(args args.ProgArg) (SearchQueryModel, error) {
	reqUrl := fmt.Sprintf("https://api.unsplash.com/search/photos?query=%v&page=%v", args.Query, args.Page)
	request, err := api.prepareRequest(reqUrl)
	if err != nil {
		return SearchQueryModel{}, err
	}
	client := &http.Client{}
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		return SearchQueryModel{}, err
	}
	if response.StatusCode != http.StatusOK {
		return SearchQueryModel{}, errors.New(fmt.Sprintf("status code should be %v but is %v instead", http.StatusOK, response.StatusCode))
	}
	var body []byte
	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return SearchQueryModel{}, err
	}

	var data SearchQueryModel
	if err = json.Unmarshal(body, &data); err != nil {
		return SearchQueryModel{}, err
	}
	return data, nil
}
