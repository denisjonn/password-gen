# Генератор паролей

## Как запустить?

### Через Go 

```bash
go run main.go
```

### Docker 

Windows Powershell:

Создать image 

```Powershell
docker build -t password-gen .
```

Запустить контейнер
```Powershell
docker run -it --rm -v ${PWD}\data:/app/data password-gen
```
Файл passwords.dat будет сохраняться в папку data

## Софт
Go 1.23.2 и выше, 
Docker 

