# C# PostgreSQL Connection Example using Npgsql

This example shows you how to connect to PostgreSQL using Npgsql in the C# programming language.

## Prerequisites

Ensure you have the following installed on your machine:

- .NET Core SDK (https://dotnet.microsoft.com/download)
- Npgsql .NET data provider for PostgreSQL (This can be installed through NuGet)

## Usage

1. Clone this repository:

```bash
git clone https://github.com/<your_username>/postgresql-connection-examples.git
```

2. Navigate to the directory:

```bash
cd postgresql-connection-examples/csharp
```

3. Create a new console application:

```bash
dotnet new console -n PostgresConnection

4. Navigate to the newly created project:

```bash
cd PostgresConnection
```
5. Install the Npgsql package:

```bash
dotnet add package Npgsql
```

4. Replace the content of the `Program.cs` file with the following code:

```csharp
using System;
using Npgsql;

class Program
{
    static void Main()
    {
        var connString = "Host=<hostname>;Username=<user>;Password=<password>;Database=<dbname>";

        using var conn = new NpgsqlConnection(connString);
        conn.Open();

        using var cmd = new NpgsqlCommand("SELECT version();", conn);
        Console.WriteLine("PostgreSQL version: " + cmd.ExecuteScalar());

        cmd.CommandText = "SELECT current_timestamp;";
        Console.WriteLine("Current time: " + cmd.ExecuteScalar());
    }
}
```

To use an SSL certificate file when connecting, use this code instead:

```csharp
using System;
using Npgsql;

class Program
{
    static void Main()
    {
        var connString = "Host=<hostname>;Username=<user>;Password=<password>;Database=<dbname>;SSL Mode=Require;Trust Server Certificate=true;Root Certificate=/etc/ssl/certs/ca-certificates.crt";

        using var conn = new NpgsqlConnection(connString);
        conn.Open();

        using var cmd = new NpgsqlCommand("SELECT version();", conn);
        Console.WriteLine("PostgreSQL version: " + cmd.ExecuteScalar());

        cmd.CommandText = "SELECT current_timestamp;";
        Console.WriteLine("Current time: " + cmd.ExecuteScalar());
    }
}

```

5. Run the program:

```bash
dotnet run
```

This will print the PostgreSQL version and the current timestamp.

Remember to replace <hostname>, <dbname>, <user>, and <password> in the connString string with your own PostgreSQL credentials.

Always remember to handle database credentials securely in your applications. In a real-world scenario, you should not hardcode your credentials in the code. Instead, consider using environment variables or secure configuration files to handle such sensitive data.