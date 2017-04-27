package DOIpUpdate

import (
	"net/http"
	"io/ioutil"
)

func GetIP() (string) {
	resp, err := http.Get("http://whatismyip.akamai.com/")
	if (err != nil) {
		panic("Can not get IP from whatismyip.akamai.com")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if (err != nil) {
		panic("Can not get IP from whatismyip.akamai.com")
	}

	return string(body)
}