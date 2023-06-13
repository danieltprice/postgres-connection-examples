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
