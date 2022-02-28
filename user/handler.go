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
	"net/http"
	"strings"

	"github.com/clavinjune/bjora-project-golang/pkg"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc    pkg.UserService
	router fiber.Router
}

func (h *handler) ApplyRoute(router fiber.Router) {
	router.Get("/", h.fetchByEmail())
}

func (h *handler) fetchByEmail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Query("email")
		if strings.TrimSpace(email) == "" {
			return c.SendStatus(http.StatusNotFound)
		}

		user, err := h.svc.FetchByEmail(c.Context(), email)
		if err != nil {
			return c.SendStatus(http.StatusNotFound)
		}

		return c.JSON(user)
	}
}
