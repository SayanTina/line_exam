# Web check service status

## Project setup backend (Golang) and run
```
cd /backend
```
```
docker build --tag docker-backend-go .
```
```
docker run --publish 8080:8080
```

## Project setup frontend (Vue) and run

```
cd /frontend/line_exam
```
```
docker build --tag docker-frontend-vuejs
```
```
docker run -v ${PWD}:/app -v /app/node_modules -p 8081:8080 --rm docker-frontend-vuejs
```
