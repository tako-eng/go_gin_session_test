package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./view/*.html") //テンプレートの設定

	// セッションの設定
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session-id", store))

	// ログインページを表示
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login", gin.H{
			"UserName":     "",
			"ErrorMessage": "",
		})
	})

	// ログイン試行
	router.POST("/login", func(c *gin.Context) {
		UserName := c.PostForm("user-name")
		session := sessions.Default(c) //セッションにデータを格納する
		session.Set("UserName", UserName)
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/menu/top")
	})

	menu := router.Group("/menu")
	menu.Use(sessionCheck())
	{
		menu.GET("/top", func(c *gin.Context) {
			UserName, _ := c.Get("UserName") // ログインユーザの取得
			fmt.Println(UserName)
			c.HTML(http.StatusOK, "menu", gin.H{"UserName": UserName}) //html側に「ユーザー名：xxx」で表示するため
		})
	}

	//セッションからデータを破棄する
	router.POST("/logout", func(c *gin.Context) {
		session := sessions.Default(c) //セッションにデータを格納する
		session.Clear()
		session.Save()
		// ログインフォームに戻す
		c.HTML(http.StatusOK, "login", gin.H{
			"UserName":     "",
			"ErrorMessage": "",
		})
	})

	router.Run(":9001")

}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		type SessionInfo struct {
			UserName interface{}
		}
		var LoginInfo SessionInfo

		session := sessions.Default(c)
		LoginInfo.UserName = session.Get("UserName")

		// セッションがない場合、ログインフォームをだす
		if LoginInfo.UserName == nil {
			log.Println("ログインしていません")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort() // これがないと続けて処理されてしまう
		} else {
			c.Set("UserName", LoginInfo.UserName) // ユーザidをセット
			c.Next()                              // c.Nextの前に記述した処理は、関数が呼ばれる前に動作する。後に記述した処理は、関数実行後に動作する。
		}
		log.Println("ログインチェック終わり")
	}
}
