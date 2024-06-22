![Supported Go Versions](https://img.shields.io/badge/Go-1.21%2C%201.22-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/v/release/medant81/url_library.svg)](https://github.com/medant81/url_library/releases)


# REST API на Go для управления коллекцией книг и авторами, которые их пишут

Этот проект выполнен на Go для демонстрации работы REST API методов. Для хранения данных используется PostgreSQL.
Сервис запускается с помощью Docker.

## Используемые библиотеки:
 - github.com/jackc/pgx/v4 - для подключения к базе на PostgreSQL
 - github.com/go-chi/chi/v5 - маршрутизатор для создания HTTP-сервисов


## Установка и запуск проекта
Для установки и запуска используйте команды Makefile:

 - Создание и запуск сервисов в Docker
```
make run
```

- Создание сервисов и пересоздание сервисов
```
make build
```

- Запуск созданных сервисов
```
make up
```

- Остановка и удаление контейнеров
```
make down
```

- Просмотр созданных контейнеров
```
make ps
```

- Просмотр логов из контейнеров
```
make logs
```

- Остановка сервисов
```
make stop
```

# Как использовать
После установки и запуска доступ по адресу <b><a href="localhost:3000">localhost:3000</a></b>  

Реализованы следующие RESTful эндпоинты:
- <a href="localhost:3000">localhost:3000</a> Сведения о текущей версии
- Post <a href="localhost:3000/books">localhost:3000/books</a> Добавить книгу
- Get <a href="localhost:3000/books">localhost:3000/books</a> Получить список всех книг
- Get <a href="localhost:3000/books/{id}">localhost:3000/books/{id}</a> Получить книгу по ее идентификатору
- Put <a href="localhost:3000/books/{id}">localhost:3000/books/{id}</a> Обновить книгу по ее идентификатору
- Delete <a href="localhost:3000/books/{id}">localhost:3000/books/{id}</a> Удалить книгу
- Post <a href="localhost:3000/authors">localhost:3000/authors</a> Добавить автора
- Get <a href="localhost:3000/authors">localhost:3000/authors</a> Получить список всех авторов
- Get <a href="localhost:3000/authors/{id}">localhost:3000/authors/{id}</a> Получить автора по его идентификатору
- Put <a href="localhost:3000/authors/{id}">localhost:3000/authors/{id}</a> Обновить автора по его идентификатору
- Delete <a href="localhost:3000/authors/{id}">localhost:3000/authors/{id}</a> Удалить автора
- Put <a href="localhost:3000/books/{book_id}/authors/{author_id}">localhost:3000/books/{book_id}/authors/{author_id}</a> Обновить сведения о книге и авторе

Подробное описание методов доступно по адресу <a href="localhost:3000/swagger/index.html">localhost:3000/swagger/index.html</a>
после запуска сервиса