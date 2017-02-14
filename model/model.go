package model

import (
	"regexp"
	"errors"
	"urlShortner/structures"
	"log"
)



func MatchUrl(url string) error{
	var rp = regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	switch {
	case url == "":
		return errors.New("Empty URL")
	case rp.MatchString(url):
		return nil
	default:
		return errors.New("Bad URL")
	}
}

func GenerateShortUrl(longUrl string, number int) string {
	var urlStruct structures.MyUrl
	urlStruct.ID = structures.RandStringBytesMaskImpr(number)
	urlStruct.LongUrl = longUrl
	urlStruct.ShortUrl = structures.Domain + urlStruct.ID

	structures.Database[urlStruct.ID] = urlStruct
	log.Printf("%+v", urlStruct)
	return urlStruct.ShortUrl
}

func GetLongUrl(id string) (string, error) {
	for _, val := range structures.Database {
		if val.ID == id {
			return val.LongUrl, nil
		}
	}
	return "",errors.New("URL not found!")
}