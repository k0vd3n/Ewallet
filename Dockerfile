# Используем минимальный образ golang в качестве базового образа
FROM golang:1.21.6-alpine AS build

# Установка рабочей директории
WORKDIR /app

# Копирование go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копирование исходного кода в контейнер
COPY . .

# Сборка приложения
RUN go build -o EWallet .

# Используем минимальный образ alpine в качестве базового образа для запуска приложения
FROM alpine:latest
WORKDIR /app

# Установка зависимостей для работы приложения
RUN apk --no-cache add ca-certificates

# Копирование скомпилированного бинарного файла из предыдущего образа
COPY --from=build /app/EWallet .

# Задаем переменные окружения для подключения к базе данных
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=root
ENV DB_NAME=EWallet

# Порт, который будет прослушиваться вашим приложением
EXPOSE 8080

# Команда для запуска приложения
CMD ["./EWallet"]
