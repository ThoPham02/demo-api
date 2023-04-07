mysql:
	docker run --name demodb -p 3306:3306 -e MYSQL_PASSWORD=secret -e MYSQL_ROOT_PASSWORD=secret -d mysql:8.0-debian
createdb:
	docker exec -it demodb mysql -uroot -psecret -e "CREATE DATABASE demodb"
dropdb:
	docker exec -it demodb mysql -uroot -psecret -e "DROP DATABASE demodb"
new-account-service:
	goctl api new -home "." ./service/account/api
new-topic-service:
	goctl api new -home "." ./service/topic/api
gen-account-service:
	goctl api go -api ./api/account.api -dir ./service/account/api
gen-topic-service:
	goctl api go -api ./api/topic.api -dir ./service/topic/api
gen-account-model:
	goctl model mysql ddl -src="./schema/account.sql" -dir="./service/account/model"
gen-session-model:
	goctl model mysql ddl -src="./schema/session.sql" -dir="./service/account/model"
gen-topic-model:
	goctl model mysql ddl -src="./schema/topic.sql" -dir="./service/topic/model"
run:
	go run main.go -f ./etc/server.yaml