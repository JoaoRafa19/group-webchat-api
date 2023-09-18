package database

type InsertUserModel struct {
	Username       string `json:"username"`
	HashedPassword string `json:"password"`
}
