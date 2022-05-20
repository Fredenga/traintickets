package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


//TODO: seed value(int64), ensures that everytime we run the code, the generated values are different
func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int32{
	return int32(min + rand.Int63n((max-min+1)))
}
func RandomNum() int32{
	return RandomInt(0, 2^64)
}
func RandomMoney() int32{
	return RandomInt(0, 1000)
	
}
func RandomWords() string{
	return RandomString(10)
}
func RandomType(word string) string{
	types := []string{"express", "local"}
	class := []string{"luxury", "economy", "business"}
	switch word{
	case "class":
		n := len(class)
		return class[rand.Intn(n-1)]
	case "types":
		n := len(types)
		return types[rand.Intn(n-1)]
	}
	return ""
}
//TODO: generate a random string of n characters
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)

		//!index of alphabet is random number between 0->length of string
		//!c becomes the indexed element, is added to the string and process repeats n times
		
	}
	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}

func RandomTime() time.Time {
	return time.Now().AddDate(
		rand.Intn(1),
		rand.Intn(12),
		rand.Intn(30),		
	)
}

func RandomRoute() string {
	return fmt.Sprintf("%s to %s", RandomString(6), RandomString(6))
}