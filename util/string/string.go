package string

import (
	"strconv"
	"time"
	"bytes"
	"strings"
	"encoding/json"
	"unsafe"
	"fmt"
)

func String(value interface{}) (string,error) {
	switch v := value.(type) {
	case int:
		return strconv.FormatInt(int64(v), 10),nil
	case int8:
		return strconv.FormatInt(int64(v), 10),nil
	case int16:
		return strconv.FormatInt(int64(v), 10),nil
	case int32:
		return strconv.FormatInt(int64(v), 10),nil
	case int64:
		return strconv.FormatInt(v, 10),nil
	case uint:
		return strconv.FormatUint(uint64(v), 10),nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10),nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10),nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10),nil
	case uint64:
		return strconv.FormatUint(uint64(v), 10),nil
	case []byte:
		return bytes2str(v),nil
	case bool:
		if v == true {
			return "1",nil
		}
		return "0",nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64),nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64),nil
	case string:
		return v,nil
	default:
		return "", fmt.Errorf("data not type : %v", v)
	}
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// bytes to string
func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}


//字符串转数值---------------->
func ToInt(s string)(int,error)  {
	return strconv.Atoi(s)
}

func ToInt32(s string)(int32,error)  {
	i,err := strconv.ParseInt(s, 10, 32)
	return int32(i),err
}

func ToInt64(s string)(int64,error)  {
	return strconv.ParseInt(s, 10, 64)
}

func ToUint64(s string)(uint64,error)  {
	return strconv.ParseUint(s, 10, 64)
}

func ToFloat32(s string)(float32,error)  {
	f,err:= strconv.ParseFloat(s, 32)
	return float32(f),err
}

func ToFloat64(s string)(float64,error)  {
	return strconv.ParseFloat(s, 64)
}

func ToBool(s string)(bool,error)  {
	return strconv.ParseBool(s)
}
//-------------------------------------->



//数值转字符串----------------------->
func FormatInt(i int)string  {
	return strconv.FormatInt(int64(i),10)
}

func FormatInt32(i int32)string  {
	return strconv.FormatInt(int64(i),10)
}

func formatInt64(i int64)string  {
	return strconv.FormatInt(i,10)
}

func FormatUint64(i uint64)string  {
	return strconv.FormatUint(i,10)
}

//而 *prec* 表示有效数字（对 *fmt='b'* 无效），对于 'e', 'E' 和 'f'，有效数字用于小数点之后的位数；对于 'g' 和 'G'，则是所有的有效数字。
func FormatFloat(f float64, fmt byte, prec, bitSize int)string  {
	return strconv.FormatFloat(f,fmt,prec,bitSize)
}

func FormatBool(b bool)string  {
	return strconv.FormatBool(b)
}
//-------------------------------------->



//字符串转时间Duration----------------------->
func ToTimeDuration(s string)(time.Duration,error)  {
	return time.ParseDuration(s)
}
//-------------------------------------->


//将字符串数组（或slice）连接起来可以通过 Join 实现
func JoinSep(str []string, sep string) string {
	// 特殊情况应该做处理
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str[0]
	}
	buffer := bytes.NewBufferString(str[0])
	for _, s := range str[1:] {
		buffer.WriteString(sep)
		buffer.WriteString(s)
	}
	return buffer.String()
}

func Join(str []string) string {
	// 特殊情况应该做处理
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str[0]
	}
	buffer := bytes.NewBufferString(str[0])
	for _, s := range str[1:] {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func Split(s, sep string) []string{
	return strings.Split(s,sep)
}

func Replace(s, old, new string) string{
	return strings.Replace(s,old,new,-1)
}

func Repeat(s string, count int) string{
	return strings.Repeat(s,count)
}

func ToMap(s string)(map[string]interface{},error)  {
	var m map[string]interface{}
	err:=json.Unmarshal([]byte(s),&m)
	return m,err
}

func ToByte(s string)[]byte  {
	return []byte(s)
}

func FormatByte(b []byte)string  {
	return *(*string)(unsafe.Pointer(&b))
}


