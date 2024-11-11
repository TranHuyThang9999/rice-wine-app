package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"unicode"

	"github.com/google/uuid"
)

var (
	mu sync.Mutex
)

func GenerateUniqueKey() int64 {
	mu.Lock()
	defer mu.Unlock()

	var length = 7
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	key := int64(0)
	for i := 0; i < length; i++ {
		key = key*10 + int64(seededRand.Intn(9)) + 1
	}

	return key
}

func GeneratePassword() string {
	var length = 5
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	key := int64(0)
	for i := 0; i < length; i++ {
		key = key*10 + int64(seededRand.Intn(9)) + 1
	}

	ranStr := make([]rune, length)

	for i := 0; i < length; i++ {
		ranStr[i] = rune(65 + rand.Intn(25))
	}

	keyinit := fmt.Sprintf("%d%s", key, string(ranStr))
	shuff := []rune(keyinit)
	rand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})

	for i := 0; i < len(shuff); i++ {
		shuff[i] = unicode.ToUpper(shuff[i])
	}

	return string(shuff)
}
func GenerateOtp() int64 {
	var length = 6
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	key := int64(0)
	for i := 0; i < length; i++ {
		key = key*10 + int64(seededRand.Intn(9)) + 1
	}

	return key
}
func GenerateTimestamp() int64 {
	timeNow := time.Now()
	return timeNow.Unix()
}
func ConvertTimestampToDateTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)

	formattedDateTime := t.Format("2006-01-02")

	return formattedDateTime
}
func GenerateTimestampExpiredAt(expiredAt int) *int {
	timeNow := time.Now()

	expirationTime := timeNow.Add(time.Duration(expiredAt) * time.Minute)

	timestamp := int(expirationTime.Unix())
	return &timestamp
}
func GenerateNameFile() string {
	return uuid.NewString()
}
