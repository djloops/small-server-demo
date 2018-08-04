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
    	Name: "董佳礼",
		AvatarUrl:  "https://images.unsplash.com/photo-1531203236206-56458355ba2a?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&h=200&fit=crop&ixid=eyJhcHBfaWQiOjF9&s=885060889caf3f1ba98c70a0c85b25b8",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})

	db.CreateUser(&models.User{
		Name: "zoujing",
		AvatarUrl:  "https://images.unsplash.com/photo-1530890192129-e7c9ec4a2025?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&h=200&fit=crop&ixid=eyJhcHBfaWQiOjF9&s=d035ec770d0222ca956c30614059e096",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "韩速",
		AvatarUrl:  "https://images.unsplash.com/photo-1532578941548-30cbde3ad4b1?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&h=200&fit=crop&ixid=eyJhcHBfaWQiOjF9&s=e5a0841bc2bb2488f08edcf1beda760e",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "Steve Jobs",
		AvatarUrl:  "https://images.unsplash.com/photo-1532696389727-73b88de71ef0?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&h=200&fit=crop&ixid=eyJhcHBfaWQiOjF9&s=2f5c05b5dead472391dceb93e9de1c4e",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "张超",
		AvatarUrl:  "https://images.unsplash.com/photo-1530999075172-6846540b2617?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&h=200&fit=crop&ixid=eyJhcHBfaWQiOjF9&s=d7dae86cc8dd72d19ceb3c0dab53f614",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "张飞",
		AvatarUrl:  "https://images.unsplash.com/photo-1532466522623-976ab79e4037?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&h=200&fit=crop&ixid=eyJhcHBfaWQiOjF9&s=8e3aa877123f0c7852221828fbf20512",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "陈坤",
		AvatarUrl:  "https://images.unsplash.com/photo-1533165159663-884a2837ef63?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&h=200&fit=crop&ixid=eyJhcHBfaWQiOjF9&s=4584eeb6e64c6f7e529f36a3d8dd6182",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "马化腾",
		AvatarUrl:  "https://images.unsplash.com/photo-1531579881625-f9ab707607fe?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&h=200&fit=crop&ixid=eyJhcHBfaWQiOjF9&s=d3e9ddd353a28fe905f273645b578312",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "好红叶",
		AvatarUrl:  "https://images.unsplash.com/photo-1532090912514-3b90d608e6b7?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&h=200&fit=crop&ixid=eyJhcHBfaWQiOjF9&s=f5d469bc21c7749414cc4d5b9f552335",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	db.CreateUser(&models.User{
		Name: "姆巴佩",
		AvatarUrl:  "https://images.unsplash.com/photo-1533052420930-2cda210f54ed?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&h=200&fit=crop&ixid=eyJhcHBfaWQiOjF9&s=c476b8ae13e9f0f8141b34aacdb71340",
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
