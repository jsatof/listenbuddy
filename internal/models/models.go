package models

type User struct {
	ID       int64  `json:id`
	Username string `json:username`
	Password string `json:password`
}

type SongRequest struct {
	ID     int64  `json:id`
	UserID int64  `json:userid`
	URL    string `json:url`
}
