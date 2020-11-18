package controllers

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"morningo/connections/database"
	db "morningo/connections/database/mysql"
	"morningo/filters/auth"
	m "morningo/models"
	"net/http"
	"time"
	"fmt"
	"strconv"
	"sync"
	mlog "morningo/modules/log"
)

func IndexApi(c *gin.Context) {

	// 返回html
	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"title": "GO GO GO GO GO!",
	})
}

func DBExample(c *gin.Context) {

	// 数据库插入
	insertRs, _ := db.Exec("insert into users (first_name, last_name, age) values (?, ?, ?)", "人才", "unknown", 1)
	insertId, _ := insertRs.LastInsertId()
	log.Printf("insert id: %d\n", insertId)

	// 数据库更新
	db.Exec("update users set first_name = ? where id = ?", "饭桶", insertId)

	// 数据库中间件
	_, _ = database.Table("users").Where("id", "=", insertId).Update(database.H{
		"first_name": "你好",
	})

	// 数据库查询
	rs := db.Query("select first_name,last_name,id from users where id < ?", 100)
	log.Println(rs[0]["first_name"])

	rs1, _ := database.Table("users").
		Select("first_name", "last_name", "id").
		Where("id", "<", 100).
		All()
	log.Println(rs1[0])

	// 数据库事务
	_, _ = db.WithTransaction(func(tx *db.SqlTxStruct) (error, map[string]interface{}) {
		_, err := tx.Query("select first_name,last_name,id from users where id < ?", 100)
		if err != nil {
			return err, map[string]interface{}{}
		}
		return nil, map[string]interface{}{}
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"query_result": rs,
		},
	})
}

func StoreExample(c *gin.Context) {

	// session存储
	session := sessions.Default(c)
	session.Set("key", "value") // 0表示不过期

	str := session.Get("key")
	log.Printf("session key: %s", str)

	// cache存储
	cacheStore, _ := c.MustGet(cache.CACHE_MIDDLEWARE_KEY).(*persistence.CacheStore)
	_ = (*cacheStore).Set("key", "value", time.Minute)
	_ = (*cacheStore).Delete("key")

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"store_result": str,
		},
	})
}

func OrmExample(c *gin.Context) {

	// Create
	m.Model.Create(&m.User{FirstName: "L1212", LastName: "unknown"})

	// Read
	var user m.User
	m.Model.First(&user, 1) // find user with id 1
	log.Printf("user model insert %d\n", user.Model.ID)
	m.Model.First(&user, "first_name = ?", "L1212") // find user with name l1212

	// Update
	m.Model.Model(&user).Update("last_name", "123456")

	// Delete
	m.Model.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"orm_result": user,
		},
	})
}

func CreateUsers(c *gin.Context) {
	startTime := time.Now()

	for i := 0; i < 4000; i++ {
		var age int = i + 1
		email := fmt.Sprintf("testmail%s@test.com", strconv.Itoa(i))
	    user := m.User{	FirstName: "First Name 01", LastName: "Last Name 01", Email: email, Age: age}
	    m.Model.Create(&user)
	}

	elapsed := time.Since(startTime)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  fmt.Sprintf("ElapsedTime in seconds: %f", elapsed.Seconds()),
	})
}


func CreateUserGoroutinesStandard(c *gin.Context) {
	startTime := time.Now()

	var total int = 1000
	var chunk int = 100

	mlog.Info(mlog.E{Info: mlog.M{"data": "started",},})

	i := 0
	for i < total {
		// run concurrent chunks
		go func(s int) {
			mlog.Info(mlog.E{Info: mlog.M{"chunk is:": s,},})
			for i := 0; i < chunk; i++ {
				h := s + (i + 1);
				email := fmt.Sprintf("testmail%s@test.com", strconv.Itoa(h))
    			user := m.User{	FirstName: "First Name 01", LastName: "Last Name 01", Email: email, Age: h}
    			m.Model.Create(&user)			
			}			
		}(i)

		i = i + chunk
	}

	elapsed := time.Since(startTime)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  fmt.Sprintf("ElapsedTime in seconds: %f", elapsed.Seconds()),
	})
}

func CreateUserGoroutines(c *gin.Context) {
	startTime := time.Now()

	var total int = 20000
	// var chunk int = 500

	mlog.Info(mlog.E{Info: mlog.M{"data": "started",},})

	var wg sync.WaitGroup
	// allow max 7 concurrent workers/threads
	wg.Add(7)
	for n := 0; n < 7; n++ {
        mlog.Info(mlog.E{Info: mlog.M{"added worker:": n,},})

        // each worker processing by 3000 records
        total_per_worker := 3000
        offset := total_per_worker * n
        current_total := total_per_worker * (n + 1 ) 
        if current_total > total {
        	total_per_worker = total_per_worker - (current_total - total)
        }

        go runWorker(total_per_worker, &wg, offset)	
    }

    // no need to wait for threads to finish
    // wg.Wait() 

	elapsed := time.Since(startTime)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  fmt.Sprintf("ElapsedTime in seconds: %f", elapsed.Seconds()),
	})
}

func runWorker(total int, wg *sync.WaitGroup, s int) {
	defer wg.Done()

	mlog.Info(mlog.E{Info: mlog.M{"chunk is:": s,},})
	for i := 0; i < total; i++ {
		h := s + (i + 1);
		email := fmt.Sprintf("testmail%s@test.com", strconv.Itoa(h))
		user := m.User{	FirstName: "First Name 01", LastName: "Last Name 01", Email: email, Age: h}
		m.Model.Create(&user)			
	}
	mlog.Info(mlog.E{Info: mlog.M{"chunk finished is:": s,},})
}

func CreateUser(c *gin.Context) {
	var name string = "First Name 01"
	user := m.User{FirstName: name, LastName: "Last Name 01", Email: "testmail@test01.com", Age: 33}
	m.Model.Create(&user)

	// var user m.User
	// m.Model.First(&user, "first_name = ?", name)
	// userId := strconv.Itoa(int(user.Model.ID))

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": user,
	})
}

func GetUser(c *gin.Context) {
	var user m.User
	userId := c.Param("userId")

	if err := m.Model.Where("id = ?", userId).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user m.User
	userId := c.Param("userid") 
	
	m.Model.First(&user, "id = ?", userId)

	m.Model.Model(&user).Update("last_name", "Last Name Updated")

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"orm_result": "user updated",
		},
	})
}

func CookieSetExample(c *gin.Context) {
	authDr, _ := c.MustGet("web_auth").(auth.Auth)

	id := c.Param("userid")

	rs := db.Query("select first_name,last_name,id from users where id = ?", id)

	log.Printf("len(rs): %d", len(rs))
	if len(rs) == 0 {
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"title": "wrong user id",
		})
		return
	}

	authDr.Login(c.Request, c.Writer, map[string]interface{}{"id": id})

	// 返回html
	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"title": "login success!",
	})
}

func CookieGetExample(c *gin.Context) {
	authDr, _ := c.MustGet("web_auth").(auth.Auth)

	userInfo := authDr.User(c).(map[interface{}]interface{})
	id, _ := userInfo["id"].(string)
	log.Println("id: " + id)

	rs := db.Query("select first_name,last_name,id from users where id = ?", id)

	// 返回html
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"user": rs,
		},
	})
}

func JwtSetExample(c *gin.Context) {
	authDr, _ := c.MustGet("jwt_auth").(auth.Auth)

	token, _ := authDr.Login(c.Request, c.Writer, map[string]interface{}{"id": "123"}).(string)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"token": token,
		},
	})
}

func JwtGetExample(c *gin.Context) {
	authDr, _ := c.MustGet("jwt_auth").(auth.Auth)

	info := authDr.User(c).(map[string]interface{})

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"id": info["id"],
		},
	})
}
