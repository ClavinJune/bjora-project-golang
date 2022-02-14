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

package errutil

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

type Opt struct {
	WithCaller bool
	Message    string
}

type OptFunc func(*Opt)

func WithCaller() OptFunc {
	return func(opt *Opt) {
		opt.WithCaller = true
	}
}

func WithMessage(msg string) OptFunc {
	return func(opt *Opt) {
		opt.Message = msg
	}
}

// New wraps error
func New(err error, opts ...OptFunc) error {
	o := &Opt{}
	for _, opt := range opts {
		opt(o)
	}

	var caller string

	if o.WithCaller {
		if pc, _, line, ok := runtime.Caller(0); ok {
			fn := regexp.
				MustCompile(`\\.func\\d$`).
				Split(runtime.FuncForPC(pc).Name(), -1)[0]
			caller = fmt.Sprintf("%v:%v", fn, line)
		}
	}

	prefix := strings.Join([]string{caller, o.Message}, " - ")
	return fmt.Errorf("%s: %w", prefix, err)
}
