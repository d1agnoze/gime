package internal

import (
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
  if m.Mime == "" || m.Link == ""{
    return ERR_MISSING_ARG
  }
  if !MimeChecker(Mimetype(m.Mime)){
    return ERR_UNSUPPORT_MIME
  }
	if strings.Contains(m.Mime, "video") {
		link, err := GetEmbedYTURL(m.Link)
		if err != nil {
			return err
		}
		m.Link = link
	}
	return nil
}

func MimeChecker(mimetype Mimetype) bool {
	enumValues := []Mimetype{V_youtube, I_png, I_gif, I_jpeg, I_webp}
	for _, v := range enumValues {
		if mimetype == v {
			return true
		}
	}
	return false
}
