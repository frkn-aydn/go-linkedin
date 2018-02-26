package GoLinkedin

import (
	"io/ioutil"
	"net/http"
)

func (li LinkedIn) APICall(access_token, call string) (string, error) {
	var endpoint string = "/" + li.APIVersion + "/" + li.APICalls[call] + "&format=json&oauth2_access_token=" + access_token

	req, err := http.NewRequest("POST", li.BaseURL+endpoint, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// [TODO:] Parse json before return
	return string(body), nil
}
