package designDataStructure

// 355
type Twitter struct {
	tweets     [][2]int
	followList map[int]map[int]struct{}
}

/** Initialize your data structure here. */
func ConstructorTwitter() Twitter {
	t := Twitter{
		tweets:     make([][2]int, 0),
		followList: make(map[int]map[int]struct{}),
	}
	return t
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets = append(this.tweets, [2]int{userId, tweetId})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0, 10)
	for i := len(this.tweets) - 1; i >= 0; i-- {
		// check every tweets in the list
		// if the author user id in the follow list of current user, or the author user is current user
		// add to result
		if _, ok := this.followList[userId][this.tweets[i][0]]; ok || userId == this.tweets[i][0] {
			res = append(res, this.tweets[i][1])
		}
		if len(res) == 10 {
			break
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.followList[followerId]; !ok {
		this.followList[followerId] = make(map[int]struct{})
	}
	this.followList[followerId][followeeId] = struct{}{}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.followList[followerId]; !ok {
		return
	}
	delete(this.followList[followerId], followeeId)
}
