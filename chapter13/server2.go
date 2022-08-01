package chapter13

import (
	"fmt"
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`<doctype html>
<html>
     <head>
           <title>Hello World</title>
     </head>
     <body>
           Hello World!
     </body>
</html>`,
	)
}

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

<h2>HTML Forms</h2>
<h1 style="background-color:DodgerBlue;">Hello World</h1>
<p style="background-color:Tomato;">Lorem ipsum...</p>
<form action="/action_page.php">
  <label for="fname">First name:</label><br>
  <input type="text" id="fname" name="fname" value="John"><br>
  <label for="lname">Last name:</label><br>
  <input type="text" id="lname" name="lname" value="Doe"><br><br>
  <input type="submit" value="Submit">
</form> 

<p>If you click the "Submit" button, the form-data will be sent to a page called "/action_page.php".</p>

</body>
</html>

`,
	)
}
func Chapter13main9() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", ztfunc)
	fmt.Println("start server: http://127.0.0.1:9000")
	http.ListenAndServe(":9000", nil)
}
