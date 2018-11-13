package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	route := gin.New()
	/*gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}*/
	route.Use(Logger())
	//route.Use(gin.Logger())
	//route.Use(gin.Recovery())
	route.GET("/testing", startPage)
	route.Run(":8085")
	//

}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
	c.JSON(200,&person)
}

//定义
//记录请求历史记录，可以入库或es,或.log文件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 在gin上下文中定义变量
		c.Set("example", "12345")

		// 请求前

		c.Next()//处理请求

		// 请求后
		latency := time.Since(t)
		log.Print("耗时:",latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println("statuscode:",status)
	}
}