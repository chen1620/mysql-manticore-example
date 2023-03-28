# mysql + manticore-search example

# How to setup local environment
- Download all go modules by
```shell
GOSUMDB=off go mod download
```
- Starts infra environment by running, this would use docker-compose (with file `docker-compose.yaml`) to startup
  all necessary infra components:
```shell
make infra
```

- Run the components by make command:
    - Run the app in the development mode (8080 is default HTTP port): `go run main.go`

# How to setup database schema
- For mysql:
```sql
create table users
(
    id        bigint unsigned auto_increment
        primary key,
    user_name varchar(255) null,
    password  varchar(255) null
);

create table posts
(
    id         bigint unsigned auto_increment
        primary key,
    title      text        null,
    content    longtext        null,
    user_id    bigint unsigned null,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    constraint fk_posts_user
        foreign key (user_id) references users (id)
);
```

- For manticore-search:
```sql
CREATE TABLE posts(title text, content text);
```