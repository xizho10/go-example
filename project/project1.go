package project

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func writeLog(title string, note string, link string) {
	filePath := "./project/project1.csv"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	defer file.Close()

	str := strconv.Itoa(rand.Intn(100)) + "," + title + "," + time.Now().Format("2006-01-02 15:03:04") + "," + note + "," + link
	file.WriteString(str + "\r\n")
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	values := req.URL.Query()
	title := values.Get("title")
	note := values.Get("note")
	link := values.Get("link")
	fmt.Println(title + "," + note + "," + link)
	writeLog(title, note, link)
	io.WriteString(
		res,
		`<doctype html>
<html>
    <head>
          <title>Success</title>
    </head>
<script language="javascript">
setTimeout(function(){
window.history.back(-1);
},1000);

</script>
    <body>
          Submit Success
    </body>
</html>`,
	)
}

//<h1 style="background-color:DodgerBlue;">Hello World</h1>
//<p style="background-color:Tomato;">Lorem ipsum...</p>

func ztfunc(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`
<!DOCTYPE html>
<html>
<body>

<h2>Project 1</h2>

<form action="/hello">
  <label for="fname">Title:</label>
  <input type="text" id="id_title" name="title" value="chapter13-server2"><br>
  <label for="lname">Note:</label>
  <input type="text" id="id_note" name="note" value="learn http server"><br>
  <label for="lname">Link:</label>
  <input type="text" id="id_link" name="link" value="https://github.com/tongzheng2048/go-example"><br><br>   
  <input type="submit" value="Submit">
</form> 

<p>If you click the "Submit" button, the form-data will be sent to a page called "/hello".</p>

</body>
</html>

`,
	)
}
func Project1() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", ztfunc)
	fmt.Println("start server: http://127.0.0.1:9000")
	http.ListenAndServe(":9000", nil)
}
