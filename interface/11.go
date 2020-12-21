package main

type device struct {
}

type FirstTester interface {
	test1(i int) int
	test2(s string) string
}

// 实现io.Writer的Write()方法
func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *device) test1(i int) int {
	return 111
}

func (d *device) test2(s string) string {
	return "111"
}

func main() {

}
