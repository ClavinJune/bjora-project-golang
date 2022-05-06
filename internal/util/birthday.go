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
	"time"
)

const (
	birthdayTimeLayout string = "2006-01-02"
)

// BirthdayFromStr parses string to time.Time using birthdayTimeLayout
func BirthdayFromStr(str string) (time.Time, error) {
	t, err := time.Parse(birthdayTimeLayout, str)
	if err != nil {
		return time.Time{}, WrapError(err)
	}

	return t, nil
}

// BirthdayFromTime parses time to string using birthdayTimeLayout
func BirthdayFromTime(t time.Time) string {
	return t.Format(birthdayTimeLayout)
}
