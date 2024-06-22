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
После установки и запуска доступ по адресу <b>[localhost:3000](http://localhost:3000)</b>  

Реализованы следующие RESTful эндпоинты:
- [localhost:3000](localhost:3000) Сведения о текущей версии
- Post [localhost:3000/books](localhost:3000/books) Добавить книгу
- Get [localhost:3000/books](localhost:3000/books) Получить список всех книг
- Get [localhost:3000/books/{id}](localhost:3000/books/{id}) Получить книгу по ее идентификатору
- Put [localhost:3000/books/{id}](localhost:3000/books/{id}) Обновить книгу по ее идентификатору
- Delete [localhost:3000/books/{id}](localhost:3000/books/{id}) Удалить книгу
- Post [localhost:3000/authors](localhost:3000/authors) Добавить автора
- Get [localhost:3000/authors](localhost:3000/authors) Получить список всех авторов
- Get [localhost:3000/authors/{id}](localhost:3000/authors/{id}) Получить автора по его идентификатору
- Put [localhost:3000/authors/{id}](localhost:3000/authors/{id}) Обновить автора по его идентификатору
- Delete [localhost:3000/authors/{id}](localhost:3000/authors/{id}) Удалить автора
- Put [localhost:3000/books/{book_id}/authors/{author_id}](localhost:3000/books/{book_id}/authors/{author_id}) Обновить сведения о книге и авторе

Подробное описание методов доступно по адресу [localhost:3000/swagger/index.html](localhost:3000/swagger/index.html)
после запуска сервиса