package user

import (
	"net/http"
	"strconv"

	"class4_gin/models"
	"class4_gin/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 绑定 JSON

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"param error": err.Error()})
		return
	}
	if err := models.Db.First(models.User{UserName: user.UserName}); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := models.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser models.User
	if err := models.Db.Where("user_name = ?", user.UserName).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	// 剩下的逻辑...
	returnUser, tokenString := utils.LoginToken(storedUser, c)
	c.JSON(http.StatusOK, models.LoginResponse{
		User:  returnUser,
		Token: tokenString,
	})
}

func DddPost(c *gin.Context) {
	value, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claims, ok := value.(*models.CustomClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var addPost models.AddPostMsg
	if err := c.ShouldBindJSON(&addPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:   addPost.Title,
		Content: addPost.Content,
		UserID:  claims.Uid,
	}

	if err := models.Db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": post})
}

func GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var postInfo models.Post
	if err := models.Db.Where("id = ?", id).First(&postInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
		return
	}

	c.JSON(http.StatusOK, postInfo)
}

func GetPostList(c *gin.Context) {
	var listMsg models.ListMsg
	if err := c.ShouldBindJSON(&listMsg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var postList []models.Post
	offset := int((listMsg.Page - 1) * listMsg.Size)
	if err := models.Db.Limit(int(listMsg.Size)).Offset(offset).First(&postList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GetPostList"})
		return
	}

	var count int64
	if listMsg.Page == 1 {
		if err := models.Db.Count(&count).First(&postList).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "GetPostList"})
			return
		}
	}

	c.JSON(http.StatusOK, models.PostListRes{
		List:  postList,
		Total: count,
	})
}

func SetPostInfo(c *gin.Context) {
	var setMsg models.PostSetMsg
	if err := c.ShouldBindJSON(&setMsg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	if err := models.Db.Where("id = ?", setMsg.PostID).First(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
		return
	}
	claims := c.MustGet("claims").(models.CustomClaims)

	if post.UserID != claims.Uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this post"})
		return
	}

	post.Title = setMsg.Title
	post.Content = setMsg.Content
	if err := models.Db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func DeletePost(c *gin.Context) {
	var post models.PostInfoMsg
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var postInfo models.Post
	if err := models.Db.Where("id = ?", postInfo.ID).First(&postInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
		return
	}
	claims := c.MustGet("claims").(models.CustomClaims)

	if postInfo.UserID != claims.Uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this post"})
		return
	}

	if err := models.Db.Delete(&postInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func AddPostComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claims := c.MustGet("claims").(models.CustomClaims)

	comment.UserID = claims.Uid

	if err := models.Db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"comment": comment})
}

func GetPostComments(c *gin.Context) {
	var cList models.CommentListMsg
	if err := c.ShouldBindJSON(&cList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var comments []models.Comment
	offset := int((cList.Page - 1) * cList.Size)
	if err := models.Db.Where("post_id = ?", cList.PostID).Limit(int(cList.Size)).Offset(offset).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}

	var count int64
	if cList.Page == 1 {
		if err := models.Db.Model(models.Comment{}).Where("post_id = ?", cList.PostID).Count(&count).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments count"})
			return
		}
	}

	c.JSON(http.StatusOK, models.CommentListRes{
		List:  comments,
		Total: count,
	})
}
