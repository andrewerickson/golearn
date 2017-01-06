package main

import (
	"flag"
	"fmt"
	// nice stuff, https://github.com/gin-gonic/gin
	"github.com/andrewerickson/golearn/routes"
	"gopkg.in/gin-gonic/gin.v1"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	// "time"
)

func add(x int, y int) int {
	return x + y
}

var listenAddr = flag.String("addr", "127.0.0.1:1718", "http service address") // Q=17, R=18

func myHandler(w http.ResponseWriter, req *http.Request) {
	log.Output(1, fmt.Sprintf("processing request with form value %v", req.FormValue("s")))
	templ.Execute(w, req.FormValue("s"))
}

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("my fav number is", rand.Intn(10))
	fmt.Println("lets add some stuff", add(rand.Intn(12), rand.Intn(11)))

	flag.Parse()
	/* server := &http.Server{
		Addr:           *listenAddr,
		Handler:        http.HandlerFunc(myHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
	*/
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/index", routes.GetIndexPage)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}

const templateStr = `
<html>
<head>
<title>Come at me bro</title>
</head>
<body>
<br>
Gin n' joiuce!
<br>
</body>
</html>
`
