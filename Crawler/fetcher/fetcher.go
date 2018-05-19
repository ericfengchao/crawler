package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rateTicker = time.NewTicker(time.Millisecond * 50)

func Fetch(url string) ([]byte, error) {
	// rate limit to avoid being brutal
	<-rateTicker.C
	// log.Printf("Fetching %s...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("error fetching url " + url + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s, %v", resp.StatusCode, resp)
	}
	r := bufio.NewReader(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, determineEncoding(r).NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
