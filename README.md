![](https://github.com/supercaracal/scram-sha-256/workflows/Test/badge.svg?branch=master)
![](https://github.com/supercaracal/scram-sha-256/workflows/Release/badge.svg)

A password generator for PostgreSQL
===============================================================================

This is a password generator for PostgreSQL.  
It encrypt a raw string with SCRAM-SHA-256.  

## Installation
```
## Linux AMD64
$ curl -sL https://github.com/supercaracal/scram-sha-256/releases/download/v1.0.0/scram-sha-256_1.0.0_linux_amd64.tar.gz | tar zx -C /tmp
```

If you want to get more information, please see [the release contents](https://github.com/supercaracal/scram-sha-256/releases/tag/v1.0.0).

## Usage
```
$ /tmp/scram-sha-256
Raw password:
SCRAM-SHA-256$4096:Wj5Wd30IrYzvZxJJ5NSuNg==$27HbcXbdKQd55FvzhEFq8DpPKQBRCvcnOOYFJ07Fr7s=:CibjLDWHtzbFjhSejBWcRkRGmuB2njOC8GxDh3gSwrE=
```

```
$ /tmp/scram-sha-256 mysecret
SCRAM-SHA-256$4096:1Iuyc2XTVSv/GFgCWSv9Xw==$nU96dFyIuV+uWwiOly7HU5yinIJh55GsItyFAYrU2sc=:fEC668A2ufIsGS+9WC8xqD0hHvHQBbLiDxZ8hWlwkCw=
```

## Build locally
```
$ git clone https://github.com/supercaracal/scram-sha-256.git
$ cd scram-sha-256/
$ make
$ ./encrypt
Raw password:
SCRAM-SHA-256$4096:Mg8UNqSaPstxvBVRVYPQTw==$Zl7Rhln+rus3z+4YwC+7CgL/uKSUvqWH8mHMUizh1EI=:G9dSawW20CNLxTnZdcwHEHg9U9hG2noNEV2/t7ptq3s=
```

## Testing
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
