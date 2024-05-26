package internal

import (
	"fmt"
	"net/url"
	"strings"
)

func GetEmbedYTURL(videoURL string) (string, error) {
	parsedURL, err := url.Parse(videoURL)
	if err != nil {
		return "", err
	}
	var videoID string
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
	// Include optional parameters if available
	if si := queryParams.Get("si"); si != "" {
		embedURL = fmt.Sprintf("%s?si=%s", embedURL, si)
	}
	if t := queryParams.Get("t"); t != "" {
		embedURL = fmt.Sprintf("%s&t=%s", embedURL, t)
	}

	return embedURL, nil
}
