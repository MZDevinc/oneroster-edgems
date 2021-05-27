package oauth1

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/MZDevinc/edgemsv2/src/utils"
)

// Object for the OneRoster containing client_id and client_secret
type OneRoster struct {
	clientId     string
	clientSecret string
}

// Constructor for OneRoster
func OneRosterNew(clientId string, clientSecret string) OneRoster {
	rr := OneRoster{clientId, clientSecret}
	return rr
}

// Makes the request to the given OneRoster URL with the stored key and secret
// It returns a the status code and JSON response
func (rr OneRoster) MakeRosterRequest(reqUrl string) (int, string, http.Header) {
	timestamp := fmt.Sprint(time.Now().Unix())
	nonce := generateNonce(len(timestamp))

	oauth := make(map[string]string)
	oauth["oauth_consumer_key"] = rr.clientId
	oauth["oauth_signature_method"] = "HMAC-SHA256"
	oauth["oauth_timestamp"] = timestamp
	oauth["oauth_nonce"] = nonce
	oauth["oauth_callback"] = "about:blank"
	oauth["oauth_version"] = "1.0"

	urlPieces := strings.Split(reqUrl, "?")
	allParams := make(map[string]string)
	if len(urlPieces) == 2 {
		urlParams := paramsToMap(urlPieces[1])
		allParams = mergeParams(urlParams, oauth)
	} else {
		allParams = oauth
	}
	baseInfo := buildBaseString(urlPieces[0], "GET", allParams)
	compositeKey := url.QueryEscape(rr.clientSecret) + "&"

	oauth["oauth_signature"] = generateSig(baseInfo, compositeKey)

	authHeader := buildAuthHeader(oauth)
	fmt.Println("AUTH")
	fmt.Println("clientId", rr.clientId)
	fmt.Println("clientSecret", rr.clientSecret)
	fmt.Println("compositeKey", compositeKey)
	fmt.Println("AUTH HEADER")
	utils.PrintJSON(authHeader)
	return makeGetRequest(reqUrl, authHeader)
}

// Makes the actual request to the URL with the generated auth header
func makeGetRequest(reqUrl, authHeader string) (int, string, http.Header) {
	hc := http.Client{}
	req, _ := http.NewRequest("GET", reqUrl, nil)

	req.Header.Add("Authorization", authHeader)

	fmt.Println("REQUEST URL")
	fmt.Println(reqUrl)
	fmt.Println("AUTH HEADER")
	fmt.Println(authHeader)

	resp, err := hc.Do(req)
	if err != nil {
		return 0, "An error occurred, check your URL", nil
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, string(bodyBytes), resp.Header
}

// Creates the auth header from a map of the oauth parameters
func buildAuthHeader(oauthInfo map[string]string) string {
	result := "OAuth "
	values := make([]string, 0)
	for key, value := range oauthInfo {
		values = append(values, key+"=\""+url.QueryEscape(value)+"\"")
	}
	result += strings.Join(values, ",")
	return result
}

// Generates the auth signature from the base info and composite key
func generateSig(baseInfo, compositeKey string) string {
	h := hmac.New(sha256.New, []byte(compositeKey))
	h.Write([]byte(baseInfo))
	sha := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return sha
}

// Generates the base string from the base URL, method, and all of the params
func buildBaseString(baseUrl, method string, params map[string]string) string {
	r := make([]string, 0)
	keys := make([]string, 0)
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		r = append(r, key+"="+urlEscape(params[key]))
	}
	return method + "&" + url.QueryEscape(baseUrl) + "&" + url.QueryEscape(strings.Join(r, "&"))
}

// URL encodes the string
func urlEscape(str string) string {
	escaped := url.QueryEscape(str)
	escaped = strings.Replace(escaped, "+", "%20", -1)
	return escaped
}

// Merges the two maps of params into one map of params
func mergeParams(urlParams, oauth map[string]string) map[string]string {
	result := make(map[string]string)
	for key, value := range urlParams {
		result[key] = value
	}

	for key, value := range oauth {
		result[key] = value
	}
	return result
}

// Converts the params in the url string to a map
func paramsToMap(urlParams string) map[string]string {
	result := make(map[string]string)
	params := strings.Split(urlParams, "&")

	for _, value := range params {
		value, _ := url.PathUnescape(value)
		split := strings.Split(value, "=")
		if len(split) == 2 {
			result[split[0]] = split[1]
		} else {
			result["filter"] = value[7:]
		}
	}
	return result
}

// Generates a random string for the nonce of a given length
func generateNonce(length int) string {
	rand.Seed(time.Now().UnixNano())
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	characterLen := len(characters)
	randomString := ""
	for i := 0; i < length; i++ {
		randomString += string(characters[rand.Intn(characterLen-1)])
	}
	return randomString
}
