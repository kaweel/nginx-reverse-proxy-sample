GOOS=linux GOARCH=amd64 go build -o serv-1 -v main.go
docker build . -t simple_serve_1
docker run --name serv1 -d simple_serve_1:latest

GOOS=linux GOARCH=amd64 go build -o serv-2 -v main.go
docker build . -t simple_serve_2
docker run --name serv2 -d simple_serve_2:latest

docker run --name simple_proxy \
-p 80:80 \
-p 81:81 \
-it \
--link serv1 \
--link serv2 \
-v $HOME/Develop/docker/nginx/micro-nginx/nginx/conf.d:/etc/nginx/conf.d \
nginx:1.15.5-alpine
