# Цитатник - REST API сервис

Простой REST API сервис на Go для управления цитатами.

## Функциональность

- ✅ Добавление новой цитаты (POST /quotes)
- ✅ Получение всех цитат (GET /quotes)
- ✅ Получение случайной цитаты (GET /quotes/random)
- ✅ Фильтрация по автору (GET /quotes?author=Confucius)
- ✅ Удаление цитаты по ID (DELETE /quotes/{id})

## Требования

- Go 1.24 или выше
- Git

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone https://github.com/bezlant/quotes
cd quotes-service
```

2. Установите зависимости:
```bash
go mod tidy
```

3. Запустите сервер:
```bash
make serve
```

Сервер запустится на порту 8080.

## API Endpoints

### 1. Добавление новой цитаты
```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

**Ответ:**
```json
{
  "id": 1,
  "author": "Confucius",
  "quote": "Life is simple, but we insist on making it complicated."
}
```

### 2. Получение всех цитат
```bash
curl http://localhost:8080/quotes
```

**Ответ:**
```json
[
  {
    "id": 1,
    "author": "Confucius",
    "quote": "Life is simple, but we insist on making it complicated."
  },
]
```

### 3. Получение случайной цитаты
```bash
curl http://localhost:8080/quotes/random
```

**Ответ:**
```json
  {
    "id": 1,
    "author": "Confucius",
    "quote": "Life is simple, but we insist on making it complicated."
  }
```

### 4. Фильтрация по автору
```bash
curl http://localhost:8080/quotes?author=Confucius
```

**Ответ:**
```json
[
  {
    "id": 1,
    "author": "Confucius",
    "quote": "Life is simple, but we insist on making it complicated."
  }
]
```

### 5. Удаление цитаты по ID
```bash
curl -X DELETE http://localhost:8080/quotes/1
```

**Ответ:** HTTP 204 No Content

## Запуск тестов

Для запуска unit-тестов выполните:

```bash
make test
```

## Структура проекта

```
quotes/
├──  cmd/
│   └──  quotes/
│       └──  main.go
├──  internal/
│   ├──  app/
│   │   └──  quotes/
│   │       ├──  handler.go
│   │       ├──  handler_test.go
│   │       ├──  model.go
│   │       ├──  service.go
│   │       └──  service_test.go
│   └──  pkg/
│       └──  middleware/
│           └──  json.go
├──  Makefile
└── 󰂺 README.md
├──  go.mod
├──  go.sum
```

## Архитектура

Проект использует следующую архитектуру:

- **Service** - бизнес-логика для управления цитатами
- **Handler** - HTTP обработчики для REST API
- **Quote** - структура данных для цитаты
- Данные хранятся в памяти

## Технические особенности

- Использует только стандартные библиотеки Go + gorilla/mux для роутинга
- Поиск по автору
- Валидация входных данных
- Unit-тесты с покрытием всех сценариев

## Статус коды HTTP

- **200 OK** - Успешное получение данных
- **201 Created** - Цитата успешно создана
- **204 No Content** - Цитата успешно удалена
- **400 Bad Request** - Неверный формат данных
- **404 Not Found** - Цитата не найдена

## Что можно улучшить?
- Добавить логгирование
- Блокировка чтения-записи

