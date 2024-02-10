![](https://github.com/supercaracal/scram-sha-256/workflows/Test/badge.svg?branch=master)
![](https://github.com/supercaracal/scram-sha-256/workflows/Release/badge.svg)

SCRAM-SHA-256
===============================================================================

This is a password-encryption tool for PostgreSQL with SCRAM-SHA-256.

## Installation
Please download an executable file from [release pages](https://github.com/supercaracal/scram-sha-256/releases).

## Usage
```
$ scram-sha-256
Raw password:
SCRAM-SHA-256$4096:Wj5Wd30IrYzvZxJJ5NSuNg==$27HbcXbdKQd55FvzhEFq8DpPKQBRCvcnOOYFJ07Fr7s=:CibjLDWHtzbFjhSejBWcRkRGmuB2njOC8GxDh3gSwrE=
```

```
$ scram-sha-256 mysecret
SCRAM-SHA-256$4096:1Iuyc2XTVSv/GFgCWSv9Xw==$nU96dFyIuV+uWwiOly7HU5yinIJh55GsItyFAYrU2sc=:fEC668A2ufIsGS+9WC8xqD0hHvHQBbLiDxZ8hWlwkCw=
```

```
$ echo -n mysecret | scram-sha-256
SCRAM-SHA-256$4096:67e60Pre+3h6dhUm+K2tWA==$MRZtokLiZoWqNLf05HKH7STvtAtWEOy1CZU+vg9hj/M=:jzbp7PPDFT8aBPuFk91KBO2HswNJrvMMuMkUgR1LClI=
```

```
$ echo mysecret | scram-sha-256
SCRAM-SHA-256$4096:wvtRpXoTijsOR2py/yjIjQ==$iQV2GGKBAnN3v339hDOSZWxbl7YH8I3ERh+RCHjOqGQ=:Ea9Pyj4/IR53wmdCISCIOsSINUirJzz6EzD0NJqa05M=
```

```go
import "github.com/supercaracal/scram-sha-256/pkg/pgpasswd"

func doSomething(rawPassword []byte) error {
    encrypted, err := pgpasswd.Encrypt(rawPassword)
    if err != nil {
        return err
    }

    // do something
    return nil
}
```

## Build locally
```
$ git clone https://github.com/supercaracal/scram-sha-256.git
$ cd scram-sha-256/
$ make
$ ./cmd/tool/encrypt
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
