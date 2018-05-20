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

var rateTicker = time.NewTicker(time.Millisecond * 10)

func Fetch(url string) ([]byte, error) {
	// rate limit to avoid being brutal
	<-rateTicker.C
	// log.Printf("Fetching %s...\n", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", " Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("error fetching url " + url + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v, %v", resp.StatusCode, resp)
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
