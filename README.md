This repo contains the source code for a simple Hello World application, written in Go. The app uses [GORM](https://gorm.io/) to connect to an existing [CockroachDB](https://www.cockroachlabs.com/docs/stable/) cluster.

To run the code:

1. Start a [local CockroachDB cluster](https://www.cockroachlabs.com/docs/stable/secure-a-cluster.html), or use [CockroachCloud](https://www.cockroachlabs.com/docs/cockroachcloud/create-a-free-cluster.html).
2. From the command line, execute the following:

    
    ~~~
    go mod init basic-sample && go mod tidy
    ~~~

    ~~~
    go run main.go
    ~~~

3. Enter the connection string for the cluster.

