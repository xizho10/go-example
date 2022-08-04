package project

import (
	"encoding/json"
	"example/project/project2/storage"
	"fmt"
	"io"
	"net/http"
)

type GetRequestData struct {
	Name string
	Id   string
}

type UpdateRequestData struct {
	Title string
	Id    string
}

type Response struct {
	Code   int
	Msg    string
	Result interface{}
}
type Router struct {
	storage storage.Storage
} //结构体

//<h1 style="background-color:DodgerBlue;">Hello World</h1>
//<p style="background-color:Tomato;">Lorem ipsum...</p>

func (r *Router) Index(res http.ResponseWriter, req *http.Request) {
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

console.log("==hello===");

$(document).ready(function(){
 $("#getBtn").click(function(){
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
           url: "/getlog",
           async: false,
           data: JSON.stringify({ Name: "Donald Duck", Id: eleValue }),
           contentType: "application/json",
           complete: function (data) {
				console.log(data.responseJSON);
				alert(JSON.stringify(data.responseJSON));
           
           }
    });
  });

 $("#updateBtn").click(function(){

    var eleUpdateId = document.getElementById("eleUpdateId").value;
    var eleUpdateTitle = document.getElementById("eleUpdateTitle").value;
    $.ajax({
           type: "POST",
           url: "/updatelog",
           async: false,
           data: JSON.stringify({ Title: eleUpdateTitle, Id: eleUpdateId }),
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



<h3>1. Add Log</h3>

<form action="/addlog">
  <label for="fname">Title:</label>
  <input type="text" id="id_title" name="title" value="chapter13-server2"><br>
  <label for="lname">Note:</label>
  <input type="text" id="id_note" name="note" value="learn http server"><br>
  <label for="lname">Link:</label>
  <input type="text" id="id_link" name="link" value="https://github.com/tongzheng2048/go-example"><br><br>   
  <input type="submit" value="Add Log">
</form> 

<p>If you click the "Submit" button, the form-data will be sent to a page called "/hello".</p>

<h3>2. Get Log</h3>

<label >Log Id:</label>
<input id="element" className="element" name="element">
<button id="getBtn">Get Log</button>


<h3>3. Update Log</h3>

<label >Log Id:</label>
<input id="eleUpdateId" className="eleUpdateId" name="eleUpdateId" value="47">
<br>
<label >Log Title:</label>
<input id="eleUpdateTitle" className="eleUpdateTitle" name="eleUpdateTitle" value="hello">
<br>
<button id="updateBtn">Update Log</button>

</body>
</html>

`,
	)
}

func (r *Router) AddLog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	values := req.URL.Query()
	title := values.Get("title")
	note := values.Get("note")
	link := values.Get("link")
	fmt.Println(title + "," + note + "," + link)
	r.storage.WriteLog(title, note, link)
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

func (r *Router) GetLog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"application/json",
	)
	fmt.Println("===", req.Header.Values("Content-Type"))
	// Declare a new Person struct.
	var rp GetRequestData

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
		Result: r.storage.ReadLog(rp.Id),
	}
	data, _ := json.Marshal(resData)

	io.WriteString(res, string(data))

}

func (r *Router) UpdateLog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"application/json",
	)
	fmt.Println("===", req.Header.Values("Content-Type"))
	// Declare a new Person struct.
	var rp UpdateRequestData

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
		Result: r.storage.ChangeLog(rp.Id, rp.Title),
	}
	data, _ := json.Marshal(resData)

	io.WriteString(res, string(data))

}
