# go-store

## Local development

Run tests:

```
$ make unit_test
```

Run local server:

```
$ make local
...

# in another tab, hit the app port:
$ curl -I -X GET http://localhost:3000/v1/up

# hit the healthz endpoint on the probe port:
$ curl -I -X GET http://localhost:8085/healthz
```