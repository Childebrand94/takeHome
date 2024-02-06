package logic

import (
	"embed"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/url"
	"regexp"

	"github.com/Childebrand94/takeHomePhoneNumber/pkg/models"
	"mvdan.cc/xurls/v2"
)

// Embed the data.json file

//go:embed data.json
var dataFS embed.FS

var globalDataMap map[string]models.PrefixInfo

// Initialize the URL regex
var urlRegex = regexp.MustCompile(xurls.Strict().String())

// Initialize data to a map once on program load
func init() {
	data, err := dataFS.ReadFile("data.json")
	if err != nil {
		log.Fatalf("Error reading embedded data.json: %v", err)
	}

	var tempData []models.PrefixInfo
	if err := json.Unmarshal(data, &tempData); err != nil {
		log.Fatalf("Error unmarshaling data.json: %v", err)
	}

	globalDataMap, err = createMap(tempData)
	if err != nil {
		log.Fatalf("Error creating data map: %v", err)
	}
}

// Convert the data.json into a hash map due to constant look ups for prefix match
func createMap(data []models.PrefixInfo) (map[string]models.PrefixInfo, error) {
	prefixMap := make(map[string]models.PrefixInfo)
	for _, item := range data {
		prefixMap[fmt.Sprint(item.Prefix)] = item
	}

	return prefixMap, nil
}

// Create the longest prefix by iterating over the given phone number creating a
// string that matches the longest prefix
func processPhoneNumber(phoneNumber string) (string, error) {

	var longestPrefix string

	for i := 1; i <= len(phoneNumber); i++ {
		prefix := phoneNumber[:i]

		if _, ok := globalDataMap[prefix]; ok {
			longestPrefix = prefix
		}
	}

	return longestPrefix, nil
}

// Return empty Prefix Info struct on non matches
func getPrefixInfo(prefix string) models.PrefixInfo {
	if prefixInfo, ok := globalDataMap[prefix]; ok {
		return prefixInfo
	}
	return models.PrefixInfo{}
}

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

	prefix, err := processPhoneNumber(payload.Phone_number)
	if err != nil {
		return models.Resp{},
			fmt.Errorf("Error has occurred finding longest prefix: %w", err)
	}

	prefixInfo := getPrefixInfo(prefix)

	formatted_message, err := processMessage(payload.Message)
	if err != nil {
		return models.Resp{},
			fmt.Errorf("Error has occurred processing message: %w", err)
	}

	var resp models.Resp
	resp.PrefixInfo = prefixInfo
	resp.Message = formatted_message

	return resp, nil

}
