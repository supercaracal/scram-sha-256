Let's encrypt password with SCRAM-SHA-256 for PostgreSQL
===============================================================================

### Build
```
$ make
```

### Encryption
```
$ ./encrypt
Raw password:
SCRAM-SHA-256$4096:N+VemGods5JLBB35+a+h/w==$Q2xpZW50IEtleYsjoJTmSKUdN1z04fklksJEf9TDTOpXmgr0So7I+eiM:U2VydmVyIEtleYsjoJTmSKUdN1z04fklksJEf9TDTOpXmgr0So7I+eiM
```

### Testing
```
$ docker run --rm --name=test -e POSTGRES_PASSWORD=postgres postgres
```

```
$ docker exec -it test psql -h 127.0.0.1 -U postgres -c "CREATE ROLE test WITH LOGIN PASSWORD 'SCRAM-SHA-256$4096:N+VemGods5JLBB35+a+h/w==$Q2xpZW50IEtleYsjoJTmSKUdN1z04fklksJEf9TDTOpXmgr0So7I+eiM:U2VydmVyIEtleYsjoJTmSKUdN1z04fklksJEf9TDTOpXmgr0So7I+eiM'"
CREATE ROLE
```

```
$ PGPASSWORD=type-here-the-above-raw-password docker exec -it test psql -h 127.0.0.1 -U test -d postgres -c 'SELECT 1'
 ?column?
----------
        1
(1 row)
```
