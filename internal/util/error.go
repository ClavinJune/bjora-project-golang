// Copyright 2022 ClavinJune/bjora
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// Error is a custom error
type Error struct {
	Err      error
	FileLine string
	FuncName string
	Message  string
}

// Unwrap returns internal error
func (e *Error) Unwrap() error {
	return e.Err
}

// Error returns error message
func (e *Error) Error() string {
	var msg string

	var err *Error
	if errors.As(e.Err, &err) {
		msg = errFormat(err.FuncName, err.FileLine, e.Message, e.Err.Error())
	} else {
		msg = fmt.Sprintf("%q", e.Err.Error())
	}

	return errFormat(e.FuncName, e.FileLine, e.Message, msg)
}

func errFormat(fn, fl, msg, err string) string {
	var b strings.Builder
	b.WriteString(`{"funcname":"`)
	b.WriteString(fn)
	b.WriteString(`","fileline":"`)
	b.WriteString(fl)
	b.WriteString(`","caused":`)
	b.WriteString(err)

	if strings.TrimSpace(msg) == "" {
		b.WriteString(`}`)
	} else {
		b.WriteString(`,"message":"`)
		b.WriteString(msg)
		b.WriteString(`"}`)
	}
	return b.String()
}

// WrapError creates *Error by wrapping error
func WrapError(err error) *Error {
	fl, fn := "?", "?"
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fl = errConstructFileLine(file, line)
		fn = errGetFuncName(pc)
	}

	return &Error{
		Err:      err,
		FileLine: fl,
		FuncName: fn,
	}
}

// NewErrorFromMsg creates *Error from string
func NewErrorFromMsg(msg string) *Error {
	fl, fn := "?", "?"
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fl = errConstructFileLine(file, line)
		fn = errGetFuncName(pc)
	}

	return &Error{
		Err:      errors.New(msg),
		FileLine: fl,
		FuncName: fn,
	}
}

// WrapErrorWithMsg creates *Error by wrapping error with a custom message
func WrapErrorWithMsg(err error, msg string) error {
	fl, fn := "?", "?"
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fl = errConstructFileLine(file, line)
		fn = errGetFuncName(pc)
	}

	return &Error{
		Err:      err,
		FileLine: fl,
		FuncName: fn,
		Message:  msg,
	}
}

func errConstructFileLine(file string, line int) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	return fmt.Sprintf("%s:%d", short, line)
}

func errGetFuncName(pc uintptr) string {
	f := runtime.FuncForPC(pc)

	split := strings.Split(f.Name(), "/")
	return split[len(split)-1]
}
