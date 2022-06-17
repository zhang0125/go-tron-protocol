package tron

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type EventData struct {
	BlockNumber           int64       `json:"block_number"`
	BlockTimestamp        int64       `json:"block_timestamp"`
	CallerContractAddress string      `json:"caller_contract_address"`
	ContractAddress       string      `json:"contract_address"`
	Event                 string      `json:"event"`
	EventIndex            int         `json:"event_index"`
	EventName             string      `json:"event_name"`
	Result                interface{} `json:"result"`
	TransactionId         string      `json:"transaction_id"`
}

type EventResponse struct {
	Data    []EventData `json:"data"`
	Success bool        `json:"success"`
	Meta    struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}

func GetContractEvents(tronGridUrl, apiKey, contract, start, end string) ([]EventData, error) {
	gridUrl := fmt.Sprintf("%v/v1/contracts/%v/events", tronGridUrl, contract)
	req, err := http.NewRequest(http.MethodGet, gridUrl, nil)
	if err != nil {
		return nil, err
	}
	params := make(url.Values)
	params.Add("min_block_timestamp", start)
	params.Add("max_block_timestamp", end)
	params.Add("limit", "200")
	params.Add("order_by", "block_timestamp,asc")
	//params.Add("only_confirmed", "true")

	req.URL.RawQuery = params.Encode()
	req.Header.Add("TRON-PRO-API-KEY", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Printf("Response failed with status code: %d and\nbody: %s\n\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	var event EventResponse
	if err = json.Unmarshal(body, &event); err != nil {
		return nil, err
	}
	if !event.Success {
		return nil, fmt.Errorf("fail to get contract event")
	}

	return event.Data, nil
}
