# Генератор паролей

## Как запустить?

### Через Go 

```bash
go run main.go
```

### Docker 

Создать image 

```bash
docker build -t password-gen .
```

Запустить контейнер
```bash
docker run -it --rm -v $(pwd)/data:/app password-generator
```
Файл passwords.dat будет сохраняться в папку data

## Софт
Go 1.23.2 и выше
Docker 

