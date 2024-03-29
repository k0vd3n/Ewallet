# Минимальный образ golang в качестве базового образа
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

# Использование минимального образа alpine в качестве базового образа для запуска приложения
FROM alpine:latest
WORKDIR /app

# Установка зависимостей для работы приложения
RUN apk --no-cache add ca-certificates

# Копирование скомпилированного бинарного файла из предыдущего образа
COPY --from=build /app/EWallet .

# Порт, который будет прослушиваться приложением
EXPOSE 8080

# Команда для запуска приложения
CMD ["./EWallet"]
