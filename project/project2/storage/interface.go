package storage

type Storage interface {
	WriteLog(title string, note string, link string)
	ReadLog(id string) string
	ChangeLog(id string, title string) string
}
