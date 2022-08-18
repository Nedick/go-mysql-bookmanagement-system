#  Bookstore management system build with Golang and Mysql

```
FOLDERS       FOLDERS         FILES

[CMD]---------------------->[main.go]

        |-->[config]------->[app.go]
        |
        |-->[controllers]-->[book-controller.go]
        |
[PKG]---|-->[models]------->[book.go]
        |
        |-->[routes]------->[bookstore-routes.go]
        |
        |-->[utils]-------->[utils.go]

                ROUTES
controller funcs
[GETBOOKS]<--------[/book/]<----------[GET]
[CREATE BOOK]<-----[/book/]<----------[POST]
[GET BOOK BY ID]<--[/book/{bookid}]<--[GET]
[UPDATE BOOK]<-----[/book/{bookid}]<--[PUT]
[DELETE BOOK]<-----[/book/{bookid}]<--[DELETE]
```

# Requirements
- [go](https://go.dev/doc/tutorial/getting-started)
- [Mysql](https://dev.mysql.com/doc/mysql-getting-started/en/)
- [Gorm](https://gorm.io/index.html)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Postman](https://www.postman.com/)

# Quickstart
```
git clone https://github.com/Nedick/go-mysql-bookmanagement-system.git
cd go-mysql-bookmanagement-system/cmd/main
go build
go run main.go
```

# Usage
> Make sure your Mysql server is running

Add credentials to .env file in the following format:
```
DB_CREDENTIALS="user:password@/yourdb?charset=utf8&parseTime=True&loc=Local"
```

Open browser and navigate to:
```
localhost:7171/book/
```
Using [Postman](https://www.postman.com/) you can play with the endpoints.