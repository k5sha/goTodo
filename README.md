<div align="center">

# TodosApp

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![React](https://img.shields.io/badge/React-61DAFB?style=flat&logo=react&logoColor=black)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=flat&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-4FC08D?style=flat&logo=vite&logoColor=white)

</div>

---

## Description

**TodosApp** is a task management application with a backend built using Go with the Chi framework and PostgreSQL for data storage. The frontend is developed using React and Vite.

## Technology Stack

- **Backend:** Go, Chi, PostgreSQL
- **Frontend:** React, Vite
- **Deployment:** Docker, Docker Compose

---

## Installation and Running

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/k5sha/golang-todo-example.git
   cd golang-todo-example
   ```

2. Create a `development.yaml` file in the project root with the following contents:

   ```yaml
   env: “local”
   http_server:
    address: “0.0.0.0.0:8080”
    timeout: 4s
    idle_timeout: 30s
   ```

3. Run the project:

   ```bash
   docker-compose up --build
   ```

4. open a browser and go to [http://localhost:8080](http://localhost:8080).

---

## API Usage

**Endpoints:**

- ``GET /todo`` - Get all tasks
- `GET /todo/{id}` - Get task by ID
- `POST /todo` - Create a new task
- `DELETE /todo/{id}/delete` - Delete task by ID
- `PATCH /todo/{id}/status` - Update Task Status

---


## Contribute

If you want to contribute, create a Pull Request or open a Suggestion Issue.

---

## License

This project is licensed under the MIT License. See. [LICENSE](LICENSE) for details.
