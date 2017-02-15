package model

import (
	"strings"
	"errors"
	"regexp"
	"urlShortner/app/util"
)

const domainRemote = "https://afternoon-scrubland-76540.herokuapp.com/"
const domainLocal = "http://127.0.0.1/"
var database = make(map[string]Url) //da togliere

type Url struct {
	ID string `json:"id,omitempty"`
	LongUrl string `json:"longurl,omitempty"`
	ShortUrl string `json:"shortul,omitempty"`
}


func (u *Url) GenerateShortUrl(longUrl string, genesisNumber int) string {
	u.ID = util.RandStringBytesMaskImpr(genesisNumber) //leggere dal json
	if !strings.HasPrefix(longUrl, "http://") && !strings.HasPrefix(longUrl,"https://") {
		longUrl = "http://" + longUrl
	}
	u.LongUrl = longUrl
	u.ShortUrl = domainLocal + u.ID

	database[u.ID] = *u

	return u.ShortUrl
}

func GetLongUrl(id string) (string,error) {
	for _, val := range database {
		if val.ID == id {
			return val.LongUrl, nil
		}
	}

	return "", errors.New("URL not found!")
}

func MatchUrl(url string) error{
	var rp = regexp.MustCompile(`[https?:\/\/]?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	switch {
	case url == "":
		return errors.New("Empty URL")
	case rp.MatchString(url):
		return nil
	default:
		return errors.New("Bad URL")
	}
}