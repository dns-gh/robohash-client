package robohashclient

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// TODO:
// - implement robohash algorithm in Go and add more data sets
// - extend the concept to landscapes ?

const (
	robohashBaseURL = "https://robohash.org/"
)

// RobohashClient represents the web Client.
// Note: if you want to change the robohash options
// you have to create another one with the desired options.
type RobohashClient struct {
	width  int
	height int
	set    string
	bgset  string
}

func makeSet(set int) string {
	if 1 <= set && set <= 3 {
		return strconv.Itoa(set)
	}
	return ""
}

func makeBgset(bgset int) string {
	if 1 <= bgset && bgset <= 2 {
		return strconv.Itoa(bgset)
	}
	return ""
}

// MakeRobohashClient creates a web client to make http request
// to the robohash API.
func MakeRobohashClient(width, height, set, bgset int) *RobohashClient {
	log.Println("[robohash] making robohash client")
	return &RobohashClient{
		width:  width,
		height: height,
		set:    makeSet(set),
		bgset:  makeSet(bgset),
	}
}

func loadImage(uri string) (string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error request status: %s != 200", resp.Status)
	}
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (r *RobohashClient) makeURI(query string) string {
	// ex: https://robohash.org/dns-gh?size=300x300&set=set3&bgset=bg3
	uri := robohashBaseURL + url.QueryEscape(query) + "?"
	uri += "size=" + url.QueryEscape(strconv.Itoa(r.width)) + "x" + url.QueryEscape(strconv.Itoa(r.height))
	if len(r.set) != 0 {
		uri += "&set=" + url.QueryEscape(r.set)
	}
	if len(r.bgset) != 0 {
		uri += "&bgset=" + url.QueryEscape(r.bgset)
	}
	return uri
}

// Fetch fetches the robohash image using a uri
// build with the RobohashClient parameters.
func (r *RobohashClient) Fetch(query string) (string, error) {
	return loadImage(r.makeURI(query))
}
