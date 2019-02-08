package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	"golang.org/x/text/width"
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
	targetAccountList := configureTargetAccount()
	keywordPatterns := generateKeywordPattern()
	filterTweets := filterTweet(targetAccountList, keywordPatterns)
	printTweet(filterTweets)
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

func filterTweet(userIdList []string, keywordPatterns []string) []anaconda.Tweet {
	api := getTwitterApi()
	v := url.Values{}
	v.Set("count", "200")

	targetUsersTweets := []anaconda.Tweet{}
	for _, userId := range userIdList {
		v.Set("screen_name", userId)
		tweets, err := api.GetUserTimeline(v)
		if err != nil {
			fmt.Println("****************occur error*************")
			fmt.Println(err)
		}

		for _, tweet := range tweets {
			for _, keywordPattern := range keywordPatterns {
				if strings.Contains(tweet.FullText, keywordPattern) {
					targetUsersTweets = append(targetUsersTweets, tweet)
				}
			}
		}
	}
	return targetUsersTweets
}

func printTweet(targetUsersTweet []anaconda.Tweet) {
	for _, tweet := range targetUsersTweet {
		fmt.Println(tweet.FullText)
		url := "https://twitter.com/" + tweet.User.ScreenName + "/status/" + strconv.FormatInt(tweet.Id, 10)
		fmt.Println(url)
	}
}

func generateKeywordPattern() []string {
	var keywordPattern []string
	keywordPattern = append(keywordPattern, os.Getenv("TARGET"))
	keywordPattern = append(keywordPattern, width.Narrow.String(os.Getenv("TARGET")))
	return keywordPattern
}
