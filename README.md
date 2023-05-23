## Формат ошибок, которые возвращаем на фронт:
``` 
type HackError struct {
	Code      int         // статус код ошибки
	Err       error       // системное сообщение об ошиюке (error)
	Message   string      // Описание ошибки/Решение ее проблемы (может быть пустым)
	Timestamp time.Time   // время, когда произошла ошибка
}
```

### Не забываем, про логирование.
Лушче всего делать лог в месте ошибки и потом ее прокидывать на вверх (в хендлер), уже без вызова log.Print(...)

## ФАЙЛ config.yml СО СВОИМИ ДАННЫМИ НЕ КОМИТИМ И НЕ ПУШИМ!

## Наш адрес для тестов: localhost:3000

## Connect
указатель к бд лежит в internal/models.go -> Tools

## Логгер
Нужен кто-то, кто разберется в нем и всем объяснит, как им пользоваться (хд)

## Как запустить проект:
https://learn.microsoft.com/ru-ru/azure/developer/go/configure-visual-studio-code 
Проходим первые 4 пункта
Если терминал выдает ошибку на комамнду go, то  Терминал > Новый терминал
Прописываем:
```
go mod download
go run cmd/main.go
```
## Для запускаа нужно:
1) скачать Golang 1.20
2) Скачать все либы через go mod download
3) Иметь активную бд Postgres
4) Изменить параметы в файл config на свои
### Пример
```
Postgres:
  user: "almaz"
  password: "almaz"
  host: "localhost"
  port: "5431"
  dbName: "postgres"
```
