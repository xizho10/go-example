package storage

type Storage interface {
	WriteLog(title string, note string, link string)
	ReadLog(id string) interface{}
	ChangeLog(id string, title string) interface{}
}
