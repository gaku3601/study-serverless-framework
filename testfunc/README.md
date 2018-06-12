# deploy
GOOS=linux go build -o bin/main ./hello
GOOS=linux go build -o bin/hello2 ./hello2
GOOS=linux go build -o bin/authHello ./authHello
sls deploy
