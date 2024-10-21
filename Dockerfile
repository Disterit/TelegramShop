# Используем официальный образ Golang для сборки приложения
FROM golang:1.22

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /api

# Копируем файлы go.mod и go.sum в рабочую директорию
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь код в рабочую директорию
COPY .. .

# Собираем приложение
RUN go build -o main ./api/cmd/main.go
RUN go build -o consumer ./kafka/consumer.go

# Указываем команду запуска собранного приложения
CMD ["./main"]

# Указываем, какой порт должен быть открыт в контейнере
EXPOSE 8000
