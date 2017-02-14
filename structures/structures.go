package structures

import (
	"math/rand"
	"time"
	"sync"

)
var database = make(map[string]MyUrl)

type MyUrl struct {
	ID string `json:"id,omitempty"`
	LongUrl string `json:"longurl,omitempty"`
	ShortUrl string `json:"shortul,omitempty"`
}

const Domain = "http://127.0.0.1/"
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImpr(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

var mutex sync.Mutex

func int63() int64 {
	mutex.Lock()
	v := rand.Int63()
	mutex.Unlock()
	return v
}