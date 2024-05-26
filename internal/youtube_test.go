package internal

import (
	"strings"
	"testing"
)

func TestGetEmbedYTURL(t *testing.T) {
	testUrls := []string{
		"https://www.youtube.com/watch?v=Aavt_JRhDZA",
		"https://youtu.be/Aavt_JRhDZA",
		"https://youtu.be/Aavt_JRhDZA?si=PhsiVtemJm3EEE-v",
		"https://youtu.be/Aavt_JRhDZA?si=PhsiVtemJm3EEE-v&t=818s",
	}
	t.Run("Test with standart Url", func(t *testing.T) {
		res, err := GetEmbedYTURL(testUrls[0])
		println(res)
		if err != nil && strings.Contains(res, "embed") {
			t.Errorf("GetEmbedYTURL() error = %v", err)
			return
		}
	})

	t.Run("Test with shortened Url", func(t *testing.T) {
		res, err := GetEmbedYTURL(testUrls[1])
		println(res)
		if err != nil && strings.Contains(res, "embed") {
			t.Errorf("GetEmbedYTURL() error = %v", err)
			return
		}
	})

	t.Run("Test with shortened Url with query parameters", func(t *testing.T) {
		res, err := GetEmbedYTURL(testUrls[2])
		println(res)
		if err != nil && strings.Contains(res, "embed") && strings.Contains(res, "si=PhsiVtemJm3EEE-v") {
			t.Errorf("GetEmbedYTURL() error = %v", err)
			return
		}
	})
	t.Run("Test with shortened Url with query parameters and time parameter", func(t *testing.T) {
		res, err := GetEmbedYTURL(testUrls[3])
		println(res)
		if err != nil && strings.Contains(res, "embed") && strings.Contains(res, "t=818s") && strings.Contains(res, "si=PhsiVtemJm3EEE-v") {
			t.Errorf("GetEmbedYTURL() error = %v", err)
			return
		}
	})
}
