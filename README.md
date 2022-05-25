# Web check service status

## Project setup backend (Golang) and run

### change directory to backend folder
```
cd /backend
```
### docker build images
```
docker build --tag docker-backend-go .
```
### docker run image as a container
```
docker run --publish 8080:8080 docker-backend-go
```

## Project setup frontend (Vue) and run

### change directory to frontend folder
```
cd /frontend/line_exam
```
### docker build images
```
docker build --tag docker-frontend-vuejs
```
### docker run image as a container
```
docker run -v ${PWD}:/app -v /app/node_modules -p 8081:8080 --rm docker-frontend-vuejs
```
