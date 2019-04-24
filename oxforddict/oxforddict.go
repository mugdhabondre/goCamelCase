package oxforddict

import (
	"fmt"
	"strings"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
)

var (
	credentialsfile	string = "oxforddict/credentials.json"
	endpoint string = "entries"
	language_code string = "en-us"
	oxfordUrl string = "https://od-api.oxforddictionaries.com/api/v2/"
)

type config struct{
    Appid string
    Appkey string
}

// method to connect to oxford API and check if the given word exists 
func ConnectAndCheck(word_id string) bool{
	config := getcredentials(credentialsfile)

	categories :=  "verb,adjective,adverb,conjunction,numeral,particle,preposition,pronoun,noun"
	url :=  oxfordUrl + endpoint + "/" + language_code + "/" + strings.ToLower(word_id) + "?lexicalCategory=" + categories

	// Create client connection
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		os.Exit(1)
	}
	req.Header.Add("app_id", config.Appid)
	req.Header.Add("app_key", config.Appkey)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// Read the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	responseString := string(body)
	var jsonResp map[string]interface{}
	json.Unmarshal([]byte(responseString), &jsonResp)

	// Check if the response contains error, if it does, the word does not exist
	_, errorexists := jsonResp["error"]

	return !errorexists
}

// Get appi_id and app_key for Oxford Dictionary API
func getcredentials(filename string) config {
	// Open the file, get contents
    content, err := ioutil.ReadFile(filename)
    if err!=nil{
        fmt.Print("Error:",err)
    }

    // Unmarshal contents
    var conf config
    err=json.Unmarshal(content, &conf)
    if err!=nil{
        fmt.Print("Error:",err)
    }
    
    return conf
}