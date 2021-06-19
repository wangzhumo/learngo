package impls

import (
	"fmt"
	"strings"
)

// Queue  interface{}  表示任何类型
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// 实现了接口 stringer ，类似 toString
func (q *Queue) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	for _, value := range *q {
		sb.WriteString(fmt.Sprintf("%v", value))
		sb.WriteString(",")
	}
	sb.WriteString(")")
	return fmt.Sprintf("Queue %s", sb.String())
}
