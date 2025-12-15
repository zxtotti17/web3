package models

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger     *zap.Logger
	Db         *gorm.DB
	PrivateKey string = "zx123456"
)

type CustomClaims struct {
	Uid      uint   `json:"uid"`      // 自定义字段：用户 ID
	UserName string `json:"username"` // 自定义字段：用户名
	Exp      int64  `json:"exp"`      // 自定义字段：过期时间
}

// Valid implements jwt.Claims.
func (c *CustomClaims) Valid() error {
	panic("unimplemented")
}

type User struct {
	gorm.Model
	UserName string `gorm:"index" json:"username" form:"username" binding:"required,min=2"`
	Email    string
	Password string    `gorm:"not null" json:"password" form:"password" binding:"required,contains=!@#"`
	PostNum  uint      `gorm:"default:0"`
	Posts    []Post    `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title      string    `gorm:"type:varchar(200);not null"`
	Content    string    `gorm:"type:text;not null"`
	UserID     uint      `gorm:"index"`
	CommentNum uint      `gorm:"default:0"`
	Comments   []Comment `gorm:"foreignKey:PostID"`
	User       User      `gorm:"foreignKey:UserID"`
}

type Comment struct {
	gorm.Model
	UserID uint `gorm:"index"`
	PostID uint `gorm:"index" json:"post_id" form:"post_id" binding:"required"`
	//IndexName string `gorm:"index:idx_comment_post_user"` // 联合索引
	Content string `gorm:"type:varchar(5000);not null" json:"content" form:"content" binding:"required"`
	User    User   `gorm:"foreignKey:UserID"`
	Post    Post   `gorm:"foreignKey:PostID"`
}

type MysqlCfg struct {
	Url          string
	Dbname       string
	MaxIdleConns int
	MaxOpenConns int
	Prefix       string
	Singular     bool
	Engine       string
}

type LoginResponse struct {
	User      CustomClaims `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type PostInfoMsg struct {
	PostID uint `json:"post_id" form:"post_id" binding:"required"`
}

type PostSetMsg struct {
	PostID  uint   `json:"post_id" form:"post_id" binding:"required"`
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
}

type ListMsg struct {
	Page uint `json:"page" form:"page" binding:"required"`
	Size uint `json:"size" form:"size" binding:"required"`
}

type PostListRes struct {
	Total int64  `json:"total"`
	List  []Post `json:"list"`
}

type CommentListMsg struct {
	PostID uint `json:"post_id" form:"post_id" binding:"required"`
	Page   uint `json:"page" form:"page" binding:"required"`
	Size   uint `json:"size" form:"size" binding:"required"`
}

type CommentListRes struct {
	Total  int64     `json:"total"`
	List   []Comment `json:"list"`
	PostID uint      `json:"post_id" form:"post_id" binding:"required"`
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
