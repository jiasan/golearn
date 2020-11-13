package main

import "fmt"

type filer interface {
	Read() string
}

type File struct {
	Name string
}

func (f *File) Read() string {
	return f.Name
}

func Get(file interface{}) string {
	f := file.(filer)
	return f.Read()
}

type Service interface {
	Start()     // 开启服务
	Log(string) // 日志输出
}

// 日志器
type Logger struct {
}

// 实现Service的Log()方法
func (g *Logger) Log(l string) {
	fmt.Println(l)
}

// 游戏服务
type GameService struct {
	Logger // 嵌入日志器
}

// 实现Service的Start()方法
func (g *GameService) Start() {
	fmt.Println("GameService.start")
}

func main() {
	var s Service = new(GameService)
	s.Start()
	s.Log("hello")

	f := &File{Name: "hello"}
	var w filer
	w = filer(f)
	fmt.Println(w.Read())

	fmt.Println(Get(f))
}
