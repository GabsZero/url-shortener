# Url Shortener

![GitHub repo size](https://img.shields.io/github/repo-size/gabszero/url-shortener?style=for-the-badge)
![GitHub language count](https://img.shields.io/sourceforge/languages/url-shortener?logo=go&style=for-the-badge)


> A small url shortener project, written in Go, that act as a solid foundation to apply scalability concepts.

### Features and planned features

This project is currently under development, but I do not expect to launch it in production in the future. This is somekind of a toy project to learn concepts, and thats all. With that said, here are the main functionalities available:

- [x] Endpoint to shorten an url with a random string
- [x] Endpoint to redirect to the destiny url
- [x] Endpoint to shorten an url with a custom string defined by the user
- [x] Load balancer with default round robin algorithm to distribute requests
- [x] Cache results of redirect for better performance using redis
- [x] Simulating shards with application level logic to distribute data
- [ ] Refactor to look more like a Go project
- [ ] Tests
- [ ] More to come? Who knows

## 💻 Requirements

This project was designed to run in parallel to a database container, you could add database services inside the project and run everything together, tho. It was just a personal choice.

Here is a suggestion to start a mysql database container with shards in mind:
```
# ...docker compose omitted here
mysql1:
    container_name: mysql-1
    image: mysql:5.7
    environment:
      - MYSQL_ROOT_PASSWORD=mysqlpw
    ports:
      - 3306:3306
    volumes:
      - ./data:/var/lib/mysql1
    networks: 
      - yourNetwork

mysql2:
    container_name: mysql-2
    image: mysql:5.7
    environment:
      - MYSQL_ROOT_PASSWORD=mysqlpw
    ports:
      - 3318:3306
    volumes:
      - ./data:/var/lib/mysql2
    networks: 
      - yourNetwork
```

- You will need docker and docker compose installed in your machine
- You should have no problem even if you

## 🚀 Installation

If you have a separate container for database, you should start it now.
After that, Just run `docker compose up` inside the root folder and you should be good to go

## ☕ Usage
You will find the file `endpoints.json` in the root folder, which holds all information you need to make requests. It was generated from insomnia, but it's possible to import in postman as well.

