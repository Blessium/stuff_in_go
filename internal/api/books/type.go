package books

import (
	"time"
)

type Book struct {
	ISBN      string
	Title     string
	Author    string
	Published time.Time
	Pages     uint
}
