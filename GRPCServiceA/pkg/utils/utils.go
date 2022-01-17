package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"

	"github.com/google/uuid"
)

func TimeNow() string {
	return strings.Replace(time.Now().Format("01-02-2006 15:04"), " ", "(", 1) + ")"
}

func HashPwd(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func NewId() string {
	return uuid.New().String()
}
