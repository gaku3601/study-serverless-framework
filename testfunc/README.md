# deploy
GOOS=linux go build -o bin/main ./hello
GOOS=linux go build -o bin/hello2 ./hello2
GOOS=linux go build -o bin/authHello ./authHello
GOOS=linux go build -o bin/helloDynamo ./helloDynamo
sls deploy

# 切り替え
sls deploy --stage prod
