package utils

import (
	"math/big"
	"strconv"
	"time"
)

const (
	formatTime     = "15:04:05"
	formatDate     = "2006-01-02"
	formatDateTime = "2006-01-02 15:04:05"
)

type ArrayList struct {
	elements []interface{}
	size     int
}

func New(values ...interface{}) *ArrayList {
	list := &ArrayList{}
	list.elements = make([]interface{}, 10)
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *ArrayList) Add(values ...interface{}) {
	if list.size+len(values) >= len(list.elements)-1 {
		newElements := make([]interface{}, list.size+len(values)+1)
		copy(newElements, list.elements)
		list.elements = newElements
	}

	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

type String string

// Set string
func (f *String) Set(v string) {
	if v != "" {
		*f = String(v)
	} else {
		f.Clear()
	}
}

// Clear string
func (f *String) Clear() {
	*f = String(0x1E)
}

// Exist check string exist
func (f String) Exist() bool {
	return string(f) != string(0x1E)
}

// Bool string to bool
func (f String) Bool() (bool, error) {
	return strconv.ParseBool(f.String())
}

// Float32 string to float32
func (f String) Float32() (float32, error) {
	v, err := strconv.ParseFloat(f.String(), 32)
	return float32(v), err
}

// Float64 string to float64
func (f String) Float64() (float64, error) {
	return strconv.ParseFloat(f.String(), 64)
}

// Int string to int
func (f String) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int(v), err
}

// Int8 string to int8
func (f String) Int8() (int8, error) {
	v, err := strconv.ParseInt(f.String(), 10, 8)
	return int8(v), err
}

// Int16 string to int16
func (f String) Int16() (int16, error) {
	v, err := strconv.ParseInt(f.String(), 10, 16)
	return int16(v), err
}

// Int32 string to int32
func (f String) Int32() (int32, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int32(v), err
}

// Int64 string to int64
func (f String) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	if err != nil {
		i := new(big.Int)
		ni, ok := i.SetString(f.String(), 10) // octal
		if !ok {
			return v, err
		}
		return ni.Int64(), nil
	}
	return v, err
}

// Uint string to uint
func (f String) Uint() (uint, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint(v), err
}

// Uint8 string to uint8
func (f String) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

// Uint16 string to uint16
func (f String) Uint16() (uint16, error) {
	v, err := strconv.ParseUint(f.String(), 10, 16)
	return uint16(v), err
}

// Uint32 string to uint32
func (f String) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint32(v), err
}

// Uint64 string to uint64
func (f String) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(f.String(), 10, 64)
	if err != nil {
		i := new(big.Int)
		ni, ok := i.SetString(f.String(), 10)
		if !ok {
			return v, err
		}
		return ni.Uint64(), nil
	}
	return v, err
}

func (f String) DateTime(format ...string) (time.Time, error) {
	return f.Date(formatDateTime)
}

func (f String) Time(format ...string) (time.Time, error) {
	return f.Date(formatTime)
}

func (f String) Date(format ...string) (time.Time, error) {
	var ft string
	if len(format) == 0 {
		ft = formatDate
	} else {
		ft = format[0]
	}
	t, err := time.Parse(ft, f.String())
	return t, err
}

// String string to string
func (f String) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}

func (list *ArrayList) Remove(index int) interface{} {
	if index < 0 || index >= list.size {
		return nil
	}

	curEle := list.elements[index]
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	return curEle
}

func (list *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= list.size {
		return nil
	}
	return list.elements[index]
}

func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

func (list *ArrayList) Size() int {
	return list.size
}
func (list *ArrayList) Contains(value interface{}) (bool, int) {
	for index, curValue := range list.elements {
		if curValue == value {
			return true, index
		}
	}
	return false, -1
}