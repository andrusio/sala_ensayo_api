# Sala de Ensayo App Backend

## For Testing (No DB persistance)
### Run local MySQL instance on docker and expose port
```$ sudo docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_ROOT_HOST=% mysql/mysql-server ```

### Connect with a mysql client to the server
Connect with User and Pass: root \
Run querys from create tables.sql

### Start gingonic server
```go run main.go```


