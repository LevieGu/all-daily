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
	// 通过 ajax 请求检测用户名是否可以使用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	http.HandleFunc("/getBooks", controller.GetBooks)
	http.HandleFunc("/addBook", controller.AddBook)
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	http.HandleFunc("/updateBook", controller.UpdateBookByID)
	// 分页获取
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)

	//测试
	http.HandleFunc("/setTest", controller.SetCookieDome)
	http.HandleFunc("/gettTest", controller.GetCookieDome)

	http.ListenAndServe(":8080", nil)
}
