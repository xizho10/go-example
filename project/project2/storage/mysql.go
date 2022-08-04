package storage

type Mysql struct {
} //结构体

func (m *Mysql) WriteLog(title string, note string, link string) {

}
func (m *Mysql) ReadLog(id string) string {

	return ""
}
func (m *Mysql) ChangeLog(id string, title string) string {

	return ""
}
