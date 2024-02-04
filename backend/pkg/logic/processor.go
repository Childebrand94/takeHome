package logic

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/url"
	"os"
	"regexp"

	"github.com/Childebrand94/takeHomePhoneNumber/pkg/models"
	"mvdan.cc/xurls/v2"
)

// Open data set and create a hash map of all prefix's where the prefix
// is the key
func createMap() (map[string]models.PrefixInfo, error) {
	jsonFile, err := os.Open("pkg/data/data.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open data.json: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read data.json: %w", err)
	}

	var tempData []struct {
		Prefix int `json:"prefix"`
		models.PrefixInfo
	}

	if err = json.Unmarshal(byteValue, &tempData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	prefixMap := make(map[string]models.PrefixInfo)
	for _, item := range tempData {
		prefixMap[fmt.Sprint(item.Prefix)] = item.PrefixInfo
	}

	return prefixMap, nil
}

// Create the longest prefix by iterating over the given phone number creating a
// string that matches the longest prefix
func processPhoneNumber(phoneNumber string) (models.PrefixInfo, error) {
	dataMap, err := createMap()
	if err != nil {
		return models.PrefixInfo{}, fmt.Errorf("error creating data map: %w", err)
	}

	var longestPrefix string

	for i := 1; i <= len(phoneNumber); i++ {
		prefix := phoneNumber[:i]

		if _, ok := dataMap[prefix]; ok {
			longestPrefix = prefix
		}
	}

	return dataMap[longestPrefix], nil
}

// initialize once not on every function call to process message
var urlRegex = regexp.MustCompile(xurls.Strict().String())

// Using the xurls library to find and replace all URLs, with URLs wrapped in
// <a> tags for the front end to render.
// The found URLS are sanitized with Go's url.parse and Go's html.EscapeString
// for added security
func processMessage(message string) (string, error) {

	result := urlRegex.ReplaceAllStringFunc(message, func(match string) string {
		parsedURL, err := url.Parse(match)
		if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
			return match
		}
		safeURL := html.EscapeString(match)
		return fmt.Sprintf(`<a href='%s'>%s</a>`, safeURL, safeURL)
	})

	return result, nil
}

func ProcessData(payload models.Query) (models.Resp, error) {

	prefix_info, err := processPhoneNumber(payload.Phone_number)
	if err != nil {
		return models.Resp{},
			fmt.Errorf("Error has occurred processing phone number: %w", err)
	}

	formatted_message, err := processMessage(payload.Message)
	if err != nil {
		return models.Resp{},
			fmt.Errorf("Error has occurred processing message: %w", err)
	}

	var resp models.Resp
	resp.PrefixInfo = prefix_info
	resp.Message = formatted_message

	return resp, nil

}
