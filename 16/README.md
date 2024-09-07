### Usage

```
$ docker image build -t postgres .
$ docker container run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 psql
```

exec

```
$ docker exec -it postgres psql -U postgres -d postgres
psql (16.4 (Debian 16.4-1.pgdg120+1))
Type "help" for help.

postgres=#
```
