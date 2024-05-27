package internal

import (
	"fmt"
	"net/url"
	"strings"
)

func GetEmbedYTURL(videoURL string) (string, error) {
	parsedURL, err := url.Parse(videoURL)
	var videoID string

	if err != nil {
		return "", err
	}
	if !(strings.Contains(parsedURL.Host, "youtube.com") || strings.Contains(parsedURL.Host, "youtu.be")) {
		return "", fmt.Errorf("Not a youtube url")
	}
	if strings.Contains(parsedURL.Host, "youtube.com") && strings.HasPrefix(parsedURL.Path, "/embed/") {
		return videoURL, nil
	}

	queryParams := parsedURL.Query()
	if videoID = queryParams.Get("v"); videoID == "" {
		// Handle the case for youtu.be links with or without query parameters
		if strings.Contains(parsedURL.Host, "youtu.be") {
			videoID = strings.TrimPrefix(parsedURL.Path, "/")
		} else {
			return "", fmt.Errorf("invalid YouTube URL")
		}
	}

	embedURL := fmt.Sprintf("https://www.youtube.com/embed/%s", videoID)
	embedURL = strings.TrimSpace(embedURL)
	// Include optional parameters if available
	parameters := []string{}
	if si := queryParams.Get("si"); si != "" {
		parameters = append(parameters, "si="+si)
	}
	if t := queryParams.Get("t"); t != "" {
		parameters = append(parameters, "t="+t)
	}
	// enable autoplay
	if true {
		parameters = append(parameters, "autoplay=1")
	}
	if len(parameters) > 0 {
		embedURL = fmt.Sprintf("%s?%s", embedURL, strings.Join(parameters, "&"))
	}

	return embedURL, nil
}
