mysql:
	docker run --name demodb -p 3306:3306 -e MYSQL_PASSWORD=secret -e MYSQL_ROOT_PASSWORD=secret -d mysql:8.0-debian
createdb:
	docker exec -it demodb mysql -uroot -psecret -e "CREATE DATABASE demodb"
dropdb:
	docker exec -it demodb mysql -uroot -psecret -e "DROP DATABASE demodb"
new-service:
	goctl api new -home "." ./account/api
gen-service:
	goctl api go -api ./account/api/api.api -dir ./account/api
gen-model:
	goctl model mysql ddl -src="./account/model/*.sql" -dir="./account/model"
run:
	go run ./account/api/api.go -f ./account/api/etc/api-api.yaml