RabbitMQProducer
-----------------------------

### Setup libraries
```
glide install
```

### RUN
```
go run main.go 
```

defaults:
* username  - `guest`
* passowrd  - `guest`
* host      - `localhost`
* port      - `5672`
* message   - `test message`


##### run with parameters:

```
go run main.go -u admin -p p@ssw0rd -host example.com -port 5672 -message "test message"
```