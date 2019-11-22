/*
 * @lc app=leetcode.cn id=355 lang=golang
 *
 * [355] 设计推特
 *
 * https://leetcode-cn.com/problems/design-twitter/description/
 *
 * algorithms
 * Medium (33.96%)
 * Likes:    39
 * Dislikes: 0
 * Total Accepted:    1.7K
 * Total Submissions: 4.9K
 * Testcase Example:  '["Twitter","postTweet","getNewsFeed","follow","postTweet","getNewsFeed","unfollow","getNewsFeed"]\n[[],[1,5],[1],[1,2],[2,6],[1],[1,2],[1]]'
 *
 *
 * 设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：
 *
 *
 * postTweet(userId, tweetId): 创建一条新的推文
 * getNewsFeed(userId):
 * 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
 * follow(followerId, followeeId): 关注一个用户
 * unfollow(followerId, followeeId): 取消关注一个用户
 *
 *
 * 示例:
 *
 *
 * Twitter twitter = new Twitter();
 *
 * // 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
 * twitter.postTweet(1, 5);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
 * twitter.getNewsFeed(1);
 *
 * // 用户1关注了用户2.
 * twitter.follow(1, 2);
 *
 * // 用户2发送了一个新推文 (推文id = 6).
 * twitter.postTweet(2, 6);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
 * // 推文id6应当在推文id5之前，因为它是在5之后发送的.
 * twitter.getNewsFeed(1);
 *
 * // 用户1取消关注了用户2.
 * twitter.unfollow(1, 2);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
 * // 因为用户1已经不再关注用户2.
 * twitter.getNewsFeed(1);
 *
 *
 */

// @lc code=start
import (
	"container/heap"
)

type Tweet struct {
	Id    int
	Index int
}

type TweetHeap []Tweet

func (h *TweetHeap) Less(i, j int) bool {
	return (*h)[i].Index < (*h)[j].Index
}

func (h *TweetHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *TweetHeap) Len() int {
	return len(*h)
}

func (h *TweetHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *TweetHeap) Push(v interface{}) {
	*h = append(*h, v.(Tweet))
}

type Twitter struct {
	FollowMap map[int]map[int]int // uid: map{uid: 1}
	TweetMap  map[int][]Tweet     // uid: tweet
	MaxIndex  int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		FollowMap: map[int]map[int]int{},
		TweetMap:  map[int][]Tweet{},
		MaxIndex:  0,
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.MaxIndex++
	index := this.MaxIndex
	tweet := Tweet{Id: tweetId, Index: index}
	if _, ok := this.TweetMap[userId]; !ok {
		this.TweetMap[userId] = []Tweet{}
	}
	this.TweetMap[userId] = append(this.TweetMap[userId], tweet)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	th := make(TweetHeap, 0)
	heap.Init(&th)
	limit := 10 // 限制10条
	uids := []int{userId}
	for uid := range this.FollowMap[userId] {
		if uid != userId {
			uids = append(uids, uid)
		}
	}
	for _, uid := range uids {
		tweets := this.TweetMap[uid]
		for _, el := range tweets {
			if th.Len() >= limit {
				if th[0].Index >= el.Index {
					continue // 发表时间比堆顶小，跳过
				}
				heap.Pop(&th)
			}
			// 判断堆头
			heap.Push(&th, el)
		}
	}
	rev := []int{}
	for th.Len() > 0 {
		rev = append(rev, heap.Pop(&th).(Tweet).Id)
	}
	ret := make([]int, 0)
	for i := len(rev) - 1; i >= 0; i-- {
		ret = append(ret, rev[i])
	}
	return ret
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.FollowMap[followerId]; !ok {
		this.FollowMap[followerId] = map[int]int{}
	}
	this.FollowMap[followerId][followeeId] = 1
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.FollowMap[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end

