import psycopg2
import sys

def connect():
    conn = None
    try:
        # Connection string with a 10-second timeout
        conn_str = "postgres://<user>:<password>@<hostname>/<dbname>?connect_timeout=10"
        
        # connect to the PostgreSQL server
        conn = psycopg2.connect(conn_str)
        
        # create a cursor
        cur = conn.cursor()
        
        # execute a statement
        cur.execute('SELECT version(), current_timestamp')
        
        # display the PostgreSQL version and current timestamp
        db_version = cur.fetchone()
        print(db_version)
        
        # close the communication with the PostgreSQL
        cur.close()
    except (Exception, psycopg2.DatabaseError) as error:
        print(error)
    finally:
        if conn is not None:
            conn.close()
            print('Database connection closed.')

if __name__ == '__main__':
    connect()
