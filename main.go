package main

import (
	"GinDemo04/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "我是首页")
	//})
	//
	//r.GET("/users", func(c *gin.Context) {
	//	c.HTML(200, "userAdd.html", gin.H{})
	//})
	//
	//r.POST("/doUpload", func(c *gin.Context) {
	//	username := c.PostForm("username")
	//	file, err := c.FormFile("face")
	//
	//	dst := path.Join("./static/upload", file.Filename)
	//	if err == nil {
	//		c.SaveUploadedFile(file, dst)
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"success":  true,
	//		"username": username,
	//		"dst":      dst,
	//	})
	//})
	//
	//r.GET("/edit", func(c *gin.Context) {
	//	c.HTML(200, "userEdit.html", gin.H{})
	//})
	//
	//r.POST("/doEdit", func(c *gin.Context) {
	//	username := c.PostForm("username")
	//	file1, err1 := c.FormFile("face1")
	//	file2, err2 := c.FormFile("face2")
	//
	//	dst1 := path.Join("./static/upload", file1.Filename)
	//	if err1 == nil {
	//		c.SaveUploadedFile(file1, dst1)
	//	}
	//	dst2 := path.Join("./static/upload", file2.Filename)
	//	if err2 == nil {
	//		c.SaveUploadedFile(file2, dst2)
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"success":  true,
	//		"username": username,
	//		"dst1":     dst1,
	//		"dst2":     dst2,
	//	})
	//})

	//查询
	r.GET("/user", func(c *gin.Context) {

		userList := []models.User{}

		models.DB.Find(&userList)

		c.HTML(http.StatusOK, "userFind.html", userList)

		//c.JSON(http.StatusOK, gin.H{
		//	"result": userList,
		//})

		c.String(http.StatusOK, "用户列表")
	})

	//增加
	r.GET("/add", func(c *gin.Context) {

		//user := models.User{
		//	Username: "CC",
		//	Age:      22,
		//	Email:    "222@qq.com",
		//	AddTime:  12345333,
		//}
		c.HTML(http.StatusOK, "userAdd.html", gin.H{})
	})
	r.POST("/doAdd", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		username := c.PostForm("username")
		age, _ := strconv.Atoi(c.PostForm("age"))
		email := c.PostForm("email")
		addTime, _ := strconv.Atoi(c.PostForm("addTime"))

		user := models.User{
			Id:       id,
			Username: username,
			Age:      age,
			Email:    email,
			AddTime:  addTime,
		}

		models.DB.Create(&user)

		c.String(http.StatusOK, "添加成功")

	})

	//修改
	r.GET("/ed", func(c *gin.Context) {
		c.HTML(http.StatusOK, "userEdit.html", gin.H{})
	})
	r.POST("/doFind", func(c *gin.Context) {

		id, _ := strconv.Atoi(c.PostForm("id"))

		user := models.User{Id: id}

		models.DB.Find(&user)

		c.HTML(http.StatusOK, "userEd.html", user)
	})
	r.POST("/doEd", func(c *gin.Context) {

		id, _ := strconv.Atoi(c.PostForm("id"))
		fmt.Println(id)
		username := c.PostForm("username")
		age, _ := strconv.Atoi(c.PostForm("age"))
		email := c.PostForm("email")
		addTime, _ := strconv.Atoi(c.PostForm("addTime"))

		user := models.User{Id: id}
		models.DB.Find(&user)

		user.Username = username
		user.Age = age
		user.Email = email
		user.AddTime = addTime

		models.DB.Save(&user)

		//user := models.User{}
		//models.DB.Model(&user).Where("id = ?", 5).Update("username", "李四")

		//user := models.User{}
		//models.DB.Where("id=?", 7).Find(&user)
		//user.Username = "张三"
		//user.Email = "niuniuniu@qq.com"
		//models.DB.Save(&user)

		c.String(http.StatusOK, "修改成功！")
	})

	//删除
	r.GET("/del", func(c *gin.Context) {
		c.HTML(http.StatusOK, "userDelete.html", gin.H{})
	})
	r.POST("/doDel", func(c *gin.Context) {

		id, _ := strconv.Atoi(c.PostForm("id"))

		user := models.User{Id: id}
		models.DB.Delete(&user)

		//user := models.User{}
		//models.DB.Where("id=?", 7).Delete(&user)

		c.String(http.StatusOK, "删除成功")
	})

	r.Run(":8080")
}
