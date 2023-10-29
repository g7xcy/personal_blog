package main

import (
	"log"
	"net/http"
	"strconv"

	"example.com/auth"
	"example.com/blog"
	"example.com/model"
	"example.com/util"
	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	DBErr := model.HealthCheck()
	if DBErr != nil {
		c.JSON(http.StatusOK, DBErr.Error())
	}
	c.JSON(http.StatusOK, "OK")
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}

func registerUser(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}
	userID, err := auth.Register(&newUser)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, userID)
	}
}

func login(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}
	user, err := auth.Login(&user)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func getBlogs(c *gin.Context) {
	var blogs []model.Blog
	pageSize, _ := strconv.Atoi(c.Query("size"))
	pageNumber, _ := strconv.Atoi(c.Query("page"))
	blogs, err := blog.BlogsByPage(pageSize, pageNumber)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, blogs)
	}
}

func getBlogsByID(c *gin.Context) {
	var blg model.Blog
	blog_id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	blg, err := blog.BlogsByID(blog_id)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, blg)
	}
}

func getBlogsByUser(c *gin.Context) {
	var blogs []model.Blog
	user := c.Param("user")
	blogs, err := blog.BlogsByUser(user)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, blogs)
	}
}

func postBlog(c *gin.Context) {
	var newBlog model.Blog
	if err := c.BindJSON(&newBlog); err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}
	blogID, err := blog.AddBlog(&newBlog)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, blogID)
	}
}

func main() {
	if err := util.LoadEnv(); err != nil {
		log.Fatal(err)
		return
	}
	if err := model.InitDB(); err != nil {
		log.Fatal(err)
		return
	}
	router := gin.Default()
	router.GET("/", home)
	router.GET("/health", health)

	router.POST("/register", registerUser)
	router.GET("/login", login)

	router.GET("/blogs/", getBlogs)
	router.GET("/blog/id/:id", getBlogsByID)
	router.GET("/blog/user/:user", getBlogsByUser)
	router.POST("/blog", postBlog)

	router.Run(util.Env["URL"] + ":" + util.Env["PORT"])
	defer model.DB.Close()
}
