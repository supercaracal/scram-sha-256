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
SCRAM-SHA-256$4096:yTo5lMI+1XyqZOcvYz99Kw==$VJcML25bB3h0xiMUFw9D4spAJwp8IxD1CxnkR7XPty8=:NE05auswTZk1ntaXa8DrO9tYekyhfv1qRMXmugXpGPc=
```

### Testing
```
$ docker run --rm --name=test -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_INITDB_ARGS=--auth-host=scram-sha-256 postgres
```

```
$ docker exec -it test psql -h 127.0.0.1 -U postgres -c "CREATE ROLE test WITH LOGIN PASSWORD 'SCRAM-SHA-256$4096:yTo5lMI+1XyqZOcvYz99Kw==$VJcML25bB3h0xiMUFw9D4spAJwp8IxD1CxnkR7XPty8=:NE05auswTZk1ntaXa8DrO9tYekyhfv1qRMXmugXpGPc='"
Password for user postgres:
CREATE ROLE
```

```
$ docker exec -it test psql -h 127.0.0.1 -U postgres -c 'SELECT usename, passwd FROM pg_catalog.pg_shadow'
Password for user postgres:
 usename  |                                                                passwd                                                                 
----------+---------------------------------------------------------------------------------------------------------------------------------------
 postgres | SCRAM-SHA-256$4096:N+t+PZUQAu25roNaMJiQIw==$MNmcJjqjLwfWBTvKq2zRCWSWPFQX6KnDqqyrqA1XU5g=:jL3qX7jzS4wSP1rOmEbbmLReYL98WeKukK8SfLcdpvU= 
 test     | SCRAM-SHA-256$4096:yTo5lMI+1XyqZOcvYz99Kw==$VJcML25bB3h0xiMUFw9D4spAJwp8IxD1CxnkR7XPty8=:NE05auswTZk1ntaXa8DrO9tYekyhfv1qRMXmugXpGPc= 
(2 rows)
```

```
$ psql -h 127.0.0.1 -p 5432 -U test -W -d postgres -c 'SELECT 1'
Password:
 ?column?
----------
        1
(1 row)
```
