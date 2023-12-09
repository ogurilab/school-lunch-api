package util

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomUlid() string {
	return NewUlid()
}

func RandomDate() time.Time {
	// YYYY-MM-DD

	year := RandomInt(1111, 2020)
	month := RandomInt(1, 12)
	day := RandomInt(1, 31)

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomURL() string {
	return fmt.Sprintf("https://%s.com", RandomString(10))
}

func RandomNullURL() sql.NullString {
	return sql.NullString{
		String: RandomURL(),
		Valid:  true,
	}
}

func RandomInt32() int32 {
	return int32(RandomInt(0, 100))
}

func RandomCityCode() int32 {
	return int32(RandomInt(0, 4713))
}
