// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package oauth2 // import "miniflux.app/v2/internal/oauth2"

import (
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/oauth2"

	"miniflux.app/v2/internal/crypto"
)

type Authorization struct {
	url          string
	state        string
	codeVerifier string
}

func (u *Authorization) RedirectURL() string {
	return u.url
}

func (u *Authorization) State() string {
	return u.state
}

func (u *Authorization) CodeVerifier() string {
	return u.codeVerifier
}

func GenerateAuthorization(config *oauth2.Config) *Authorization {
	codeVerifier := crypto.GenerateRandomStringHex(32)
	sum := sha256.Sum256([]byte(codeVerifier))

	state := crypto.GenerateRandomStringHex(24)

	authUrl := config.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", base64.RawURLEncoding.EncodeToString(sum[:])),
	)

	return &Authorization{
		url:          authUrl,
		state:        state,
		codeVerifier: codeVerifier,
	}
}
