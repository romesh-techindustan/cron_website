package controllers

import (
	"fmt"
	"goapi/database"
	"goapi/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
func AddWebsites(c *gin.Context) {
	var website models.Websites
	name := c.PostForm("name")
	web := c.PostForm("web")
	website.Name = name
	website.Web = web
}

func AddSOSUser(c *gin.Context) {
	var user models.SOS_User
	email := c.PostForm("email")
	user.Email = email

}

func GetAllWebsites(c *gin.Context) {
	var web models.Websites

	id := c.Param("id")
	err := database.Init().QueryRow("select * from websites where id=?", id)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Invalid id"})
	}
	c.JSON(http.StatusOK, &web)

}

func GetSOSUsername(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := database.Init().QueryRow("select * from sos_users where id=?", id)
	if err == nil {
		c.JSON(200, &user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func StatusCode(c *gin.Context) {
	var status models.Status_code
	id := c.Param("id")
	err := database.Init().QueryRow("select * from status_code where id=?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not a valid ID"})
	} else {
		c.JSON(http.StatusOK, &status)
	}
	return
}

func PostStatusCode(c *gin.Context) {
	var status models.Status_code
	var web_id = c.PostForm("id") 
	code := c.PostForm("code")
	status.Web_id = web_id
	status.Code = code

	err:= database.Init().QueryRow("insert into status_code (web_id, code) values (?,?)", status.Web_id, status.Code)
	if err!= nil{

	}

}

func Login(c *gin.Context) {
	var user models.User
	email := c.PostForm("email")
	password := c.PostForm("password")

	user.Email = email
	var storedHashedPassword string
	err := database.Init().QueryRow("select * from flower where id = ?;", user.Email).Scan(&storedHashedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or Password"})
	}
	fmt.Println(err)

	match := comparePasswords(storedHashedPassword, []byte(password))
	if match != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
}

func SignUp(c *gin.Context) {
	var user models.User
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	user.Name = name
	user.Email = email

	hashedPassword := hashPassword([]byte(password))
	user.HashPassword = hashedPassword

	err := database.Init().QueryRow("INSERT INTO users (name, email, password) VALUES (?,?,?,?);", user.Name, user.Email, user.HashPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot insert into table"})
	}

	// c.JSON(http.StatusOK, gin.H{"email": user.Email,"username":user.Name})
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Main website",
	})
}

func hashPassword(password []byte) string {

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
