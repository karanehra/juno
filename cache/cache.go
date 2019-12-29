package cache

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

const vapourPort int = 3009

//CacheClient represents the the cache interaction layer
var CacheClient *Client

func init() {
	CacheClient = &Client{
		Port:    vapourPort,
		BaseURL: "http://localhost",
	}
	err := CacheClient.Create()
	if err != nil {
		log.Fatal("Cannot connect to caching service")
	}
}

//Client defines the getter and setter cache methods
type Client struct {
	Port        int
	BaseURL     string
	StatusCheck string
	GetKeyURL   string
	SetKeyURL   string
	Instance    *http.Client
}

//Entry defines a key value struct
type Entry struct {
	Key   string
	Value interface{}
}

//Create checks for a connection with the vapour server
func (client *Client) Create() error {
	client.createURLs()
	client.Instance = &http.Client{Timeout: 1 * time.Second}
	req, _ := http.NewRequest("GET", client.StatusCheck, nil)
	res, err := client.Instance.Do(req)
	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return errors.New("Connection failed")
	}
	fmt.Println("Connected to KV store...")
	return nil
}

func (client *Client) createURLs() {
	client.StatusCheck = fmt.Sprintf("%s:%d/status", client.BaseURL, client.Port)
	client.GetKeyURL = fmt.Sprintf("%s:%d/get/", client.BaseURL, client.Port)
	client.SetKeyURL = fmt.Sprintf("%s:%d/set", client.BaseURL, client.Port)
}

func (client *Client) createGetURL(key string) string {
	return client.GetKeyURL + key
}

//Get returns a key value from the cache
func (client *Client) Get(key string) (interface{}, error) {
	req, _ := http.NewRequest("GET", client.createGetURL(key), nil)
	req.Header.Set("content-type", "application/json")
	res, err := client.Instance.Do(req)
	var body interface{} = map[string]string{}
	json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		{
			return nil, errors.New("Cache: GET failed")
		}
	}
	return body, nil
}

//Set puts a key value to the cache
func (client *Client) Set(key string, value interface{}) error {
	payload, err := json.Marshal(map[string]interface{}{
		"key":    key,
		"value":  value,
		"expiry": 0,
	})
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("POST", client.SetKeyURL, bytes.NewReader(payload))
	req.Header.Set("content-type", "application/json")
	res, err := client.Instance.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode == 200 {
		return nil
	}
	return nil
}
