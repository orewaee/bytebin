package domain

import "time"

type Meta struct {
	Id          string
	ContentType string
	Ip          string
	UserAgent   string
	CreatedAt   time.Time
	Lifetime    time.Duration
}
