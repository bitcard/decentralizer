package utils

import (
	"time"
	"fmt"
	"github.com/iain17/logger"
	"errors"
)

//Returns true if the second passed record is newer than the first one.
func IsNewerRecord(current uint64, new uint64) bool {
	if new == 0 {
		return false
	}
	if current == 0 && new != 0 {
		return true
	}
	now := time.Now().UTC()
	publishedTime := time.Unix(int64(new), 0).UTC()
	publishedTimeText := publishedTime.String()
	expireTime := time.Unix(int64(current), 0).UTC()
	expireTimeText := expireTime.String()
	if !publishedTime.After(expireTime) {
		err := fmt.Errorf("record with publish date %s is not newer than %s", publishedTimeText, expireTimeText)
		logger.Warning(err)
		return false
	}
	if publishedTime.After(now) {
		err := errors.New("new peer with publish date %s was published in the future")
		logger.Warning(err)
		return false
	}
	logger.Infof("record with publish date %s IS newer than %s", publishedTimeText, expireTimeText)
	return true
}