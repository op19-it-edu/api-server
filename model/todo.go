package model

import (
	"time"
)

func CreateTodo(userID uint, todo string, deadline time.Time) {

	db := InitDB()
	todoTweet := TodoTweet{
		UserID:   userID,
		Todo:     todo,
		Deadline: deadline,
		Done:     false,
	}

	db.Create(&todoTweet)

	defer db.Close()
}

func GetTodo(todoID int) TodoTweet {

	db := InitDB()
	todoTweet := TodoTweet{}
	db.First(&todoTweet, todoID)

	defer db.Close()

	return todoTweet
}

func UpdateTodo(todoID int, todo string, deadline time.Time, done bool) {

	db := InitDB()
	todoTweet := TodoTweet{}

	// 空白のフィールドは無視してレコードを更新する
	db.Model(&todoTweet).Where("id = ?", todoID).Updates(TodoTweet{Todo: todo, Deadline: deadline, Done: done})

	defer db.Close()
}

func DeleteTodo(todoID int) {

	db := InitDB()
	todoTweet := TodoTweet{}

	// Deleteする際のid指定はuintじゃないといけないので、intからuintにキャストしておく
	castedID := uint(todoID)

	// キャストしたuserIDを構造体に代入する
	todoTweet.ID = castedID

	db.Unscoped().Delete(&todoTweet)

	defer db.Close()
}

func GetFolloweeTodos(followeeIDs []uint) []TodoTweet {

	db := InitDB()
	todoTweets := [][]TodoTweet{}

	// FolloweeIDの数だけレコードを取得して追加していく
	for n := range followeeIDs {
		todoTweet := []TodoTweet{}
		db.Where(&TodoTweet{UserID: followeeIDs[n]}).Find(&todoTweet)
		todoTweets = append(todoTweets, todoTweet)
	}

	defer db.Close()

	// ネストしている構造体をフラットにする
	todoTweet := []TodoTweet{}
	for i, _ := range todoTweets {
		result := todoTweets[i]
		for i, _ := range result {
			todoTweet = append(todoTweet, result[i])
		}
	}

	return todoTweet
}
