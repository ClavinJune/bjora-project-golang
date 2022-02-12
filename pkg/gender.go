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

package pkg

import "strings"

// Gender define gender enum
type Gender string

const (
	// GenderMale defines male
	GenderMale Gender = "male"
	// GenderFemale defines female
	GenderFemale Gender = "female"
	// GenderUndefined defines undefined gender
	GenderUndefined Gender = "gender_undefined"
)

// GenderFrom parse str to Gender
func GenderFrom(str string) Gender {
	switch strings.ToUpper(strings.TrimSpace(str)) {
	case "MALE":
		return GenderMale
	case "FEMALE":
		return GenderFemale
	}

	return GenderUndefined
}
