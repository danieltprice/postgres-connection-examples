# PostgreSQL Connection Example using psql

This example shows you how to connect to PostgreSQL using `psql`.

## Prerequisites

Ensure you have the following installed on your machine:

- psql 

## Usage

To connect:

```bash
psql 'postgres://<user>@<hostname>/<dbname>'
```

To connect using `sslmode=verify-full`:

```bash
PGSSLROOTCERT=/path/to/cert.pem psql 'postgres://<user>@<hostname>/<dbname>?sslmode=verify-full'
```

### Root certificate file locations

See [Location of system root certificates](https://neon.tech/docs/connect/connect-securely#location-of-system-root-certificates).