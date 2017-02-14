package model

import (
	"regexp"
	"errors"
	"urlShortner/structures"
)

var rp = regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)


func MatchUrl(url string) error{
	switch {
	case rp.MatchString(url):
		return nil
	default:
		return errors.New("Bad URL")
	}
}

func GenerateShortUrl(longUrl string, number int) structures.MyUrl {
	var urlStruct structures.MyUrl
	urlStruct.ID = structures.RandStringBytesMaskImpr(number)
	urlStruct.LongUrl = longUrl
	urlStruct.ShortUrl = structures.Domain + urlStruct.ID

	return urlStruct
}