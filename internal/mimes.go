package internal

import (
	"fmt"
	"strings"
)

type Media struct {
	Mime string
	Link string
}
type Mimetype string

// All supported mime types
const (
	//Custom mimetype: youtube
	V_youtube Mimetype = "video/x-youtube"

	I_png  Mimetype = "image/png"
	I_gif  Mimetype = "image/gif"
	I_jpeg Mimetype = "image/jpeg"
	I_webp Mimetype = "image/webp"
	// for -i flag
	I_wildcard Mimetype = "image/*"
)

func (m *Media) MediaProcess() error {
	if m.Mime == "" || m.Link == "" {
    return fmt.Errorf("%s - MIME:%s - LINK:%s\n", ERR_MISSING_ARG.Error(), m.Mime, m.Link)
	}
	if !MimeChecker(Mimetype(m.Mime)) {
		return fmt.Errorf("%s: %s\n", ERR_UNSUPPORT_MIME.Error(), m.Mime)
	}
	if m.Link = cleanString(m.Link); !isValidURL(m.Link) {
		return ERR_INVALID_URL
	}
	if strings.Contains(m.Mime, "video") {
		var err error
		if m.Link, err = GetEmbedYTURL(m.Link); err != nil {
			return err
		}
	}
	return nil
}

func MimeChecker(mimetype Mimetype) bool {
	enumValues := []Mimetype{V_youtube, I_png, I_gif, I_jpeg, I_webp, I_wildcard}
	for _, v := range enumValues {
		if mimetype == v {
			return true
		}
	}
	return false
}
