#include <stdio.h>
#include <stdlib.h>
#include <libpq-fe.h>

void exit_nicely(PGconn *conn) {
    PQfinish(conn);
    exit(1);
}

int main() {
    PGconn     *conn;
    PGresult   *res;

    /* Specify the connection string */
    const char *conninfo = "postgres://<user>:<password>@ep-autumn-voice-196875.us-east-2.aws.neon.tech/neondb";

    /* Make a connection to the database */
    conn = PQconnectdb(conninfo);

    /* Check to see that the backend connection was successfully made */
    if (PQstatus(conn) != CONNECTION_OK) {
        fprintf(stderr, "Connection to database failed: %s", PQerrorMessage(conn));
        exit_nicely(conn);
    }

    /* Do SQL query and get result */
    res = PQexec(conn, "SELECT version(), current_timestamp");
    if (PQresultStatus(res) != PGRES_TUPLES_OK) {
        printf("No data retrieved\n");        
        PQclear(res);
        exit_nicely(conn);
    }

    /* Print the version and current time */
    printf("PostgreSQL version: %s\n", PQgetvalue(res, 0, 0));
    printf("Current time: %s\n", PQgetvalue(res, 0, 1));

    /* Clear the result */
    PQclear(res);

    /* Close the connection to the database and cleanup */
    PQfinish(conn);

    return 0;
}
