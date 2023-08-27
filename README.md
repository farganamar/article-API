# alpha-indo-soft-backend-test
This project using Golang `1.21.0` version

# How to use
On terminal:

run DB and redis from docker

    docker compose up -d 


then run golang Restful API

    go run main.go 


# API
 To create new Article, copy and paste this cURL: 
```
curl --location 'localhost:8080/articles' \
--header 'Content-Type: application/json' \
--data '{
    "author": "Fargan",
    "Title": "Alpha indo",
    "Body": "Alpha indo article"
}'
```
  
 To Search article by some query param, copy and paste this cURL: 

 ```curl --location 'localhost:8080/articles?author=Fargan&keyword=Alpha%20Indo'```