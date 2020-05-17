package model

func CreateRelation(loginUserID, targetFolloweeID uint) {

	db := InitDB()
	relation := Relation{
		UserID:   loginUserID,
		FollowID: targetFolloweeID,
	}

	db.Create(&relation)

	defer db.Close()

}

func GetFolloweeID(targetUserID uint) []int {

	db := InitDB()
	followee := []Relation{}

	// 該当ユーザーのfolloweeを取り出す
	db.Where("user_id = ?", targetUserID).Find(&followee)

	defer db.Close()

	// 構造体からFolloweeIDを取り出す
	followeeIDs := make([]uint, len(followee))
	for i, v := range followee {
		followeeIDs[i] = v.FollowID
	}

	// 取り出したFolloweeIDをintに変換
	castedIDs := make([]int, len(followeeIDs))
	for n := range followeeIDs {
		castedIDs[n] = int(followeeIDs[n])
	}

	return castedIDs

}

func GetFollowerID(loginUserID uint) []int {

	db := InitDB()
	follower := []Relation{}

	// 該当ユーザーのfolloweeを取り出す
	db.Where("follow_id = ?", loginUserID).Find(&follower)

	defer db.Close()

	// 構造体からFolloweeIDを取り出す
	followerIDs := make([]uint, len(follower))
	for i, v := range follower {
		followerIDs[i] = v.UserID
	}

	// 取り出したFolloweeIDをintに変換
	castedIDs := make([]int, len(followerIDs))
	for n := range followerIDs {
		castedIDs[n] = int(followerIDs[n])
	}

	return castedIDs

}

func DeleteRelation(loginUserID, targetFolloweeID uint) {

	db := InitDB()
	relation := Relation{
		UserID:   loginUserID,
		FollowID: targetFolloweeID,
	}

	db.Unscoped().Delete(relation, "user_id LIKE ? AND follow_id LIKE ?", relation.UserID, relation.FollowID)

	defer db.Close()

}
