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

//go:generate stringer -type=Environment -trimprefix=Environment

package enum

import "strings"

// Environment define environment enum
type Environment int8

const (
	// EnvironmentUndefined defines undefined gender
	EnvironmentUndefined Environment = iota
	// EnvironmentTest defines test
	EnvironmentTest
	// EnvironmentDev defines dev
	EnvironmentDev
	// EnvironmentStg defines stg
	EnvironmentStg
	// EnvironmentProd defines prod
	EnvironmentProd
)

// EnvironmentFrom parse str to Gender
func EnvironmentFrom(str string) Environment {
	switch strings.ToUpper(strings.TrimSpace(str)) {
	case "TEST":
		return EnvironmentTest
	case "DEV":
		return EnvironmentDev
	case "STG":
		return EnvironmentStg
	case "PROD":
		return EnvironmentProd
	}

	return EnvironmentUndefined
}
