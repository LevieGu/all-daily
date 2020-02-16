package main

import (
	"html/template"
	"net/http"
	"shopping/controller"
)

// indexHandler 首页处理器
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// 解析模版
	t := template.Must(template.ParseFiles("views/index.html"))
	// 执行
	t.Execute(w, "")
}

func main() {
	// 设置处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/", indexHandler)
	// 登陆
	http.HandleFunc("/login", controller.Login)
	// 注册
	http.HandleFunc("/regist", controller.Regist)
	http.ListenAndServe(":8080", nil)
}
