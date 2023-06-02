# libpq PostgreSQL Connection Example

This example shows you how to connect to PostgreSQL using `libpq` in C programming language.

## Prerequisites

Ensure you have the following installed on your machine:

- PostgreSQL
- libpq (PostgreSQL C API)
- GCC (GNU Compiler Collection)

If you're an Ubuntu user, you can install these with the following commands:

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo apt install libpq-dev
sudo apt install build-essential
```

## Usage

1. Clone this repository:

```bash
git clone https://github.com/<your_username>/postgresql-connection-examples.git
```

2. Navigate to the directory:

```bash
cd postgresql-connection-examples
```

3. Compile the C program:

```bash
gcc connect_pg.c -o connect_pg -I /usr/include/postgresql -lpq
```

4. Run the program:

```bash
./connect_pg
```

This will print the PostgreSQL version and the current timestamp.
