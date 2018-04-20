package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func GetMeasure(w http.ResponseWriter, r *http.Request) {
	cx := appengine.NewContext(r)

	values := url.Values{}
	values.Add("action", "getmeas")
	values.Add("oauth_consumer_key", os.Getenv("NOKIA_OAUTH_CONSUMER_KEY"))
	values.Add("oauth_nonce", os.Getenv("NOKIA_OAUTH_NONCE"))
	values.Add("oauth_signature", os.Getenv("NOKIA_OAUTH_SIGNATURE"))
	values.Add("oauth_signature_method", os.Getenv("NOKIA_OAUTH_SIGNATURE_METHOD"))
	values.Add("oauth_timestamp", os.Getenv("NOKIA_OAUTH_TIMESTAMP"))
	values.Add("oauth_token", os.Getenv("NOKIA_OAUTH_TOKEN"))
	values.Add("oauth_version", os.Getenv("NOKIA_OAUTH_VERSION"))
	values.Add("userid", os.Getenv("NOKIA_USER_ID"))

	client := urlfetch.Client(cx)
	query := values.Encode()
	resp, err := client.Get("http://api.health.nokia.com/measure?" + query)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var d interface{}

	decoder.Decode(&d)

	ret, err := json.Marshal(d)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(ret)
}
