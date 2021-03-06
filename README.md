# Тестовое задание для backend-стажёра в команду Advertising

## Задача
Необходимо создать сервис для хранения и подачи объявлений. Объявления должны храниться в базе данных. Сервис должен предоставлять API, работающее поверх HTTP в формате JSON.


## Установка и Запуск
Для запуска нужно собрать и запустить контейнер

`docker-compose build`

`docker-compose up`

## API
* `POST /advertisements` - Создание нового объявления
* `GET /advertisements` - Получение списка всех объявлений
* `GET /advertisements/:id` - Получение объявления по id

## Дополнительная информация
Сервис реализован с использованием "Чистой Архитектуры". 
Ворк-флоу обработки процесса разделён на 3 слоя: Транспортный слой - *Delivery*,
Бизнес логику - *Usecase* и хранилище - *Repository*. 

Для валидации полей при создании объявления реализована функция промежуточной обработки **(middleware)**.

Конфигурация сервиса задается через файл .env.

Технологический стек: Golang + Gin, Postgres.