package dtos

import (
	"encoding/json"
	"go-store/cmd/models"
)

type Auth struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type UserToken struct {
	Token string `form:"token" json:"token" xml:"token"  binding:"required"`
	User  struct {
		UserName string `form:"username" json:"username" xml:"username" binding:"required"`
		ID       uint   `form:"id" json:"id" xml:"id" binding:"required"`
	} `form:"user" json:"user" xml:"user" binding:"required"`
}

func UserDetails(user *models.User) map[string]interface{} {
	data := UserToken{Token: user.SignClaims()}
	data.User.UserName = user.Username
	data.User.ID = user.ID
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(data)
	json.Unmarshal(inrec, &inInterface)
	return Success(inInterface)
}
