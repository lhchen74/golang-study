### GOPATH

项目的工作路径，可以通过 go env 查看，可以设置多个 GOPATH, GOPATH 下的文件夹

- src 存放项目的源代码，通过 go get 获取的依赖包也会放在这里
- bin 编译过程中生成的可执行文件
- pkg 编译后的包文件，包文件名和所在的目录名一致

### go doc

`godoc -http=:9000` web doc
`go doc strings` 查看 strings 文档
`go doc strings.Split` 查看 strings.Split 文档
`go doc builtin.append` 查看内置函数文档

### beego

1. beego new
2. beego run
3. bee generate scaffold user -fields="id:int64,name:string,gender:int,age:int" -driver=mysql -conn="root:root@tcp(127.0.0.1:3306)/imooc" //create model and controller
4. 注册路由
// beego.Router("/", &controllers.MainController{})
// 使用注解
beego.Include(&controllers.UserController{})
5. 使用orm
orm.RegisterDataBase("default", "mysql", "root:root@/imooc?charset=utf8")
