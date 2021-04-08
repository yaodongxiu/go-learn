// RBAC（attribute-based-access-control）
package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

type Object struct {
	Name  string
	Owner string
}

type Subject struct {
	Name string
	Hour int
}

func check(e *casbin.Enforcer, sub Subject, obj Object, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s at %d:00\n", sub.Name, act, obj.Name, sub.Hour)
	} else {
		fmt.Printf("%s CANNOT %s %s at %d:00\n", sub.Name, act, obj.Name, sub.Hour)
	}
}

// RBAC模型对于实现比较规则的、相对静态的权限管理非常有用。但是对于特殊的、动态的需求，RBAC就显得有点力不从心了。
// 例如，我们在不同的时间段对数据data实现不同的权限控制。
// 正常工作时间9:00-18:00所有人都可以读写data，其他时间只有数据所有者能读写
func main() {
	e, err := casbin.NewEnforcer("./model.conf")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	o := Object{"data", "dajun"}
	s1 := Subject{"dajun", 10}
	check(e, s1, o, "read")

	s2 := Subject{"lizi", 10}
	check(e, s2, o, "read")

	s3 := Subject{"dajun", 20}
	check(e, s3, o, "read")

	s4 := Subject{"lizi", 20}
	check(e, s4, o, "read")
}
