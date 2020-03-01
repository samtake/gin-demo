# gin-demo
 gin框架学习

 

学习资料:
[gin项目地址](https://github.com/gin-gonic/gin)
[gin中文文档](https://gin-gonic.com/zh-cn/)
[Gin入门实战](https://www.imooc.com/learn/1175)
[文中demo](https://github.com/samtake/gin-demo)


## Gin运行流程

![](gin.png)

`engine`:实现了ServeHTTP接口的handler
`methodTree`:根据http请求方法分别维护的路由树
`routerGroup`:将路由表分组，方便中间件统一处理
`Context`:Gin的上下文，在handler之间传递参数


## 示例代码

Router ：路由规则定义
```bash
import (
	hdl "filestore-server/handler"

	"github.com/gin-gonic/gin"
)

// Router ：路由规则定义
func Router() *gin.Engine {
	// gin framework
	router := gin.Default()

	// 静态资源处理
	router.Static("/static/", "./static")

	// 定义接口
	router.GET("/user/signup", hdl.SignupHandler)
    router.POST("/user/signup", hdl.DoSignupHandler)
    
    return router
}

```


handler

```bash
// SignupHandler : 处理用户注册请求
func SignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "http://"+c.Request.Host+"/static/view/signup.html")
}
```

```bash
// DoSignupHandler : 处理用户注册请求
func DoSignupHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	if len(username) < 3 || len(passwd) < 5 {
		c.JSON(http.StatusOK,
			gin.H{
				"msg": "Invalid parameter",
			})
		return
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	// 将用户信息注册到用户表中
	suc := dblayer.UserSignup(username, encPasswd)
	if suc {
		c.JSON(http.StatusOK,
			gin.H{
				"code":    0,
				"msg":     "注册成功",
				"data":    nil,
				"forward": "/user/signin",
			})
	} else {
		c.JSON(http.StatusOK,
			gin.H{
				"code": 0,
				"msg":  "注册失败",
				"data": nil,
			})
	}
}
  
```

最终main.go可简洁为

```bash
func main() {
	// gin framework
	router := route.Router()

	// 启动服务并监听端口
	err := router.Run(config.UploadServiceHost)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s\n", err.Error())
	}
}

```



## 快速开始
```bash
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("first", func(c *gin.Context) {
		fmt.Println("first .........")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authorized := r.Group("/try")

	authorized.POST("/second", second)
	authorized.POST("/third", third)

	// 嵌套路由组
	testing := authorized.Group("testing")
	testing.GET("/forth", fourth)

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

func second(c *gin.Context) {
	fmt.Println("second .........")
}

func third(c *gin.Context) {
	fmt.Println("third .........")
}

func fourth(c *gin.Context) {
	fmt.Println("fourth .........")
}

```



```bash
➜  gin-demo git:(master) ✗ lsof -i tcp:8080    
COMMAND     PID    USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
QQ          327 samtake   29u  IPv4 0xfdcbb6b79716f31f      0t0  TCP huangloshansmbp:50202->157.255.13.190:http-alt (ESTABLISHED)
QQ          327 samtake   48u  IPv4 0xfdcbb6b79716f31f      0t0  TCP huangloshansmbp:50202->157.255.13.190:http-alt (ESTABLISHED)
Google      420 samtake   26u  IPv4 0xfdcbb6b793e58fb7      0t0  TCP localhost:51492->localhost:http-alt (ESTABLISHED)
Postman   66875 samtake   67u  IPv4 0xfdcbb6b778668ca7      0t0  TCP localhost:51499->localhost:http-alt (ESTABLISHED)
main      67553 samtake    3u  IPv6 0xfdcbb6b793ea9b8f      0t0  TCP *:http-alt (LISTEN)
main      67553 samtake    7u  IPv6 0xfdcbb6b793eac3cf      0t0  TCP localhost:http-alt->localhost:51492 (ESTABLISHED)
main      67553 samtake    8u  IPv6 0xfdcbb6b7912a55cf      0t0  TCP localhost:http-alt->localhost:51499 (ESTABLISHED)
➜  gin-demo git:(master) ✗ kill -9  67553      
➜  gin-demo git:(master) ✗ go run start/main.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /first                    --> main.main.func1 (3 handlers)
[GIN-debug] POST   /try/second               --> main.second (3 handlers)
[GIN-debug] POST   /try/third                --> main.third (3 handlers)
[GIN-debug] GET    /try/testing/forth        --> main.fourth (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
first .........
[GIN] 2020/02/25 - 11:51:19 | 200 |     176.613µs |       127.0.0.1 | GET      /first
```

![](gin-first.png)



## 1.请求路由

### 多种请求
```bash
func main() {
	r := gin.Default()

	r.GET("get", func(c *gin.Context) {
		c.String(200, "get")
	})

	r.POST("post", func(c *gin.Context) {
		c.String(200, "post")
	})

	r.DELETE("delete", func(c *gin.Context) {
		c.String(200, "delete")
	})

	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})

	r.Run()
}
```


给我们创建的any请求几乎覆盖了所有种类的请求
```bash
[GIN-debug] GET    /get                      --> main.main.func1 (3 handlers)
[GIN-debug] POST   /post                     --> main.main.func2 (3 handlers)
[GIN-debug] DELETE /delete                   --> main.main.func3 (3 handlers)
[GIN-debug] GET    /any                      --> main.main.func4 (3 handlers)
[GIN-debug] POST   /any                      --> main.main.func4 (3 handlers)
[GIN-debug] PUT    /any                      --> main.main.func4 (3 handlers)
[GIN-debug] PATCH  /any                      --> main.main.func4 (3 handlers)
[GIN-debug] HEAD   /any                      --> main.main.func4 (3 handlers)
[GIN-debug] OPTIONS /any                      --> main.main.func4 (3 handlers)
[GIN-debug] DELETE /any                      --> main.main.func4 (3 handlers)
[GIN-debug] CONNECT /any                      --> main.main.func4 (3 handlers)
[GIN-debug] TRACE  /any                      --> main.main.func4 (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

### 静态文件夹

```bash
func main() {
	r := gin.Default()
	// 静态文件夹绑定+路由，有两种写法：
	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("static"))

	//又或者路由+资源
	r.StaticFile("/favicon.ico", "/favicon.ico")
}
```
`go build -o router_static && ./router_static`这需要在routerstatic文件夹下运行，不然找不到资源文件，最后访问`http://localhost:8080/assets/a.html`以及`http://localhost:8080/static/b.html`测试即可。



### 参数作为url



Get请求`http://localhost:8080/Sam/520`即可得到相应参数
```bash
{
    "id": "520",
    "name": "Sam"
}
```


## 2.获取请求参数

- 获取get请求参数
- 获取post请求参数
- 获取body值
- 获取参数绑定结构体


### 泛绑定

```bash
func main() {
	r := gin.Default()
	r.GET("/user/*action", func(c *gin.Context) { //设置获取name和id的参数
		c.String(200, "泛绑定～")
	})

	r.Run()
}
```
所有user前缀的请求都能请求到：`http://localhost:8080/user/<XXXX>`

### 获取get参数

```bash
func main() {
	r := gin.Default()
	r.GET("/testGetParam", func(c *gin.Context) { //设置获取name和id的参数
		param1 := c.Query("parm1")
		param2 := c.Query("parm2")

		c.String(http.StatusOK, "%s, %s", param1, param2)
	})

	r.Run()
}
```



```bash
➜  ~ curl -X GET 'http://127.0.0.1:8080/testGetParam?parm1=11111' 
11111, %  
➜  ~ curl -X GET 'http://127.0.0.1:8080/testGetParam?parm1=11111&parm2=2222'
11111, 2222% 
```



### 获取body内容

```bash
func main() {
	r := gin.Default()
	r.POST("/parmBody", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) //重新回写到缓存才能拿到parm1 parm2
		parm1 := c.PostForm("parm1")
		parm2 := c.PostForm("parm2")
		c.String(http.StatusOK, "%s %s %s", parm1, parm2, string(bodyBytes))
	})

	r.Run()
}
```



```bash
➜  ~ curl -X POST 'http://127.0.0.1:8080/parmBody' -d 'parm1=value1&parm2=value2' 
value1 value2 parm1=value1&parm2=value2%                                                                                                                               
➜  ~ curl -X POST 'http://127.0.0.1:8080/parmBody' -d '{"parm1":"value1","parm2":"value2"}'
  {"parm1":"value1","parm2":"value2"}

```



### 获取bind参数

同时响应post和get，同时访问到同一个回调方法。

```bash
type Person struct {
	Name     string    `form:"name"` //设置tag `form`可以由参数转变成结构体
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-03"`
}

func main() {
	r := gin.Default()
	r.GET("/test", testHandler)
	r.POST("/test", testHandler)
	r.Run()
}

func testHandler(c *gin.Context) {
	var person Person
	//根据请求content-type来作不同的binding操作
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(200, "person bind error", err)
	}

}
```

```bash
➜  ~ curl -X GET  'http://127.0.0.1:8080/test?name=samtake&adress=gd'
{samtake  0001-01-01 00:00:00 +0000 UTC}
➜  ~ curl -X GET  'http://127.0.0.1:8080/test?name=samtake&address=gd'
{samtake gd 0001-01-01 00:00:00 +0000 UTC}%  
➜  ~ curl -X POST  'http://127.0.0.1:8080/test?name=samtake&address=gd&birthday=2008-09-09'
{samtake gd 2008-09-01 09:00:00 +0800 CST}%   
➜  ~ curl -X POST  'http://127.0.0.1:8080/test'  -d 'name=samtake&address=gd&birthday=2008-09-09'
{samtake gd 2008-09-01 09:00:00 +0800 CST}%   
➜  ~ curl -H "Content-Type:application/json"  -X POST "http://127.0.0.1:8080/test"  -d '{"name":"wang"}'
{wang  0001-01-01 00:00:00 +0000 UTC}%      
```


这里有个坑：如果时间格式写成像`time_format:"2006-01-03"`这种会报错，解析不了。
```bash
person bind error%!(EXTRA *time.ParseError=parsing time "2008-00-08" as "2011-01-03": cannot parse "-00-08" as "1")
```


## 验证请求参数
- 结构体验证
- 自定义验证
- 支持多语言错误信息


### 结构体验证
[validate规则](https://godoc.org/gopkg.in/go-playground/validator.v9)
binding条件满足



坑的记录：
```bash
type Person struct {
	Age     int    `form:"age"  binding:"required,gt=10"`
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}
```
![](validate-form.png)



正确的完整源码：
```bash
type Person struct {
	Age     int    `form:"Age"  binding:"required,gt=10"`
	Name    string `form:"Name" binding:"required"`
	Address string `form:"Address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/test", testHandler)
	r.Run()
}

func testHandler(c *gin.Context) {
	var person Person
	//根据请求content-type来作不同的binding操作
	if err := c.ShouldBind(&person); err != nil {
		c.String(500, "%v", err)
		c.Abort()
		return
	}
	c.String(200, "%v", person)

}

```


```bash
➜  ~ curl -X GET  "http://127.0.0.1:8080/test?age=19&name=samtake&address=gd"
Key: 'Person.Age' Error:Field validation for 'Age' failed on the 'required' tag
Key: 'Person.Name' Error:Field validation for 'Name' failed on the 'required' tag
Key: 'Person.Address' Error:Field validation for 'Address' failed on the 'required' tag% 
➜  ~ curl -X GET  "http://127.0.0.1:8080/test?Age=19&Name=samtake&Address=gd"
{19 samtake gd}%        
```

### 自定义验证

```
package main

import (
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in"  binding:"required,bookabledate" time_format:"2006-01-01"`
	CheckOut time.Time `form:"check_out"  binding:"required,gtfield=checkIn" time_format:"2006-01-01"`
}

func customFunc(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if date.Unix() > today.Unix() {
			return true
		}
	}
	return false
}

func main() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", customFunc)
	}
	r.GET("/bookable", testHandler)
	r.Run()
}

func testHandler(c *gin.Context) {
	var b Booking
	if err := c.ShouldBind(&b); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "ok!", "booking": b})
}

```

这里有个报错，我自己也还没弄明白，然而我看了下源码，上面的使用方法是正确的：

```bash
# command-line-arguments
validCustom/validCustom.go:30:24: cannot use customFunc (type func(*validator.Validate, reflect.Value, reflect.Value, reflect.Value, reflect.Typ
e, reflect.Kind, string) bool) as type validator.Func in argument to v.RegisterValidation
````

```bash
// RegisterValidation adds a validation with the given tag
//
// NOTES:
// - if the key already exists, the previous validation function will be replaced.
// - this method is not thread-safe it is intended that these all be registered prior to any validation
func (v *Validate) RegisterValidation(tag string, fn Func) error {
	return v.RegisterValidationCtx(tag, wrapFunc(fn))
}

```

### 支持多语言错误信息

## 中间件
- 使用Gin中间件
- 自定义`ip白名单`中间件


### Gin中间件
`Logger`日志
`Recovery`捕获panic

```bash
func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	r.Use(gin.Logger(),gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		c.String(200, "%s", name)
	})

	r.Run()
}
```

###  白名单
`r.User(IPAuthMiddleware())`

```bash
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIP := c.ClientIP()
		for _, host := range ipList {
			if clientIP == host {
				flag = true
				break
			}
		}

		if !flag {
			c.String(401, "%s,not in iplist", clientIP)
			c.Abort()
		}
	}
}

func main() {
	r := gin.New()
	r.Use(IPAuthMiddleware())
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "%s", "hello test")
	})

	r.Run()
}

```


## Gin延展

- 服务器优雅关停
- 模版渲染
- 自动证书配置

### 服务器关停
![](gin-server-stop.png)


### 模版渲染

```bash
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "index.html",
		})
	})
	r.Run()
}
```



```bash
➜  ~ curl -X GET  "http://127.0.0.1:8080/index"
<html>
    <h1>
        index.html
    </h1>
</html>%                                                                        
➜  ~ 
```

### 自动证书配置



## 脚手架
项目地址：
[gin_scaffold](https://github.com/e421083458/gin_scaffold)
[golang_common](https://github.com/e421083458/golang_common)
[vue-admin](https://github.com/e421083458/vue-admin)

### 轻量级Golang类库

GORM:[https://gorm.io/zh_CN/](https://gorm.io/zh_CN/)
redigo:[https://godoc.org/github.com/gomodule/redigo/redis](https://godoc.org/github.com/gomodule/redigo/redis)


├── README.md
├── conf   配置文件夹
│   └── dev
│       ├── base.toml
│       ├── mysql_map.toml
│       └── redis_map.toml
├── controller 控制器
│   └── demo.go
├── dao DB数据访问层
│   └── demo.go
├── dto  Bind结构体层
│   └── demo.go
├── gin_scaffold.inf.log  info日志
├── gin_scaffold.wf.log warning日志
├── go.mod go module管理文件
├── go.sum
├── main.go
├── middleware 中间件层
│   ├── panic.go
│   ├── response.go
│   ├── token_auth.go
│   └── translation.go
├── public  公共文件
│   ├── log.go
│   ├── mysql.go
│   └── validate.go
├── router  路由层
│   ├── httpserver.go
│   └── route.go
└── tmpl

### 输出格式统一封装

### 自定义中间件日志打印

### 请求数据绑定结构体与校验

