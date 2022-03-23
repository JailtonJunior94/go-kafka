## Golang + SQL Server + Kafka + Debezium + Slack
<br>

### Cadastrando um novo conector 
POST http://localhost:8083/connectors
```
{
	"name": "sqlserver-customers-connector",
	"config": {
        "connector.class": "io.debezium.connector.sqlserver.SqlServerConnector",
        "tasks.max": 1,
		"database.hostname": "mssql",
        "database.port": "1433",
		"database.user": "sa",
        "database.password": "@docker@2021",
		"database.dbname": "KafkaPoC",
		"database.server.name": "kafka_poc_server",
        "table.include.list": "dbo.Customers",
		"database.history.kafka.bootstrap.servers": "kafka:9092",
		"database.history.kafka.topic": "dbhistory.customers"
	}
}
``` 