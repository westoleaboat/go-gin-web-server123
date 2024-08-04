package main

import (
	// "fmt"
	// "html"
	// "io"
	"net/http"
	"sync"

	// "strings"
	// "time"

	"github.com/westoleaboat/go-gin-web-server123/cli"
	// "cli"

	"github.com/gin-gonic/gin"
)

// func rateLimit(c *gin.Context) {
// 	ip := c.ClientIP()
// 	value := int(ips.Add(ip, 1))
// 	if value%50 == 0 {
// 		fmt.Printf("ip: %s, count: %d\n", ip, value)
// 	}
// 	if value >= 200 {
// 		if value%200 == 0 {
// 			fmt.Println("ip blocked")
// 		}
// 		c.Abort()
// 		c.String(http.StatusServiceUnavailable, "you were automatically banned :)")
// 	}
// }

// Register routes
func registerRoutes(router *gin.Engine) {
	router.GET("/", index)
	// router.POST("/submit", submit)
	router.POST("/translate", translate)
	// router.GET("/room/:roomid", roomGET)
	// router.POST("/room-post/:roomid", roomPOST)
	// router.GET("/stream/:roomid", streamRoom)
}

func index(c *gin.Context) {
	// c.Redirect(http.StatusMovedPermanently, "/room/hn")
	// c.Redirect(http.StatusMovedPermanently, "/")
	c.HTML(http.StatusOK, "index.templ.html", gin.H{"title":"hello","message":"hello world",
	})
}

// func submit(c *gin.Context) {
// 	var json struct {
// 		UserInput string `json:"userInput"`
// 	}

// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Log the user input to the console
// 	c.JSON(http.StatusOK, gin.H{"status": "success", "userInput": json.UserInput})

// 	// Use the user input for further logic here
// 	// For example, save it to a database or process it further
// }

func translate(c *gin.Context) {
	var json struct {
		UserInput  string `json:"userInput"`
		SourceLang string `json:"sourceLang"`
		TargetLang string `json:"targetLang"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	strChan := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)

	reqBody := &cli.RequestBody{
		SourceLang: json.SourceLang,
		TargetLang: json.TargetLang,
		SourceText: json.UserInput,
	}

	go cli.RequestTranslate(reqBody, strChan, &wg)

	translatedStr := <-strChan
	close(strChan)

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"translatedText": translatedStr})
}

// func roomGET(c *gin.Context) {
// 	roomid := c.Param("roomid")
// 	nick := c.Query("nick")
// 	if len(nick) < 2 {
// 		nick = ""
// 	}
// 	if len(nick) > 13 {
// 		nick = nick[0:12] + "..."
// 	}
// 	c.HTML(http.StatusOK, "room_login.templ.html", gin.H{
// 		"roomid":    roomid,
// 		"nick":      nick,
// 		"timestamp": time.Now().Unix(),
// 	})

// }

// func roomPOST(c *gin.Context) {
// 	roomid := c.Param("roomid")
// 	nick := c.Query("nick")
// 	message := c.PostForm("message")
// 	message = strings.TrimSpace(message)

// 	validMessage := len(message) > 1 && len(message) < 200
// 	validNick := len(nick) > 1 && len(nick) < 14
// 	if !validMessage || !validNick {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": "failed",
// 			"error":  "the message or nickname is too long",
// 		})
// 		return
// 	}

// 	post := gin.H{
// 		"nick":    html.EscapeString(nick),
// 		"message": html.EscapeString(message),
// 	}
// 	messages.Add("inbound", 1)
// 	room(roomid).Submit(post)
// 	c.JSON(http.StatusOK, post)
// }

// func streamRoom(c *gin.Context) {
// 	roomid := c.Param("roomid")
// 	listener := openListener(roomid)
// 	ticker := time.NewTicker(1 * time.Second)
// 	users.Add("connected", 1)
// 	defer func() {
// 		closeListener(roomid, listener)
// 		ticker.Stop()
// 		users.Add("disconnected", 1)
// 	}()

// 	c.Stream(func(w io.Writer) bool {
// 		select {
// 		case msg := <-listener:
// 			messages.Add("outbound", 1)
// 			c.SSEvent("message", msg)
// 		case <-ticker.C:
// 			c.SSEvent("stats", Stats())
// 		}
// 		return true
// 	})
// }
