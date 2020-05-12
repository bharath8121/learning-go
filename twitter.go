package main

import (
	"fmt"
	"sort"
)

type Twitter struct{

	userTweets map[int][][2]int
	userFollows map[int][]int
	timeStamp int

}


func Constructor() Twitter{

	var twitter Twitter

	twitter.userTweets = make(map[int][][2]int)
	twitter.userFollows = make(map[int][]int)
	twitter.timeStamp= 0

	return twitter

}


func (this *Twitter) PostTweet(userId, tweetId int){

	fmt.Println("Befor Tweet: ", this.userTweets[userId])
	this.timeStamp++
	this.userTweets[userId] = append(this.userTweets[userId], [2]int{tweetId, this.timeStamp})
	fmt.Println("After Tweet: ", this.userTweets[userId])

}


func (this *Twitter) GetNewsFeed(userId int) []int {

	var feeds []int
    users := append(this.userFollows[userId], userId)
	var allFeeds [][2]int
	for _, user := range(users){
		allFeeds = append(allFeeds, this.userTweets[user]...)
	}

	sort.SliceStable(allFeeds, func(i, j int) bool{
		return allFeeds[i][1] > allFeeds[j][1]
	})

	for _, feed := range(allFeeds){
		if len(feeds) == 10{
			break
		}
		if isNotIn(feeds, feed[0]){
			feeds = append(feeds, feed[0])
		}
		
	}

	fmt.Println(feeds)
	return feeds
    
}


func isNotIn(feeds []int, feed int) bool{

	for _, f := range(feeds){
		if feed == f{
			return false
		}
	}
	return true

}

func (this *Twitter) Follow(followerId int, followeeId int)  {

	fmt.Println("Before Follows: ", this.userFollows[followerId])
	for _, followee := range(this.userFollows[followerId]){
		if followee == followeeId{
			return
		}
	}
	this.userFollows[followerId] = append(this.userFollows[followerId], followeeId)
	fmt.Println("After Follows: ", this.userFollows[followerId])

}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {

	fmt.Println("Before Unfollow: ", this.userFollows[followerId])
    for indx, element :=  range(this.userFollows[followerId]){
		if element == followeeId{
			this.userFollows[followerId] = append(this.userFollows[followerId][:indx], this.userFollows[followerId][indx+1:]...)
			break
		}
	}
	fmt.Println("After Unfollow: ", this.userFollows[followerId])

}
