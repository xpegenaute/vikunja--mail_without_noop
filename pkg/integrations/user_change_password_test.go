// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-present Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public Licensee as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public Licensee for more details.
//
// You should have received a copy of the GNU Affero General Public Licensee
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package integrations

import (
	"net/http"
	"testing"

	apiv1 "code.vikunja.io/api/pkg/routes/api/v1"
	"code.vikunja.io/api/pkg/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserChangePassword(t *testing.T) {
	t.Run("Normal test", func(t *testing.T) {
		rec, err := newTestRequestWithUser(t, http.MethodPost, apiv1.UserChangePassword, &testuser1, `{
  "new_password": "123456789",
  "old_password": "12345678"
}`, nil, nil)
		require.NoError(t, err)
		assert.Contains(t, rec.Body.String(), `The password was updated successfully.`)
	})
	t.Run("Wrong old password", func(t *testing.T) {
		_, err := newTestRequestWithUser(t, http.MethodPost, apiv1.UserChangePassword, &testuser1, `{
  "new_password": "12345",
  "old_password": "invalid"
}`, nil, nil)
		require.Error(t, err)
		assertHandlerErrorCode(t, err, user.ErrCodeWrongUsernameOrPassword)
	})
	t.Run("Empty old password", func(t *testing.T) {
		_, err := newTestRequestWithUser(t, http.MethodPost, apiv1.UserChangePassword, &testuser1, `{
  "new_password": "12345",
  "old_password": ""
}`, nil, nil)
		require.Error(t, err)
		assertHandlerErrorCode(t, err, user.ErrCodeEmptyOldPassword)
	})
	t.Run("Empty new password", func(t *testing.T) {
		_, err := newTestRequestWithUser(t, http.MethodPost, apiv1.UserChangePassword, &testuser1, `{
  "new_password": "",
  "old_password": "12345678"
}`, nil, nil)
		require.Error(t, err)
		assertHandlerErrorCode(t, err, user.ErrCodeEmptyNewPassword)
	})
}
