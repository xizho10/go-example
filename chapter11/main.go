// need to export the GOPATH
// export GOPATH=/Users/chee/golang/golang-book/chapter11
// math module needs to go into src folder

// to generate godoc:
// godoc src/chapter11/math Average
// godoc -http=":6060"
// http://localhost:6060/pkg

package chapter11

import (
	m "example/chapter11/src/math"
	"fmt"
) //引用

func Chapter11main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
	fmt.Println(m.Min(xs)) //输出xs最小值
	fmt.Println(m.Max(xs)) //输出xs最大值
}
