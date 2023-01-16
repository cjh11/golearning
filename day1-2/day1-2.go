package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	//局部声明变量方式1
	var a int
	fmt.Println("a=", a)
	//声明变量方式2
	var b int = 100
	fmt.Println("b=", b)
	//声明变量方式3
	var c = 200
	fmt.Println("c=", c)
	fmt.Printf("c=%v,type of c=%T\n", c, c)
	//声明变量方式4(常用)
	e := 300
	fmt.Printf("e=%v,type of e=%T\n", e, e)
	f := float64(e)
	fmt.Printf("f=%f,type of f=%T\n", f, f)

	//常量
	const eh = 2000
	const i = eh / 4
	fmt.Println(eh, math.Sin(i))

	var ll, yy = 123, "sad"
	fmt.Println(ll, yy)
	const (
		u1 = iota * 10 //iota=0
		u2             //iota=1,u2=10
		u3             //iota=2,u3=20
	)
	fmt.Println(u1, u2, u3)

	if num := 9; num < 0 {
		fmt.Println("是负数")
	} else {
		fmt.Println("不是负数哦")
	}

	qq := -5
	for qq <= 0 {
		fmt.Println(qq)
		qq++
	}

	//for {
	//	time.Sleep(5 * time.Second)
	//}
	for i := 0; i < 5; i++ {
		fmt.Println("我爱学go")
	}

	//switch
	grade := "A"
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}

	//数组
	//var variable_name [SIZE] variable_type
	var balance [10]float32 //数组长度不确定可以使用...
	balance[0] = 1

	//初始化数组
	var balance1 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//或者
	balance2 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(len(balance1))
	//可以指定具体位置的元素
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	balance1[0] = 4
	balance2[0] = 4
	balance3[0] = 4
	var twoD [2][3]int
	twoD[0][0] = 1

	//slice
	slice1 := make([]int, 3)
	slice1[0] = 0
	slice1[1] = 1
	fmt.Println(len(slice1))
	slice1 = append(slice1, 12)
	fmt.Println(slice1)

	qw := make([]int, 3)
	copy(qw, slice1[:3]) //此处不包括3位置的元素
	fmt.Println(qw)

	//map
	var map1 = make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	fmt.Println(map1)
	fmt.Println(len(map1))

	r, ok := map1["three"]
	fmt.Println(r, ok)
	delete(map1, "one")
	fmt.Println(map1)
	//range
	nums := []int{1, 2}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Printf("sum is %d,index is %d", sum, i)
	}

	map2 := make(map[string]int)
	map2["ss"] = 8
	for k, v := range map2 {
		fmt.Println(k, v)
	}

	//函数
	map3 := make(map[string]int)
	map3["sa"] = 1
	v, ok := exists(map3, "sa")
	fmt.Println(v, ok)
	n := 1
	add1(&n)
	fmt.Println(n)

	//结构体

	struct1 := user{"ccc", "1234"}
	struct2 := user{}
	struct2.name = "sdksd"
	fmt.Println(struct1)
	fmt.Println(checkPassword1(&struct1, "1234"))

	//结构体方法

	res := struct2.checkPassword("1234")
	fmt.Println(res)

	//错误处理
	user1 := []user{{"cc", "123"}}
	u, err := findUser(user1, "cc")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)

	//字符串处理
	str := "hello1"
	str1 := "hello2"
	strr := []string{str, str1}
	fmt.Println(strings.Contains(str, "ll"))
	fmt.Println(strings.Count(str1, "l"))
	fmt.Println(strings.Index(str, "l"))
	fmt.Println(strings.Join(strr, "-"))
	fmt.Println(strings.Replace(str, "he", "hh", -1))
	fmt.Println(strings.ToUpper(str))

	//格式化输出
	n1 := 123
	p := point{1, 2}
	fmt.Println(n1, p)

	fmt.Printf("s=%v\n", str)
	fmt.Printf("s=%v\n", p)

	//json
	aa := userInfo{"wang", 23, []string{"love", "ll"}}
	buf, err := json.Marshal(aa)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)
	fmt.Println(string(buf))
	//buf, err = json.MarshalIndent(aa, "", "\t")
	var bb userInfo
	err = json.Unmarshal(buf, &bb)
	if err != nil {
		panic(err)
	}
	fmt.Printf("序列化结果%v", bb)

	//时间处理
	now := time.Now()
	fmt.Println(now) //2023-01-16 20:02:02.7679312 +0800 CST m=+0.00582790
	//使用time.date构造时间
	t := time.Date(2023, 1, 16, 1, 25, 60, 0, time.UTC)
	t1 := time.Date(2024, 1, 16, 1, 25, 60, 0, time.UTC)
	fmt.Println(t)
	diff := t1.Sub(t)
	fmt.Println(diff)

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	//获取时间戳
	fmt.Println(now.Unix())

	//数字解析
	f, _ = strconv.ParseFloat("1.23456", 64)
	fmt.Println(f)
	n12, _ := strconv.ParseInt("11", 10, 64)
	fmt.Println(n12)
	nn, _ := strconv.Atoi("123456")
	fmt.Println(nn)
}

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}
type point struct {
	x, y int
}

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

type user struct {
	name     string
	password string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}
func checkPassword1(u *user, password string) bool {
	return u.password == password
}
func add(a int) {
	a = a + 2
}

func add1(a *int) {
	*a = *a + 2
}
func exists(m map[string]int, k string) (v int, ok bool) {
	v, ok = m[k]
	return v, ok
}
