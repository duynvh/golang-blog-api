package common

import (
	"errors"
	log "golang-blog-api/log"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Error("Recovery error:", err)
	}
}
