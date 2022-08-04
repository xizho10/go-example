package storage

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const filePath = "./project/project2/project2.csv"

type LocalFile struct {
} //结构体

func (l *LocalFile) WriteLog(title string, note string, link string) {

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	defer file.Close()

	str := strconv.Itoa(rand.Intn(100)) + "," + title + "," + time.Now().Format("2006-01-02 15:03:04") + "," + note + "," + link
	file.WriteString(str + "\r\n")
}
func (l *LocalFile) ReadLog(id string) interface{} {

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
func (l *LocalFile) ChangeLog(id string, title string) interface{} {

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}

	defer file.Close()

	fmt.Println(id, title)
	br := bufio.NewReader(file)
	var total int = 0
	for {

		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if total == 0 {
			total = br.Buffered()
		}
		fmt.Println(br.Size(), br.Buffered())
		fmt.Println(string(a))
		strArr := strings.Split(string(a), ",")
		if strArr[0] == id {
			strArr[1] = title
			newStr := strings.Join(strArr, ",")
			file.WriteAt([]byte(newStr), int64(total-br.Buffered()))
			return newStr
		}
	}
	return ""
}
