# Demo :: Build docker image and container
* Write a Dockerfile
  * [Building best practices](https://docs.docker.com/build/building/best-practices/)

## 1. ReactJS
```
$docker compose build reactjs
$docker compose up -d reactjs
$docker compose ps
```

Access to http://localhost:8888

## 2. Mirror site of NPM package
* [Verdaccio](https://verdaccio.org/)
* [Sonatype Nexus Repository ](https://help.sonatype.com/en/download.html)
  * Login -> Create a new repository
    * npm (proxy)
      * name=demo01
      * https://registry.npmjs.org/

```
$docker compose up -d mirror
$docker compose ps   
```

Access to http://localhost:8081

### Config npm to download packages from Mirror
```
$npm set registry http://localhost:8081/repository/demo01/
```