<div align="center">

# TodosApp

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![React](https://img.shields.io/badge/React-61DAFB?style=flat&logo=react&logoColor=black)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=flat&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-4FC08D?style=flat&logo=vite&logoColor=white)

</div>

---

## Описание

**TodosApp** — это приложение для управления задачами, использующее бэкенд на Go с фреймворком Chi и PostgreSQL для хранения данных. Фронтенд построен с использованием React и Vite. 

## Стек технологий

- **Backend:** Go, Chi, PostgreSQL
- **Frontend:** React, Vite
- **Deployment:** Docker, Docker Compose

---

## Установка и запуск

### Предварительные требования

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Шаги

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/k5sha/golang-todo-example.git
   cd golang-todo-example
   ```

2. Создайте файл `development.yaml` в корне проекта со следующим содержимым:

   ```yaml
   env: "local"
   http_server:
    address: "0.0.0.0:8080"
    timeout: 4s
    idle_timeout: 30s
   ```

3. Запустите проект:

   ```bash
   docker-compose up --build
   ```

4. Откройте браузер и перейдите по адресу [http://localhost:8080](http://localhost:8080).

---

## Использование API

**Эндпоинты:**

- `GET /todo` — Получить все задачи
- `GET /todo/{id}` — Получить задачу по ID
- `POST /todo` — Создать новую задачу
- `DELETE /todo/{id}/delete` — Удалить задачу по ID
- `PATCH /todo/{id}/status` — Обновить статус задачи

---


## Вклад

Если вы хотите внести свой вклад, создайте Pull Request или откройте Issue с предложениями.

---

## Лицензия

Этот проект лицензирован под MIT License. См. [LICENSE](LICENSE) для подробностей.
