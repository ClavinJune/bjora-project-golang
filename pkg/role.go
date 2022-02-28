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

//go:generate stringer -type=Role -trimprefix=Role

package pkg

import "strings"

// Role define role enum
type Role int8

const (
	// RoleUndefined defines undefined role
	RoleUndefined Role = iota
	// RoleAdmin defines admin
	RoleAdmin
	// RoleMember defines member
	RoleMember
)

// RoleFrom parse str to Role
func RoleFrom(str string) Role {
	switch strings.ToUpper(strings.TrimSpace(str)) {
	case "ADMIN":
		return RoleAdmin
	case "MEMBER":
		return RoleMember
	}

	return RoleUndefined
}
