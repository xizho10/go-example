package project

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func readLog(id string) string {
	filePath := "./project/project1.csv"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}
	defer file.Close()

	br := bufio.NewReader(file)
	fmt.Println(id)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
		strArr := strings.Split(string(a), ",")
		if strArr[0] == id {
			return string(a)
		}
	}
	return ""
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

type RequestData struct {
	Name string
	Id   string
}

type Response struct {
	Code   int
	Msg    string
	Result interface{}
}

func postData(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"application/json",
	)
	fmt.Println("===", req.Header.Values("Content-Type"))
	// Declare a new Person struct.
	var rp RequestData

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(req.Body).Decode(&rp)
	if err != nil {
		fmt.Println(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Person struct...
	//fmt.Fprintf(res, "%+
	fmt.Println(rp)

	resData := Response{
		Code:   0,
		Msg:    "success",
		Result: readLog(rp.Id),
	}
	data, _ := json.Marshal(resData)

	io.WriteString(res, string(data))

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


<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.6.0/jquery.min.js"></script>
<script>
$(document).ready(function(){
  $("button").click(function(){
    //$.post("/postdata",
    //{
    //  name: "Donald Duck",
    //  age: "Duckburg"
    //},
    //function(data,status){
    //  alert("Data: " + data + "\nStatus: " + status);
    //},
    //'application/json; charset=utf-8'
    // );

     var eleValue = document.getElementById("element").value;
     $.ajax({
            type: "POST",
            url: "/postdata",
            async: false,
            data: JSON.stringify({ Name: "Donald Duck", Id: eleValue }),
            contentType: "application/json",
            complete: function (data) {
				console.log(data.responseJSON);
				alert(JSON.stringify(data.responseJSON));
            
        }
     });

  });
});
</script>


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

<label >Log Id:</label>
<input id="element" className="element" name="element" />
<button>Get Log</button>

</body>
</html>

`,
	)
}
func Project1() {
	http.HandleFunc("/postdata", postData)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", ztfunc)
	fmt.Println("start server: http://127.0.0.1:9000")
	http.ListenAndServe(":9000", nil)
}
