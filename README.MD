main.go
- callGet(s string) (name string, value int)
    функция для отправки get запроса
    возвращает url на который был отправлен запрос
    возвращает количество вхождений по заданному параметру

cобираем:
go build main.go

запускаем:
main.exe -input=https://golang.org\nhttps://golang.org\nhttps://golang.org
