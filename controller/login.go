package controllers

import (
	"goapi/models"
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetUser(c *gin.Context) {
	var user []models.User
	_, err := dbmap.Select(&user, "select * from user")
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}
func GetUserDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=? LIMIT 1", id)
	if err == nil {
		user_id, _ := strconv.ParseInt(id, 0, 64)
		content := &models.User{
			Id:         user_id,
			Username:   user.Username,
			Email:      user.Email,
			Created_at: user.Created_at,
			Is_active:  false,
			Is_admin:   false,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}
func UsernameViaId(c *gin.Context) {
	id:= c.Params.ByName("id")
	var name models.User
	err := dbmap.SelectOne(&name, "SELECT Username FROM user WHERE Id=?", id)
	if err == nil{
		user_id,_:= strconv.ParseInt(id,0,64)
		content :=&models.User{
			Id: user_id,
			Username: name.Username,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error":"Unidentified user"})

	}

}
func PostUser(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	log.Println(user)
	if user.Username != "" && user.Password != "" && user.Email != "" && user.Created_at != "" {
		if insert, err := dbmap.Exec(`INSERT INTO user ( Id,Username, Password, Email, Created_at) VALUES (?,?, ?, ?, ?)`, user.Id, user.Username, user.Password, user.Email, user.Created_at); insert != nil {
			// user_id, err := insert.LastInsertId()
			if err == nil {
				content := &models.User{
					Id:        user.Id,
					Username:  user.Username,
					Email:      user.Email,
					Password:  user.Password,
					Created_at: user.Created_at,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
	if err == nil {
		var json models.User
		c.Bind(&json)
		user_id, _ := strconv.ParseInt(id, 0, 64)
		user := models.User{
			Id:        user_id,
			Username:  user.Username,
			Email:      user.Email,
			Password:  user.password,
			Created_at: user.Created_at,
		}
		if user.Email != "" && user.Password != "" {
			_, err = dbmap.Update(&user)
			if err == nil {
				c.JSON(200, user)
			} else {
				checkErr(err, "Updated failed")
			}
		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}


