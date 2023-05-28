## Наш адрес для тестов: 37.18.110.184:3000

## Connect
указатель к бд лежит в internal/models.go -> Tools

## Как запустить проект:
Для запуска из Docker:
Прописываем:
```
docker-compose up -d --build hack
```
или
```
go mod download
go run cmd/main.go
```
## Для запуска нужно:
1) Скачать все либы через go mod download
2) Изменить параметы в файл config на свои

### Пример
```
Postgres:
  user: "almaz"
  password: "almaz"
  host: "localhost"
  port: "5432"
  dbName: "postgres"
```
