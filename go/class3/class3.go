package class3

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Gorm定义 User 、 Post 和 Comment 模型
type User struct {
	gorm.Model
	Name     string `gorm:"index"`
	Email    string
	Password string
	PostNum  uint      `gorm:"default:0"`
	Posts    []Post    `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(200);not null"`
	Description string    `gorm:"type:text;not null"`
	UserID      uint      `gorm:"index"`
	CommentNum  uint      `gorm:"default:0"`
	Comments    []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	gorm.Model
	UserID uint `gorm:"index"`
	PostID uint `gorm:"index"`
	//IndexName string `gorm:"index:idx_comment_post_user"` // 联合索引
	Content string
}

// 使用Gorm创建这些模型对应的数据库表
func Class3_1() {
	dns := "root:123456@tcp(localhost:3306)/zm4?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}

	// 创建 User 表
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	u1 := User{Name: "Alice", Email: "alice@example.com", Password: "123456"}
	u2 := User{Name: "Bob", Email: "bob@example.com", Password: "654321"}
	db.Omit("PostNum").Create(&u1)
	db.Omit("PostNum").Create(&u2)

	// 创建 Post 表
	p1 := Post{Title: "Post 1", Description: "This is the first post", UserID: u1.ID}
	p2 := Post{Title: "Post 2", Description: "This is the second post", UserID: u2.ID}
	p3 := Post{Title: "Post 3", Description: "This is the first post", UserID: u1.ID}
	db.Omit("CommentNum").Create(&p1)
	db.Omit("CommentNum").Create(&p2)
	db.Omit("CommentNum").Create(&p3)

	// 创建 Comment 表
	c1 := Comment{PostID: p1.ID, UserID: u1.ID, Content: "Good job!"}
	c2 := Comment{PostID: p1.ID, UserID: u2.ID, Content: "Nice post!"}
	c3 := Comment{PostID: p2.ID, UserID: u1.ID, Content: "Great!"}
	db.Create(&c1)
	db.Create(&c2)
	db.Create(&c3)

	sqlDB, err := db.DB()
	if err != nil {
		// 处理错误
		fmt.Println("insert database error")
		return
	}
	defer sqlDB.Close()
}

// 使用Gorm查询某个用户发布的所有文章及其对应的评论信息
func Class3_2() {
	dns := "root:123456@tcp(localhost:3306)/zm4?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
	}

	// 查询某个用户发布的所有文章及其对应的评论信息
	var result User
	err = db.Preload("Posts.Comments").First(&result, User{Name: "Alice"}).Error
	if err != nil {
		fmt.Println("Failed to query user")
		return
	}

	fmt.Printf("User: %s (ID: %d)\n", result.Name, result.ID)
	for _, post := range result.Posts {
		fmt.Printf("  Post: %s\n", post.Title)
		for _, comment := range post.Comments {
			fmt.Printf("    Comment: %s\n", comment.Content)
		}
	}

	//使用Gorm查询评论数量最多的文章信息
	var postMax Post
	db.Model(&Post{}).Select("posts.*, COUNT(comments.id) AS count").Joins("left join comments on posts.id = comments.post_id").Group("posts.id").Order("count DESC").Limit(1).First(&postMax)
	fmt.Println(postMax.Title, postMax.CommentNum, postMax.UserID, postMax.Comments)

	sqlDB, err := db.DB()
	if err != nil {
		// 处理错误
		fmt.Println("select database error")
		return
	}
	defer sqlDB.Close()
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	// 增加用户的文章计数
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("PostNum", gorm.Expr("post_num + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论创建时自动更新文章的评论数量统计字段
func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	// 增加文章的评论计数
	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).UpdateColumn("CommentNum", gorm.Expr("comment_num + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}
