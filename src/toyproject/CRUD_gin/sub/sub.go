package sub

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// in-memory로 관리하기 위함
var userMap map[int]*User

// id 처리를 위한 변수
var lastId int

// User Struct
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"create_at"`
}

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// GET : User 리스트 조회
func usersHandler(c *gin.Context) {
	// 만약 등록된 유저가 한명도 없을 경우
	if len(userMap) == 0 {
		if err := c.ShouldBindUri(&userMap); err != nil {
			fmt.Printf("error - %+v \n", err)
		}
		fmt.Printf("No User - %+v \n", userMap)
		return
	}
	// 등록된 유저가 있을 경우
	users := []*User{}
	for _, u := range userMap {
		users = append(users, u)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   users,
	})
}

// POST
func createUserHandler(c *gin.Context) {
	// 유저 객체를 저장하기 위한 구조체 선언
	user := new(User)
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Printf("error - %+v \n", err)
		return
	}
	lastId++
	user.ID = lastId
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   user,
	})
}

// PUT
func updateUserHandler(c *gin.Context) {
	// 클라이언트에서 받은 내용을 updataUser 구조체에 바인딩한다.
	updateUser := new(User)
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		fmt.Printf("error - %+v \n", err)
		return
	}
	// 클라이언트에서 받은 ID값을 가진 유저가 기존 메모리에 등록되어 있는지 확인
	user, ok := userMap[updateUser.ID]
	if !ok {
		c.String(http.StatusOK, "No User ID : %d", updateUser.ID)
		return
	}

	// 업데이트된 항목만 변경
	if updateUser.FirstName != "" {
		user.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		user.LastName = updateUser.LastName
	}
	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	// c.String(http.StatusOK, "user : %s", user)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   user,
	})
}

func getUserHandler(c *gin.Context) {
	tempId := c.Param("id")
	id, err := strconv.Atoi(tempId)
	if err != nil {
		fmt.Printf("error - %+v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "This is not a valid value.",
		})
		return
	}

	user, ok := userMap[id]
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"status": "No User Id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"user":   user,
	})
}

// DELETE
func deleteUserHandler(c *gin.Context) {

	// id, err := strconv.Atoi(tempId)

	tempId := c.Param("id")
	id, err := strconv.Atoi(tempId)
	if err != nil {
		fmt.Printf("error - %+v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "This is not a valid value.",
		})
		return
	}
	_, ok := userMap[id]
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"status": "No User Id",
		})
		return
	}
	delete(userMap, id)
	c.JSON(http.StatusOK, gin.H{
		"status": "Deleted User ",
		"id":     id,
	})
}

// GET
func Index() *gin.Engine {
	// in memory 사용
	userMap = make(map[int]*User)
	// user 마다 다른 아이디를 생성해주기 위한 변수
	lastId = 0

	// gin 인스턴스 생성
	router := gin.New()

	// logger 미들웨어 사용
	router.Use(gin.Logger())
	// Recovery 미들웨어 사용
	router.Use(gin.Recovery())
	// CORS 사용
	router.Use(CORS())

	router.GET("/users", usersHandler)
	router.POST("/users", createUserHandler)
	router.PUT("/users", updateUserHandler)
	router.DELETE("/users/:id", deleteUserHandler)
	router.GET("/users/:id", getUserHandler)
	return router
}
