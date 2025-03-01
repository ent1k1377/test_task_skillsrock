# Task Management API

## Запуск проекта

### 1. Подготовка конфигурации
Скопируйте файл конфигурации и переименуйте его:
```sh
mv example.env .env
```

### 2. Запуск Docker-контейнеров
Используйте команду для сборки и запуска контейнеров с переменными окружения:
```sh
docker compose up --build
```

### 3. Миграции базы данных

#### 3.1 Установите `golang-migrate`
Установите инструмент для миграции базы данных, следуя инструкциям из официального репозитория:
[golang-migrate](https://github.com/golang-migrate/migrate)

#### 3.2 Примените миграции
Запустите команду для применения миграций:
```sh
make migrate_up
```

После успешного выполнения этих шагов API будет готово к использованию.


---
## Скриншоты

![Screenshot 1](imgs/photo_2025-03-01_17-04-33.jpg)
![Screenshot 2](imgs/photo_2025-03-01_17-05-51.jpg)
![Screenshot 3](imgs/photo_2025-03-01_17-06-07.jpg)
---
![Screenshot 4](imgs/photo_2025-03-01_17-06-30.jpg)
---
![Screenshot 5](imgs/photo_2025-03-01_17-07-16.jpg)
![Screenshot 6](imgs/photo_2025-03-01_17-07-27.jpg)
---
![Screenshot 7](imgs/photo_2025-03-01_17-07-57.jpg)
![Screenshot 8](imgs/photo_2025-03-01_17-08-12.jpg)