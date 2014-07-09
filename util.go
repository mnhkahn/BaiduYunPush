package BaiduYunPush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
 * Performs a GET request to the given url
 * Returns a []byte containing the response body
 */
func HTTPGet(url string) ([]byte, error) {

	r, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("GET failed (%s)", url)
	}

	// read the response and check
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Response read failed.")
	}

	return body, nil
}

/*
 * Performs a HTTP Post request. Takes:
 * * A url
 * * Headers, in the format [][]string{} (e.g., [[key, val], [key, val], ...])
 * * A payload (post request body) which can be nil
 * * Returns the body of the response and an error if necessary
 */
func HTTPPost(url string, headers [][]string, payload *[]byte) ([]byte, error) {
	// setup post client
	client := &http.Client{}
	if payload == nil {
		payload = new([]byte)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(*payload))

	// add headers
	if len(headers) > 0 {
		for i := range headers {
			req.Header.Add(headers[i][0], headers[i][1])
		}
	}

	// perform request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("POST request failed: %s", err))
	}

	// read response, check & return
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return body, fmt.Errorf("%s", resp.Status)
	}

	return body, nil
}

/*
 * Returns a URLEncoded version of a Param Map
 * E.g., ParamMap[foo:bar omg:wtf] => "foo=bar&omg=wtf"
 * TODO: This isn't exactly safe and there's probably a library pkg to do this already...
 */
func EncodeURLParamMap(m *URLParamMap) string {
	r := []string{}

	for k, v := range *m {
		l := len(v)
		for x := 0; x < l; x++ {
			r = append(r, fmt.Sprintf("%s=%s", k, url.QueryEscape(v[x])))
		}
	}

	return strings.Join(r, "&")
}

/*
 * Decodes a json formatted []byte into an interface{} type
 */
func BytesToJSON(b *[]byte) (*interface{}, error) {
	var container interface{}
	err := json.Unmarshal(*b, &container)
	if err != nil {
		return nil, fmt.Errorf("%v %s", err, string(*b))
	}

	return &container, nil
}

/*
 * Encodes a map[string]interface{} to bytes and returns
 * a pointer to said bytes
 */
func JSONToBytes(m map[string]interface{}) (*[]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("Failed to encode JSON")
	}

	return &b, nil
}
