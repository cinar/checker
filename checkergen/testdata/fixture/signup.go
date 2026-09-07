// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package fixture

// SignupRequest exercises a representative mix of checkergen-eligible
// checkers/normalizers, including field-relative ones.
type SignupRequest struct {
	Email           string `json:"email" checkers:"trim lower required email"`
	Password        string `json:"password" checkers:"required min-len:8"`
	ConfirmPassword string `json:"confirm_password" checkers:"required eq-field:Password"`
	Age             int    `json:"age" checkers:"gte:18"`
	Nickname        string `json:"nickname" checkers:"omitempty min-len:3"`
}
