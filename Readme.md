## hello-world-env

A docker image for tinkering with environment variables. It will try print out all the environment variables that begin with a particular prefix. The `PREFIX` environment variable has been kept mandatory intentionally. Optionally one can also choose to trim the `PREFIX` while printing by setting the `TRIM` environment variable.

```
rounakdatta/hello-world-env
```

### Using with environment variables (adding a prefix)

```
docker run --publish=3000:3000 --rm -e PREFIX=FOXY_ -e FOXY_ALLCAPS=400 -e FOXY_sweet=6996 rounakdatta/hello-world-env
```

```
GET http://localhost:3000
```

Prints

```
FOXY_ALLCAPS : 400
FOXY_sweet : 6996
```

### Using with environment variables (adding a prefix, but trimming out while printing)

```
docker run --publish=3000:3000 --rm -e TRIM=yes -e PREFIX=FOXY_ -e FOXY_ALLCAPS=400 -e FOXY_sweet=6996 rounakdatta/hello-world-env
```

```
GET http://localhost:3000
```

Prints

```
ALLCAPS : 400
sweet : 6996
```
