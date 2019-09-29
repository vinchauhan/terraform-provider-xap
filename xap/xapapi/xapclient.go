package xapapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type XapClient struct {
	BaseUrl    *url.URL
	UserAgent  string
	httpClient *http.Client
}

func newXapClient(url string) {

}

func (c *XapClient) CreateSpace(name string, partitions int, backups bool, requiresIsolation bool) error {
	rel := &url.URL{Path: "/spaces"}
	u := c.BaseUrl.ResolveReference(rel)
	data := url.Values{}
	data.Set("name", name)
	data.Set("partitions", string(partitions))
	data.Set("backups", strconv.FormatBool(backups))
	data.Set("requiresIsolation", strconv.FormatBool(requiresIsolation))

	req, err := http.NewRequest("POST", u.String(), bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(f))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
}
