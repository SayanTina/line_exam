# Web check service status

this project requires the following software to be installed locally.

- git
- docker

## Project setup backend (Golang) and run

Change directory to backend folder.
```
cd /backend
```
Run command to build images.
```
docker build --tag docker-backend-go .
```
run command to run image as a container.
```
docker run --publish 8080:8080 docker-backend-go
```
Now backend-go service ready on localhost:8080

## Project setup frontend (Vue) and run

Change directory to frontend folder.
```
cd /frontend/line_exam
```
Run command to build images.
```
docker build --tag docker-frontend-vuejs .
```
run command to run image as a container.
```
docker run --publish 8082:8080 docker-frontend-vuejs
```
Now you can access web frontend via your browser
```
localhost:8082
```
