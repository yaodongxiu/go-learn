// 常用函数
package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/url"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func dd(data ...interface{}) {
	fmt.Println(data...)
}

// TestMath 数学函数
func TestMath() {
	dd("-----TestMath-----")

	dd(math.Abs(-4.2))
	dd(math.Ceil(9.999))
	dd(math.Floor(9.999))
	dd(math.Mod(5.7, 1.3))
	dd(math.Pow(-1, 20))
	dd(math.Round(1.95583))
	dd(math.Sqrt(9))
	dd(math.Max(1, 7))
	dd(math.Min(1, 7))
	dd(rand.Intn(200) - 100)
	dd(math.Pi)

	dd()
}

// TestTrimString 去空格或或其他字符
func TestTrimString() {
	dd("-----TestTrimString-----")

	str := "!Hello World!!"
	dd(strings.Trim(str, "!"))
	dd(strings.TrimLeft(str, "!"))
	dd(strings.TrimRight(str, "!"))

	dd(filepath.Dir("c:/testweb/home.php"))
	dd(path.Dir("c:/testweb/home.php"))

	dd(strings.HasPrefix("ftp://192.168.10.1", "ftp"))
	dd(strings.HasSuffix("NLT_abc.jpg", "jpg"))

	dd()
}

// TestConvertString 字符串生成与转化
func TestConvertString() {
	dd("-----TestConvertString-----")

	i := 123
	str := strconv.Itoa(i)
	bytes := []byte(str)
	dd(i, str, bytes)

	raw := "Hello World"
	dd(fmt.Sprintf("%s%s", raw, strings.Repeat(".", 20-len(raw))))

	split := strings.Split(raw, "l")
	dd(split)
	dd(strings.Join(split, "l"))

	dd(reverseString(raw))

	// 将字符串解析成变量 将字符串解析成变量
	params, _ := url.ParseQuery("id=23&name=John%20Adams")
	dd(params)

	dd(strconv.FormatInt(int64(i), 2))
	dd(strings.ToUpper(raw))
	dd(strings.ToLower(raw))

	dd()
}

func main() {
	TestMath()
	TestTrimString()
	TestConvertString()
}

// reverseString 字符串翻转
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
