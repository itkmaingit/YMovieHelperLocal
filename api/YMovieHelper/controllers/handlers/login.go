package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/utils"
)

//	@Summary		ログイン認証
//	@Description	ログイン認証自体はFirebaseが担うので、SessionにUIDをセットするだけ
//	@Tags			users
//	@Produce		plain
//	@Param			user	query		string	true	"UID"
//	@Success		200		{string}	string	"Succeeded set session."
//	@Router			/login [post]
func Login(c *gin.Context) {
	type Request struct {
		UID string `json:"uid" example:"fjkoafjklajklfjaefwjiowa"`
	}
	var req Request
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.Login: Failed to bind request data: %v", err)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &utils.Claims{
		UserID: req.UID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(os.Getenv("JWTKey"))

	// Create the JWT token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT token."})
		log.Printf("handlers.Login: %v", err)
		return
	}

	// cookie := http.Cookie{
	// 	Name:     "token",
	// 	Value:    tokenString,
	// 	Expires:  time.Now().Add(24 * time.Hour), // 有効期限は24時間
	// 	Path:     "/",                            // パスはルート
	// 	Domain:   os.Getenv("FrontendDomain"),    // ドメインは example.com
	// 	Secure:   true,                           // HTTPSのみで送信
	// 	HttpOnly: false,                          // JavaScriptからアクセス不可
	// 	SameSite: http.SameSiteNoneMode,
	// }
	// http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded set cookie.", "token": tokenString})
}

func Auth(c *gin.Context) {
	// JWTトークンが不正なら、indexページにリダイレクトさせる

	jwtKey := []byte(os.Getenv("JWTKey"))
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	claims := &utils.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			log.Printf("middlewares.AccessControllMiddleware: %v", err)
			return
		}
		log.Printf("middlewares.AccessControllMiddleware: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if !token.Valid {
		log.Printf("middlewares.AccessControllMiddleware: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success Authorization."})
}
