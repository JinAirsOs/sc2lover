// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package battlenet provides constants for using OAuth2 to access battlenet.
package battlenet // import "golang.org/x/oauth2/battlenet"

import (
	"golang.org/x/oauth2"
)

// Endpoint is battlenet's OAuth 2.0 endpoint.
// US, EU, APAC	https://<region>.battle.net/oauth/authorize	https://<region>.battle.net/oauth/token
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.battlenet.com.cn/oauth/authorize",
	TokenURL: "https://www.battlenet.com.cn/oauth/token",
}
