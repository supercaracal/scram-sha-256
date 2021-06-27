![](https://github.com/supercaracal/scram-sha-256/workflows/Test/badge.svg?branch=master)
![](https://github.com/supercaracal/scram-sha-256/workflows/Release/badge.svg?branch=master)

Let's encrypt password for PostgreSQL using SCRAM-SHA-256
===============================================================================

### Build
```
$ make
```

### Encryption
```
$ ./encrypt
Raw password:
SCRAM-SHA-256$4096:Mg8UNqSaPstxvBVRVYPQTw==$Zl7Rhln+rus3z+4YwC+7CgL/uKSUvqWH8mHMUizh1EI=:G9dSawW20CNLxTnZdcwHEHg9U9hG2noNEV2/t7ptq3s=
```

### Testing
```
$ docker run --rm --name=test -e POSTGRES_PASSWORD=postgres -e POSTGRES_INITDB_ARGS=--auth-host=scram-sha-256 -e PGUSER=postgres postgres
```

```
$ docker exec -it test bash -c 'cat | psql'
## type the following SQL
CREATE ROLE test WITH LOGIN PASSWORD 'SCRAM-SHA-256$4096:Mg8UNqSaPstxvBVRVYPQTw==$Zl7Rhln+rus3z+4YwC+7CgL/uKSUvqWH8mHMUizh1EI=:G9dSawW20CNLxTnZdcwHEHg9U9hG2noNEV2/t7ptq3s='
## press Ctrl-D
CREATE ROLE
```

```
$ docker exec -it test psql -c 'SELECT usename, passwd FROM pg_catalog.pg_shadow'
 usename  |                                                                passwd                                                                 
----------+---------------------------------------------------------------------------------------------------------------------------------------
 postgres | SCRAM-SHA-256$4096:N+t+PZUQAu25roNaMJiQIw==$MNmcJjqjLwfWBTvKq2zRCWSWPFQX6KnDqqyrqA1XU5g=:jL3qX7jzS4wSP1rOmEbbmLReYL98WeKukK8SfLcdpvU= 
 test     | SCRAM-SHA-256$4096:Mg8UNqSaPstxvBVRVYPQTw==$Zl7Rhln+rus3z+4YwC+7CgL/uKSUvqWH8mHMUizh1EI=:G9dSawW20CNLxTnZdcwHEHg9U9hG2noNEV2/t7ptq3s= 
(2 rows)
```

```
$ docker exec -it test psql -h 127.0.0.1 -U test -W -d postgres -c 'SELECT 1'
Password:
 ?column?
----------
        1
(1 row)
```

```
$ docker stop test
```
