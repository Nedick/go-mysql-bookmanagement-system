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