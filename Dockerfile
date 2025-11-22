# 1. Билдим с последней версией Go
FROM golang:1.24-alpine AS builder

# 2. Создаём рабочую директорию
WORKDIR /app

# 3. Копируем файлы зависимостей и устанавливаем их
COPY go.mod go.sum ./
RUN go mod download

# 4. Копируем весь исходный код проекта
COPY . .

# 5. Собираем бинарник из cmd/api
RUN go build -o main ./cmd/api

# 6. Минимальный образ для запуска
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .

# 7. Команда запуска
CMD ["./main"]
