# it-life-api

Repository for backend of project "ITLife".

## ITLife API

An awesome API for project "ITLife".

### Requirements

| tool | version |
| --- | --- |
| Go | ^1.18 |

### Dependencies

See modules list in `go.mod`.

### Setup in local environments

1. Start application

    ```sh
    docker-compose build
    docker-compose up --detach
    ```

    Server starts on localhost.
    See "http://127.0.0.1:8000/".

### Deployment

You can deploy manually by the bellow command.  

```sh
heroku container:push web
heroku container:release web
```


## ITLife Database

### Structure

- it_life_database/data

    Raw data to preserve in a container.

- it_life_database/definition

    Initialization DDLs.
    They should be always synced with production database schema.

### Usage in local environments

- Initialize (Reinitialize)

    ```sh
    rm -rfi it_life_database/data/*
    docker-compose up --detach --build
    ```

- Start

    ```sh
    docker-compose up --detach
    ```

- Stop

    ```sh
    docker-compose down
    ```

- Destroy

    ```sh
    rm -rfi it_life_database/data/*
    docker-compose down --rmi all --volumes
    ```

