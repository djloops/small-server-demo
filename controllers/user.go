package controllers

import (
    "github.com/djloops/small-server-demo/models"

    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    "github.com/djloops/small-server-demo/db"
	"time"
	"math/rand"
	"fmt"
	"strconv"
)

func CreateUser(c *gin.Context) {

    db := c.MustGet("DB").(*gorm.DB)
    
    var user models.User
    if c.BindJSON(&user) != nil {
        c.JSON(406, gin.H{"message": "Invalid data", "data": user})
        c.Abort()
        return
    }

    if !db.NewRecord(user) {
        c.JSON(406, gin.H{"message": "Todo could not be created"})
        c.Abort()
        return
        
    }
    err := db.Create(&user).Error
    if err != nil {
    }
    c.JSON(200, gin.H{"message": "Todo created"})
}

func GetBoard(c *gin.Context) {
	board, _ := strconv.ParseInt(c.Query("board"), 10, 64)
	offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)
	count, _ := strconv.ParseInt(c.Query("count"), 10, 32)

	res := []models.FeedCard{}

	posts, _ := db.GetPosts(int8(board), offset, int32(count))
	for _, post := range posts {
		user, _ := db.GetUser(post.OwnerID)
		commentCount, _ := db.GetCommentCount(post.ID)

		res = append(res, models.FeedCard{
			User: *user,
			Post: post,
			CommentCount: *commentCount,
		})
	}


	c.JSON(200, gin.H{"feed_cards": res})
}

func Mock(c * gin.Context) {
    db.CreateUser(&models.User{
    	Name: "dongjiali",
		AvatarUrl:  "https://source.unsplash.com/random/100x100",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})

	db.CreateUser(&models.User{
		Name: "jingxue",
		AvatarUrl:  "https://source.unsplash.com/random/100x100",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "曾国藩",
		AvatarUrl:  "https://source.unsplash.com/random/100x100",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "Gary Dong",
		AvatarUrl:  "https://source.unsplash.com/random/100x100",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "曹操",
		AvatarUrl:  "https://source.unsplash.com/random/100x100",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "张飞",
		AvatarUrl:  "https://source.unsplash.com/random/100x100",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "陈坤",
		AvatarUrl:  "https://source.unsplash.com/random/100x100",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "马化腾",
		AvatarUrl:  "https://source.unsplash.com/random/100x100",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "Steve jobs",
		AvatarUrl:  "https://source.unsplash.com/random/100x100",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "欧阳锋",
		AvatarUrl:  "https://source.unsplash.com/random/100x100",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "姆巴佩",
		AvatarUrl:  "https://source.unsplash.com/random/100x100",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})


    for i:= 0; i < 1000; i++ {
		db.CreatePost(&models.Post{
			OwnerID:    int64(rand.Intn(10) + 1),
			Board: int8(rand.Intn(3) + 1),
			Subject:    fmt.Sprintf("测试标题 %v", i),
			Content:    fmt.Sprintf("这是一个测试内容 %v", i),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})
	}

	for i:= 0; i < 10000; i++ {
		db.CreateComment(&models.Comment{
			PostID:    int64(rand.Intn(1000) + 1),
			Content:    fmt.Sprintf("这是一个评论内容 %v", i),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})
	}
}

//func DeleteTodo(c *gin.Context) {
//    db := c.MustGet("DB").(*gorm.DB)
//    id := c.Param("id")
//    var todo models.Todo
//    db.First(&todo, id)
//    db.Delete(&todo)
//    c.JSON(200, gin.H{"message": "ok"})
//}
//
//func ListTodo(c *gin.Context) {
//    db := c.MustGet("DB").(*gorm.DB)
//    todos := []models.Todo{}
//    db.Find(&todos)
//    c.JSON(200, gin.H{"todos": &todos})
//}
