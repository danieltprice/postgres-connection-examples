Python PostgreSQL Connection Example using psycopg2

This example shows you how to connect to PostgreSQL using psycopg2 in Python programming language. Two scripts are provided, one with `sslmode=verify-full` and one without.

# Prerequisites

Ensure you have the following installed on your machine:

- Python
- pip
- psycopg2

## Usage

1. Clone this repository:

git clone https://github.com/danieltprice/postgresql-connection-examples.git

2. Navigate to the python directory:

```bash
cd postgresql-connection-examples/python
```

3. Install psycopg2:

```bash
pip3 install psycopg2
```

4. Run the program:

```bash
python3 connect.py
```
or

```bash
python3 connect-ssl.py
```

This will print the PostgreSQL version and the current timestamp.

Always remember to handle database credentials securely in your applications. In a real-world scenario, you should not hardcode your credentials in the code. Instead, consider using environment variables or secure configuration files to handle such sensitive data.
