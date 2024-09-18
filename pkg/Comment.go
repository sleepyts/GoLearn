package pkg

import "time"

type Object interface{
	String() string
}
type Comment struct{
	Id          int `json:"id"`
	Content     string `json:"content"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time ` json:"updated_time"`
	ReplyOnId 	int `json:"reply_blog_id"`
}

