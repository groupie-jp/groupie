package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
}

func main() {
	loadEnv()

	v := url.Values{}
	v.Set("count", "30")

	targetAccountList := configureTargetAccount()
	printTweet(targetAccountList)
}

func configureTargetAccount() []string {
	yoshimotoMugendaiHall := "y_mugendai_hall"
	lumineTheYoshimoto := "lumineseisaku"
	targetIdList := []string{
		yoshimotoMugendaiHall,
		lumineTheYoshimoto,
	}
	return targetIdList
}

func printTweet(userIdList []string) {
	api := getTwitterApi()
	v := url.Values{}
	v.Set("count", "200")

	for _, userId := range userIdList {
		v.Set("screen_name", userId)
		targetUsersTweet, err := api.GetUserTimeline(v)
		if err != nil {
			fmt.Println("****************occur error*************")
			fmt.Println(err)
		}

		for _, tweet := range targetUsersTweet {
			if strings.Contains(tweet.FullText, os.Getenv("TARGET")) {
				str := "https://twitter.com/" + tweet.User.ScreenName + "/status/" + strconv.FormatInt(tweet.Id, 10)
				fmt.Println(tweet.FullText)
				fmt.Println(str)
			}
		}
	}
}
