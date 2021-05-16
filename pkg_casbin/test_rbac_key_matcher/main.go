// 我们可以在匹配器中使用函数。
// casbin内置了一些函数keyMatch/keyMatch2/keyMatch3/keyMatch4都是匹配 URL 路径的，regexMatch使用正则匹配，ipMatch匹配 IP 地址。
// 参见https://casbin.org/docs/en/function。
// 使用内置函数我们能很容易对路由进行权限划分
package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	check(e, "dajun", "user/dajun/1", "read")
	check(e, "lizi", "user/lizi/2", "read")
	check(e, "dajun", "user/lizi/1", "read")
}
