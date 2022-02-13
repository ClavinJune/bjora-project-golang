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

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/clavinjune/bjora-project-golang/pkg"
	"github.com/clavinjune/bjora-project-golang/pkg/handlerutil"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	svc pkg.UserService
}

func (h *handler) Store() (string, httprouter.Handle) {

	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer handlerutil.CloseRequest(r)

		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var b RequestStore
		if err := d.Decode(&b); err != nil {
			log.Println("here", err)

			handlerutil.Response().
				Error(err).
				Write(w)
			return
		}

		handlerutil.Response().Error(fmt.Errorf("ehehehe")).Write(w)
	}

	return "/user", handle
}

func (h *handler) FetchByEmail() httprouter.Handle {
	panic("implement me")
}
