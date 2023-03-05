# image-previewer

Проект [Превьюер изображений](https://github.com/OtusGolang/final_project/blob/master/03-image-previewer.md)

По-умолчанию сервер запускается по адресу `127.0.0.0:8080`

Пример запроса ```http://127.0.0.1:8080/fill/800/600/raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/_gopher_original_1024x504.jpg```

Команды:

```bash
# сборка bin
make build

# Локальный запуск
make run

# Запуск unit-тестов
make test

# Запуск интеграционных тестов
make integration-test

# Запуск линтера
make lint
```
