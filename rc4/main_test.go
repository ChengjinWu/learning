package rc4

import (
	"crypto/rc4"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"
)

var (
	num2char = "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	key      = []byte("fd6cde7c2f4913f22297c948dd530c84") //初始化用于加密的KEY
)

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// 10进制数转换   n 表示进制， 16 or 36
func NumToBHex(num, n int) string {
	num_str := ""
	for num != 0 {
		yu := num % n
		num_str = string(num2char[yu]) + num_str
		num = num / n
	}
	return num_str
}

// 36进制数转换   n 表示进制， 16 or 36
func BHex2Num(str string, n int) int {
	v := 0.0
	length := len(str)
	for i := 0; i < length; i++ {
		s := string(str[i])
		index := strings.Index(num2char, s)
		v += float64(index) * math.Pow(float64(n), float64(length-1-i)) // 倒序
	}
	return int(v)
}

func GenInvitationCode(uid int64) string {
	rc4obj1, _ := rc4.NewCipher(key) //返回 Cipher
	rc4str1 := Int64ToBytes(uid)
	rc4Bytes := rc4str1[3:]
	plaintext := make([]byte, len(rc4Bytes))
	rc4obj1.XORKeyStream(plaintext, rc4Bytes)
	plaintext = append(make([]byte, 3), plaintext...)
	result := BytesToInt64(plaintext)
	return NumToBHex(int(result), 32)
}

func GenRandomInvitationCode(uid int) string {
	code := NumToBHex(uid, 32)
	length := len(code)
	for i := length; i < 6; i++ {
		code = string(num2char[rand.Intn(32)]) + code
	}
	return code
}

func TestRc4(t *testing.T) {
	for i := 1000000; i < 1000100; i++ {
		code := GenRandomInvitationCode(i)
		fmt.Println(code, i)

	}
	fmt.Println(strings.ToUpper("23456789abcdefghjklmnpqrstuvwxyz"))
}
