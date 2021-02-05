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
	"unicode/utf8"
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

	dd(filepath.Dir("c:/test/home.php"))
	dd(path.Dir("c:/test/home.php"))

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
	dd(str_pad(raw, 20, "."))
	dd(str_repeat(".", 13))
	dd(strrev(raw))
	dd(parse_str("id=23&name=John%20Adams"))
	dd(strconv.FormatInt(int64(i), 2))
	dd(strings.ToUpper(raw))
	dd(strings.ToLower(raw))

	dd()
}

func main() {
	//TestMath()
	//TestTrimString()
	TestConvertString()
}

func str_pad(str string, length int, pad string) string {
	//return str + strings.Repeat(pad, length-utf8.RuneCountInString(str))
	return fmt.Sprintf("%s%s", str, strings.Repeat(pad, length-utf8.RuneCountInString(str)))
}

func str_repeat(str string, count int) string {
	return strings.Repeat(str, count)
}

func strrev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func parse_str(query string) interface{} {
	params, _ := url.ParseQuery(query)
	return params
}
