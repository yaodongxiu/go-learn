package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

//////////// Date/Time Functions ////////////

func Time() int64 {
	return time.Now().Unix()
}

func Strtotime(format, strtime string) (int64, error) {
	t, err := time.Parse(format, strtime)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}

func Checkdate(month, day, year int) bool {
	if month < 1 || month > 12 || day < 1 || day > 31 || year < 1 {
		return false
	}
	switch month {
	case 4, 6, 9, 11:
		if day > 30 {
			return false
		}
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) && day > 29 {
			return false
		} else if day > 28 {
			return false
		}
	}
	return true
}

func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

func Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}

//////////// String Functions ////////////

func Strpos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

func Stripos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(haystack[offset:]), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

func Strrpos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.LastIndex(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

func Strripos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.LastIndex(strings.ToLower(haystack[offset:]), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

func StrReplace(search, replace, subject string, count int) string {
	return strings.Replace(subject, search, replace, count)
}

func Strtoupper(str string) string {
	return strings.ToUpper(str)
}

func Strtolower(str string) string {
	return strings.ToLower(str)
}

func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}

func Lcfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

func Ucwords(str string) string {
	return strings.ToTitle(str)
}

func Substr(str string, start uint, length int) string {
	if start < 0 || length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

func Strrev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ParseStr(encodedString string, result map[string][]string) error {
	result, err := url.ParseQuery(encodedString)
	return err
}

func NumberFormat(number float64, decimals uint, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimals)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	l1, l2 := len(prefix), len(sep)
	tmp := make([]byte, l2*((l1-1)/3)+l1)
	for i, pos := 0, len(tmp)-1; i < l1; i, pos = i+1, pos-1 {
		if l2 > 0 && i > 0 && i%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-1-j]
				pos--
			}
		}
		tmp[pos] = prefix[l1-1-i]
	}
	s := string(tmp)
	if dec > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

//func ChunkSplit(body string, chunklen uint, end string) string {
//
//}
//
//func StrWordCount(str string) []string {
//
//}
//
//func Wordwrap(str string, width uint, br string, cut bool) string {
//
//}

func Strlen(str string) int {
	return len(str)
}

func MbStrlen(str string) int {
	return utf8.RuneCountInString(str)
}

func StrRepeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}

func Strstr(haystack string, needle string) string {
	index := strings.Index(haystack, needle)
	return haystack[index:]
}

//
//func Strtr(haystack string, params ...interface{}) string {
//
//}
//
//func StrShuffle(str string) string {
//
//}

func Trim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimSpace(str)
	}
	return strings.Trim(str, characterMask[0])
}

func Ltrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimLeftFunc(str, unicode.IsSpace)
	}
	return strings.TrimLeft(str, characterMask[0])
}

func Rtrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}
	return strings.TrimRight(str, characterMask[0])
}

func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

func Chr(ascii int) string {
	return string(rune(ascii))
}

func Ord(char string) int {
	r, _ := utf8.DecodeRune([]byte(char))
	return int(r)
}

//func Nl2br(str string, isXhtml bool) string {
//
//}

func JSONDecode(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}

func JSONEncode(val interface{}) ([]byte, error) {
	return json.Marshal(val)
}

//func Addslashes(str string) string {
//
//}
//
//func Stripslashes(str string) string {
//
//}
//
//func Quotemeta(str string) string {
//
//}
//
//func Htmlentities(str string) string {
//
//}
//
//func HTMLEntityDecode(str string) string {
//
//}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func Md5File(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha1File(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

//func Levenshtein(str1, str2 string, costIns, costRep, costDel int) int {
//
//}
//
//func SimilarText(first, second string, percent *float64) int {
//
//}
//
//func Soundex(str string) string {
//
//}
//
//func ParseURL(str string, component int) (map[string]string, error) {
//
//}

func URLEncode(str string) string {
	return url.QueryEscape(str)
}

func URLDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

//func Rawurlencode(str string) string {
//
//}
//
//func Rawurldecode(str string) (string, error) {
//
//}
//
//func HTTPBuildQuery(queryData url.Values) string {
//
//}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//////////// Array(Slice/Map) Functions ////////////

//func ArrayFill(startIndex int, num uint, value interface{}) map[int]interface{} {
//
//}
//
//func ArrayFlip(m map[interface{}]interface{}) map[interface{}]interface{} {
//
//}
//
//func ArrayKeys(elements map[interface{}]interface{}) []interface{} {
//
//}
//
//func ArrayValues(elements map[interface{}]interface{}) []interface{} {
//
//}
//
//func ArrayMerge(ss ...[]interface{}) []interface{} {
//
//}
//
//func ArrayChunk(s []interface{}, size int) [][]interface{} {
//
//}
//
//func ArrayPad(s []interface{}, size int, val interface{}) []interface{} {
//
//}
//
//func ArraySlice(s []interface{}, offset, length uint) []interface{} {
//
//}
//
//func ArrayRand(elements []interface{}) []interface{} {
//
//}
//
//func ArrayColumn(input map[string]map[string]interface{}, columnKey string) []interface{} {
//
//}
//
//func ArrayPush(s *[]interface{}, elements ...interface{}) int {
//
//}
//
//func ArrayPop(s *[]interface{}) interface{} {
//
//}
//
//func ArrayUnshift(s *[]interface{}, elements ...interface{}) int {
//
//}
//
//func ArrayShift(s *[]interface{}) interface{} {
//
//}
//
//func ArrayKeyExists(key interface{}, m map[interface{}]interface{}) bool {
//
//}
//
//func ArrayCombine(s1, s2 []interface{}) map[interface{}]interface{} {
//
//}
//
//func ArrayReverse(s []interface{}) []interface{} {
//
//}
//
//func Implode(glue string, pieces []string) string {
//
//}
//
//func InArray(needle interface{}, haystack interface{}) bool {
//
//}
//
//func Abs(number float64) float64 {
//
//}
//
//func Rand(min, max int) int {
//
//}
//
//func Round(value float64) float64 {
//
//}
//
//func Floor(value float64) float64 {
//
//}
//
//func Ceil(value float64) float64 {
//
//}
//
//func Pi() float64 {
//
//}
//
//func Max(nums ...float64) float64 {
//
//}
//
//func Min(nums ...float64) float64 {
//
//}
//
//func Decbin(number int64) string {
//
//}
//
//func Bindec(str string) (string, error) {
//
//}
//
//func Hex2bin(data string) (string, error) {
//
//}
//
//func Bin2hex(str string) (string, error) {
//
//}
//
//func Dechex(number int64) string {
//
//}
//
//func Hexdec(str string) (int64, error) {
//
//}
//
//func Decoct(number int64) string {
//
//}
//
//func Octdec(str string) (int64, error) {
//
//}
//
//func BaseConvert(number string, frombase, tobase int) (string, error) {
//
//}
//
//func IsNan(val float64) bool {
//
//}
//
//func Stat(filename string) (os.FileInfo, error) {
//
//}
//
//func Pathinfo(path string, options int) map[string]string {
//
//}
//
//func FileExists(filename string) bool {
//
//}
//
//func IsFile(filename string) bool {
//
//}
//
//func IsDir(filename string) (bool, error) {
//
//}
//
//func FileSize(filename string) (int64, error) {
//
//}
//
//func FilePutContents(filename string, data string, mode os.FileMode) error {
//
//}
//
//func FileGetContents(filename string) (string, error) {
//
//}
//
//func Unlink(filename string) error {
//
//}
//
//func Delete(filename string) error {
//
//}
//
//func Copy(source, dest string) (bool, error) {
//
//}
//
//func IsReadable(filename string) bool {
//
//}
//
//func IsWriteable(filename string) bool {
//
//}
//
//func Rename(oldname, newname string) error {
//
//}
//
//func Touch(filename string) (bool, error) {
//
//}
//
//func Mkdir(filename string, mode os.FileMode) error {
//
//}
//
//func Getcwd() (string, error) {
//
//}
//
//func Realpath(path string) (string, error) {
//
//}
//
//func Basename(path string) string {
//
//}
//
//func Chmod(filename string, mode os.FileMode) bool {
//
//}
//
//func Chown(filename string, uid, gid int) bool {
//
//}
//
//func Fclose(handle *os.File) error {
//
//}
//
//func Filemtime(filename string) (int64, error) {
//
//}
//
//func Fgetcsv(handle *os.File, length int, delimiter rune) ([][]string, error) {
//
//}
//
//func Glob(pattern string) ([]string, error) {
//
//}
//
//func Empty(val interface{}) bool {
//
//}
//
//func IsNumeric(val interface{}) bool {
//
//}
//
//func Exec(command string, output *[]string, returnVar *int) string {
//
//}
//
//func System(command string, returnVar *int) string {
//
//}
//
//func Passthru(command string, returnVar *int) {
//
//}
//
//func Gethostname() (string, error) {
//
//}
//
//func Gethostbyname(hostname string) (string, error) {
//
//}
//
//func Gethostbynamel(hostname string) ([]string, error) {
//
//}
//
//func Gethostbyaddr(ipAddress string) (string, error) {
//
//}
//
//func IP2long(ipAddress string) uint32 {
//
//}
//
//func Long2ip(properAddress uint32) string {
//
//}
//
//func Echo(args ...interface{}) {
//
//}
//
//func Uniqid(prefix string) string {
//
//}
//
//func Exit(status int) {
//
//}
//
//func Die(status int) {
//
//}
//
//func Getenv(varname string) string {
//
//}
//
//func Putenv(setting string) error {
//
//}
//
//func MemoryGetUsage(realUsage bool) uint64 {
//
//}
//
//func VersionCompare(version1, version2, operator string) bool {
//
//}
//
//func ZipOpen(filename string) (*zip.ReadCloser, error) {
//
//}
//
//func Pack(order binary.ByteOrder, data interface{}) (string, error) {
//
//}
//
//func Unpack(order binary.ByteOrder, data string) (interface{}, error) {
//
//}
//
//func Ternary(condition bool, trueVal, falseVal interface{}) interface{} {
//
//}
