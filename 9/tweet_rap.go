package main

import (
	"github.com/ChimeraCoder/anaconda"
	"os"
	"fmt"
	"encoding/json"
	"bufio"
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

	//read input here
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	//get a user// connect to a random twitter user
	user, _ := api.GetUsersShow(text, nil)
	//fmt.Println(user)
	var userID = user.IdStr
	var bar = make(map[string][]string)
	//get their tweet
	bar["user_id"] = []string{userID}
	tweets, err := api.GetUserTimeline(bar)
	//display tweet (on console)
	for _, tweet := range tweets {
		fmt.Println(tweet.CreatedAt, tweet.Text)
	}
}






//MVP 2, take the tweet, turn it into a rap

//MVP 3 make a web client?
