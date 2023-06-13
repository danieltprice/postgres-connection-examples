# Go PostgreSQL Connection Example using pq client

This example shows you how to connect to PostgreSQL using lib/pq client in Go programming language. Two scripts are provided, one with `sslmode=verify-full` and one without.

## Prerequisites

Ensure you have the following installed on your machine:

- Go (You can install Go from [here](https://go.dev/dl/))

## Usage

1. Clone this repository:

```bash
git clone https://github.com/danieltprice/postgresql-connection-examples.git
```

2. Navigate to the directory:

```bash
cd postgresql-connection-examples/go/pq
```

3. Create a mod file

```bash
go mod init example/gopostgressql
```

4. Get the `pq` package:

```bash
go get github.com/lib/pq
```

3. Run the program:

```bash
go run main.go
```

or

```bash
go run main-ssl.go
```

This will print the PostgreSQL version and the current timestamp.

Always remember to handle database credentials securely in your applications. In a real-world scenario, you should not hardcode your credentials in the code. Instead, consider using environment variables or secure configuration files to handle such sensitive data.


