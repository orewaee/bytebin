package dto

import "time"

type Meta struct {
	Id          string        `json:"id"`
	ContentType string        `json:"content_type"`
	Ip          string        `json:"ip"`
	UserAgent   string        `json:"user_agent"`
	CreatedAt   time.Time     `json:"created_at"`
	Lifetime    time.Duration `json:"lifetime"`
}
