package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"net/http"

	"github.com/kittanutp/salesrecorder/database"
	"github.com/kittanutp/salesrecorder/schema"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func (db *DBController) CreateUser(c *gin.Context) {
	var user database.User

	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	hashPassword := sha256.New()
	hashPassword.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hashPassword.Sum((nil)))

	res := db.Database.Create(&user)

	if res.Error != nil {
		log.Fatalf("Unable to execute the query. %v", res.Error)
	}

	c.JSON(http.StatusCreated, gin.H{"username": user.Username, "id": user.ID, "created_at": user.CreatedAt})
}

func (db *DBController) LogIn(c *gin.Context) {
	var requestedUser schema.UserRequest
	if err := c.BindJSON(&requestedUser); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var user = database.User{Username: requestedUser.Username}
	db.Database.First(&user)
	hashPassword := sha256.New()
	hashPassword.Write([]byte(requestedUser.Password))
	requestedUser.Password = hex.EncodeToString(hashPassword.Sum((nil)))
	if requestedUser.Password != user.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Wrong username or password"})
		return
	} else {
		token := GenerateToken(user.Username, true)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (db *DBController) GetUserInfo(c *gin.Context) {
	user, err := GetUserFromCtx(c)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": user.Username, "id": user.ID, "created_at": user.CreatedAt})
}

func GetUserUsername(db *gorm.DB, username string) (*database.User, error) {
	user := database.User{}
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserFromCtx(c *gin.Context) (*database.User, error) {
	userData, exist := c.Get("user")
	if !exist {
		c.AbortWithStatus(400)
		return nil, errors.New("unable to get user from ctx")
	}
	user, ok := userData.(*database.User)
	if !ok {
		c.AbortWithStatus(400)
		return nil, errors.New("invalid user format")
	}
	return user, nil
}
