package model

import "github.com/golang-jwt/jwt"

type GetTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AppKey   string `json:"app_key"`
}

type GetTokenResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Data    Token  `json:"data,omitempty"`
}

type Token struct {
	Token string `json:"token"`
}

type JwtCustomClaims struct {
	UserID         int64  `json:"user_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	BranchUserType int32  `json:"branch_user_type"`
	BranchTypeID   int32  `json:"branch_type_id"`
	MksUserID      int64  `json:"mks_user_id"`
	jwt.StandardClaims
}
