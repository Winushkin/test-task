# TSV Processing Service (Golang)

## Overview

TSV Processing Service — это backend-сервис на Go, который:

- Мониторит директорию на наличие новых `.tsv` файлов
- Ставит обработку файлов в очередь
- Парсит данные устройств
- Сохраняет данные в PostgreSQL
- Логирует ошибки парсинга в БД и файл
- Генерирует выходные PDF файлы по `unit_guid`
- Предоставляет REST API с пагинацией

---
## 🛠 Используемый стек

- Go 1.21+
- PostgreSQL
- net/http
- Worker Pool (goroutines + channels)
- swaggo/swag (Swagger)
- Docker + docker-compose

---

## 📂 Структура проекта

````
├── cmd
├── docs                    # Swagger docs
├── internal
│       ├── config
│       ├── entities
│       ├── handlers
│       ├── logger
│       ├── parser
│       ├── poller
│       ├── postgres
│       ├── report
│       │       └── fonts
│       ├── repository
│       │       └── sql
│       └── worker
├── logs
├── migrator
│       └── migrations
├── reports                 # Сгенерированные PDF
└── tsv_files               # Входные .tsv файлы
````
---

## ⚙ Конфигурация

Настройки берутся из переменных окружения.

### Переменные окружения

| Переменная               | Описание                                     |
|--------------------------|----------------------------------------------|
| POSTGRES_HOST            | Хост БД                                      |
| POSTGRES_PORT            | Порт БД                                      |
| POSTGRES_USER            | Пользователь                                 |
| POSTGRES_PASSWORD        | Пароль                                       |
| POSTGRES_NAME            | Имя БД                                       |
| POSTGRES_MIN_CONNS       | Минимальное кол-во активных подключений к БД |
| POSTGRES_MAX_CONNS       | Максимальное кол-во подключений к БД         |
| TSV_DIR_PATH             | Директория для .tsv файлов                   |
| REPORT_DIR_PATH          | Директория выходных PDF отчетов              |
| POLLING_SECONDS_INTERVAL | Интервал сканирования (сек)                  |
| LOG_DIR_PATH             | Директория для сохранения файлов с логами    |

Пример `.env`:
````
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=PGUSER
POSTGRES_PASSWORD=1234
POSTGRES_NAME=DB
POSTGRES_MIN_CONNS=1
POSTGRES_MAX_CONNS=5

TSV_DIR_PATH=./tsv_files
REPORT_DIR_PATH=./reports
POLLING_SECONDS_INTERVAL=5
LOG_DIR_PATH=./logs
````
---

## 🚀 Запуск проекта
Требования:
* Установленный make
* golang 1.25+
* установленный docker + docker compose

### 1️⃣ Через Docker (рекомендуется)
```bash
  git clone https://github.com/Winushkin/test-task.git
  make start
  make up
```
Сервис будет доступен по:
http://localhost:8080/records

Swagger:
http://localhost:8080/swagger

### !!!. Чтобы получить отчеты и файлы с логами для просмотра
```bash
  make backup-from-container
```
После чего файлы будут доступны для просмотра в папке backup

### 2️⃣ Локальный запуск
```bash
  git clone https://github.com/Winushkin/test-task.git
  make start
  make pg-up
  go mod tidy
  go run cmd/main.go
```

🌐 REST API

Получение всех данных
GET /records?page=1&limit=10

Параметры:
* page: Страница для пагинации
* limit: Кол-во результатов на 1 странице

Пример запроса
```bash
  curl "http://localhost:8080/records?page=1&limit=1"
```

Пример ответа
```json
{
 "records": [
  {
   "Number": "1",
   "Mqtt": "",
   "InvID": "G-044322",
   "UnitGUID": "01749246-95f6-57db-b7c3-2ae0e8be671f",
   "MessageID": "cold7_Defrost_status",
   "MessageText": "Разморозка",
   "Context": "",
   "MessageClass": "waiting",
   "MessageLevel": "100",
   "Area": "LOCAL",
   "VarAddress": "cold7_status.Defrost_status",
   "Block": "",
   "MessageType": "",
   "BitNumber": "",
   "InvertBit": "",
   "FileID": 2
  }
 ],
 "total_records": 126,
 "page": 1,
 "limit": 1, 
 "total_pages": 126
}
```
🛡 Обработка ошибок

Ошибки парсинга:
* 	Записываются в БД 
* 	Логируются в stdout
*   Сохраняются в .log файлах
