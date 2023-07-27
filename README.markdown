### Initialization
#### 1- Reading config file:
`confgi.yaml` should include infrastructure, exchanges and strategies configurations.
#### 2- Creating domain models:
Relation between domain models described in image below:

![domains](./doc/domains.png)

An array of strategies parsed to config model. Each strategy can be executed on multiple markets in multiple exchanges.
So per each market we create an instance of `Market` and `Exchange` holds the target exchange info.
For each market, one instance of it must exist.
#### 3- Initialize exchanges and other dependencies:
Exchanges have an API and a websocket connection. Websocket connections must be initialized and managed by `ConnectionManager`.
`ConnectionManager` retrieves different exchange connections and revives connections if they got closed.
Other dependencies including InfluxDB, Mysql and redis must be initialized.
#### 4- Initialize repositories
In this step create separate instance for each exchange. Exchanges have to implement `ExchangeRepository` interface.
Other repositories such as `influxRepository`, `cacheRepository` and `mysqlRepository` must be initialized.
#### 5- Initialize strategies and register observers

