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

package user

// RequestStore defines attribute for Storing User
type RequestStore struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	Address        string `json:"address"`
	Birthday       string `json:"birthday"`
	ProfilePicture []byte `json:"profile_picture"`
}
