# kbeja - save metrics to redis/mongodb/postgresql

To install run:

```sh
$ go get github.com/erubboli/kbeja
```

### Run the workers
Worker `accountName` save the accounts to postgresql
```sh
$ kbeja -accountName
```
Worker `distinctName` count the metrics in redis and archive in a monthly bucket metrics older than 30 days
```sh
$ kbeja -distinctName
```

Worker `hourlyLog` save the log into mongodb
```sh
$ kbeja --hourlyLog
```

to send a test metric to the system use:

```sh
$ kbeja --sendMessage account_name metric_name
```



