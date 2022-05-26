# Web check service status

this project requires the following software to be installed locally"

- git
- docker

## Project setup backend (Golang) and run

change directory to backend folder
```
cd /backend
```
docker build images
```
docker build --tag docker-backend-go .
```
docker run image as a container
```
docker run --publish 8080:8080 docker-backend-go
```

## Project setup frontend (Vue) and run

change directory to frontend folder
```
cd /frontend/line_exam
```
docker build images
```
docker build --tag docker-frontend-vuejs .
```
docker run image as a container
```
docker run docker-frontend-vuejs
```
