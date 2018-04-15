package main

import (
	"github.com/ChimeraCoder/anaconda"
	"os"
	"fmt"
	"encoding/json"
)

type twitterCredentials struct {
	Key string `json:"key"`
	Secret string `json:"secret"`
	AccessToken string `json:token`
	AccessSecret string `json:access-secret`
}

func main() {
	//read config file
	var cred twitterCredentials;
	var credFile, err = os.Open("config.json")
	defer credFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	var jsonParser = json.NewDecoder(credFile)
	jsonParser.Decode(&cred)
	//set anaconda stuff
	anaconda.SetConsumerKey(cred.Key)
	anaconda.SetConsumerSecret(cred.Secret)
	api := anaconda.NewTwitterApi(cred.AccessToken, cred.AccessSecret)
	//sample use from anaconda docs
	//MVP connect to twitter api
	searchResult, _ := api.GetSearch("golang", nil)
	for _ , tweet := range searchResult.Statuses {
		fmt.Print(tweet.Text)
	}

}


// connect to a random twitter user

//get their tweet

//display tweet (on console)

//MVP 2, take the tweet, turn it into a rap

//MVP 3 make a web client?
