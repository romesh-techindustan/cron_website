package controllers

// import (
// 	"goapi/database"
// 	"goapi/models"
// 	"log"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func GetUser(c *gin.Context) {
// 	var user []models.User
// 	err := database.Init().Select(&user, "select * from user")
// 	if err == nil {
// 		c.JSON(200, user)
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}
// }
// func GetUserDetail(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user models.User
// 	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=? LIMIT 1", id)
// 	if err == nil {
// 		user_id, _ := strconv.ParseInt(id, 0, 64)
// 		content := &models.User{
// 			Id:         user_id,
// 			Username:   user.Username,
// 			Email:      user.Email,
// 			Created_at: user.Created_at,
// 			Is_active:  false,
// 			Is_admin:   false,
// 		}
// 		c.JSON(200, content)
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}
// }
// func UsernameViaId(c *gin.Context) {
// 	id:= c.Params.ByName("id")
// 	var name models.User
// 	err := dbmap.SelectOne(&name, "SELECT Username FROM user WHERE Id=?", id)
// 	if err == nil{
// 		user_id,_:= strconv.ParseInt(id,0,64)
// 		content :=&models.User{
// 			Id: user_id,
// 			Username: name.Username,
// 		}
// 		c.JSON(200, content)
// 	} else {
// 		c.JSON(404, gin.H{"error":"Unidentified user"})

// 	}

// }
// func PostUser(c *gin.Context) {
// 	var user models.User
// 	c.Bind(&user)
// 	log.Println(user)
// 	if user.Username != "" && user.Password != "" && user.Email != "" && user.Created_at != "" {
// 		if insert, err := dbmap.Exec(`INSERT INTO user ( Id,Username, Password, Email, Created_at) VALUES (?,?, ?, ?, ?)`, user.Id, user.Username, user.Password, user.Email, user.Created_at); insert != nil {
// 			// user_id, err := insert.LastInsertId()
// 			if err == nil {
// 				content := &models.User{
// 					Id:        user.Id,
// 					Username:  user.Username,
// 					Email:      user.Email,
// 					Password:  user.Password,
// 					Created_at: user.Created_at,
// 				}
// 				c.JSON(201, content)
// 			} else {
// 				checkErr(err, "Insert failed")
// 			}
// 		}
// 	} else {
// 		c.JSON(400, gin.H{"error": "Fields are empty"})
// 	}
// }
// func UpdateUser(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user models.User
// 	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
// 	if err == nil {
// 		var json models.User
// 		c.Bind(&json)
// 		user_id, _ := strconv.ParseInt(id, 0, 64)
// 		user := models.User{
// 			Id:        user_id,
// 			Username:  user.Username,
// 			Email:      user.Email,
// 			Password:  user.password,
// 			Created_at: user.Created_at,
// 		}
// 		if user.Email != "" && user.Password != "" {
// 			_, err = dbmap.Update(&user)
// 			if err == nil {
// 				c.JSON(200, user)
// 			} else {
// 				checkErr(err, "Updated failed")
// 			}
// 		} else {
// 			c.JSON(400, gin.H{"error": "fields are empty"})
// 		}
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}
// }

// func jwtGenerateToken(m *AuthorizationModel) (*jwtObj, error) {
// 	m.Password = ""
// 	expireAfterTime := time.Hour * time.Duration(viper.GetInt("app.jwt_expire_hour"))
// 	iss := viper.GetString("app.name")
// 	appSecret := viper.GetString("app.secret")
// 	expireTime := time.Now().Add(expireAfterTime)
// 	stdClaims := jwt.StandardClaims{
// 		ExpiresAt: expireTime.Unix(),
// 		IssuedAt:  time.Now().Unix(),
// 		Id:        fmt.Sprintf("%d", m.Id),
// 		Issuer:    iss,
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, stdClaims)
// 	// Sign and get the complete encoded token as a string using the secret
// 	tokenString, err := token.SignedString([]byte(appSecret))
// 	if err != nil {
// 		logrus.WithError(err).Fatal("config is wrong, can not generate jwt")
// 	}
// 	data := &jwtObj{AuthorizationModel: *m, Token: tokenString, Expire: expireTime, ExpireTs: expireTime.Unix()}
// 	return data, err
// }

// type jwtObj struct {
// 	AuthorizationModel
// 	Token    string    `json:"token"`
// 	Expire   time.Time `json:"expire"`
// 	ExpireTs int64     `json:"expire_ts"`
// }

// //JwtParseUser parse a jwt token and return an authorized identity
// func JwtParseUser(tokenString string) (*AuthorizationModel, error) {
// 	if tokenString == "" {
// 		return nil, errors.New("token is not found in Authorization Bearer")
// 	}
// 	claims := jwt.StandardClaims{}
// 	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		secret := viper.GetString("app.secret")
// 		return []byte(secret), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if claims.VerifyExpiresAt(time.Now().Unix(), true) == false {
// 		return nil, errors.New("token is expired")
// 	}
// 	appName := viper.GetString("app.name")
// 	if !claims.VerifyIssuer(appName, true) {
// 		return nil, errors.New("token's issuer is wrong,greetings Hacker")
// 	}
// 	key := fmt.Sprintf("login:%s", claims.Id)
// 	jwtObj, err := mem.GetJwtObj(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &jwtObj.AuthorizationModel, err
// }

// //get an authorized user form memory store
// func (s *memoryStore) GetJwtObj(id string) (value *jwtObj, err error) {
// 	vv, err := s.Get(id, false)
// 	if err != nil {
// 		return nil, err
// 	}
// 	value, ok := vv.(*jwtObj)
// 	if ok {
// 		return value, nil
// 	}
// 	return nil, errors.New("mem:has value of this id, but is not type of *jwtObj")
// }
