package entity

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

const (
	timeFormart = "2006-01-02 15:04:05"
)

var jsonNull = []byte("null")

// XTime 兼容JSON与Mysql数据类型
// 其中主要的类参考 https://github.com/gocraft/dbr/blob/master/types.go 生成兼容JSON与Mysql NullObject的数据类型
type XTime struct {
	time.Time
	Valid bool // Valid is true if Time is not NULL
}

// NewXTime creates a NewXTime with Scan().
func NewXTime(v interface{}) (n XTime) {
	n.Scan(v)
	return
}

// The `(*XTime) Scan(interface{})` and `parseDateTime(string, *time.Location)`
// functions are slightly modified versions of code from the github.com/go-sql-driver/mysql
// package. They work with Postgres and MySQL databases. Potential future
// drivers should ensure these will work for them, or come up with an alternative.
//
// Conforming with its licensing terms the copyright notice and link to the licence
// are available below.
//
// Source: https://github.com/go-sql-driver/mysql/blob/527bcd55aab2e53314f1a150922560174b493034/utils.go#L452-L508

// Copyright notice from original developers:
//
// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/

// Scan implements the Scanner interface.
// The value type must be time.Time or string / []byte (formatted time-string),
// otherwise Scan fails.
func (n *XTime) Scan(value interface{}) error {
	var err error

	if value == nil {
		n.Time, n.Valid = time.Time{}, false
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		n.Time, n.Valid = v, true
		return nil
	case []byte:
		n.Time, err = time.ParseInLocation(timeFormart, string(v), time.Local)
		n.Valid = (err == nil)
		return err
	case string:
		n.Time, err = time.ParseInLocation(timeFormart, v, time.Local)
		n.Valid = (err == nil)
		return err
	}

	n.Valid = false
	return nil
}

// UnmarshalJSON correctly deserializes a XTime from JSON.
func (n *XTime) UnmarshalJSON(b []byte) error {
	// scan for null
	if bytes.Equal(b, jsonNull) {
		return n.Scan(nil)
	}

	// scan for JSON string
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		return n.Scan(s)
	}

	// scan for JSON timestamp
	var t time.Time
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	return n.Scan(t)
}

// MarshalJSON 自定义时间类的序列化方法
func (n XTime) MarshalJSON() ([]byte, error) {
	if n.Valid {
		b := make([]byte, 0, len(timeFormart)+2)
		b = append(b, '"')
		b = n.Time.AppendFormat(b, timeFormart)
		b = append(b, '"')
		return b, nil
	}
	return jsonNull, nil
}

// String for X time
func (n XTime) String() string {
	return n.Time.Format(timeFormart)
}

// Value implements the driver Valuer interface.
func (n XTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}

// XInt64 兼容JSON与mysql 支持Null类型
type XInt64 struct {
	sql.NullInt64
}

// NewXInt64 creates a XInt64 with Scan().
func NewXInt64(v interface{}) (n XInt64) {
	n.Scan(v)
	return
}

// MarshalJSON correctly serializes a NullInt64 to JSON.
func (n XInt64) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Int64)
	}
	return jsonNull, nil
}

// UnmarshalJSON correctly deserializes a NullInt64 from JSON.
func (n *XInt64) UnmarshalJSON(b []byte) error {
	var s json.Number
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		return n.Scan(nil)
	}
	return n.Scan(s)
}

// XString 兼容JSON与Mysql支持Null类型
type XString struct {
	sql.NullString
}

// NewXString creates a XString with Scan().
func NewXString(v interface{}) (n XString) {
	n.Scan(v)
	return
}

// MarshalJSON correctly serializes a NullInt64 to JSON.
func (n XString) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.String)
	}
	return jsonNull, nil
}

// UnmarshalJSON correctly deserializes a NullInt64 from JSON.
func (n *XString) UnmarshalJSON(b []byte) error {
	var s interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return n.Scan(s)
}
