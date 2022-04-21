package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	// "strconv"
)

const BaseUrl = "https://api.notion.com"
const V1 = "v1"

type NotionClient struct {
	BaseURL       string
	ApiVersion    string
	NotionVersion string
	apiKey        string
	HttpClient    *http.Client
}

func CreateNotionClient(apiKey string) NotionClient {
	return NotionClient{
		BaseURL:    BaseUrl,
		ApiVersion: V1,
		apiKey:     apiKey,
		NotionVersion: "2022-02-22",
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (nc *NotionClient) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Notion-Version", nc.NotionVersion)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", nc.apiKey))

	resp, err := nc.HttpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return []byte{}, fmt.Errorf("error in request: %v, status code: %v, body: %v", req.URL, resp.StatusCode, string(body))
	}

	return body, nil
}

func (nc *NotionClient) QueryDatabase(ctx context.Context, query DatabaseQuery, database string) (DatabasePages, error) {
	jsn, err := json.Marshal(query)
	if err != nil {
		return DatabasePages{}, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/%v/databases/%v/query", nc.BaseURL, nc.ApiVersion, database), bytes.NewBuffer(jsn))
	if err != nil {
		return DatabasePages{}, err
	}

	req = req.WithContext(ctx)

	body, err := nc.doRequest(req)
	if err != nil {
		return DatabasePages{}, err
	}

	var pages DatabasePages
	err = json.Unmarshal(body, &pages)
	if err != nil {
		fmt.Println(err)
		return DatabasePages{}, err
	}
	return pages, nil
}

func (nc *NotionClient) UpdatePage(ctx context.Context, properties UpdatePage, pageId string) (Page, error) {
	jsn, err := json.Marshal(properties)
	if err != nil {
		return Page{}, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%v/%v/pages/%v", nc.BaseURL, nc.ApiVersion, pageId), bytes.NewBuffer(jsn))
	if err != nil {
		return Page{}, err
	}

	req = req.WithContext(ctx)

	resp, err := nc.doRequest(req)
	if err != nil {
		return Page{}, err
	}

	var page Page
	err = json.Unmarshal(resp, &page)
	if err != nil {
		return Page{}, err
	}

	return page, nil

}
