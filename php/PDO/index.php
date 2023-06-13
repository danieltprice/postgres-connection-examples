<?php
$dsn = 'pgsql:host=<hostname>;dbname=<dbname>;user=<user>;password=<password>';

try {
    $dbh = new PDO($dsn);

    $version = $dbh->query('SELECT version()')->fetchColumn();
    echo "PostgreSQL version: $version\n";

    $current_time = $dbh->query('SELECT NOW()')->fetchColumn();
    echo "Current Time: $current_time\n";
    
} catch (PDOException $e) {
    echo "Connection failed: " . $e->getMessage();
}
?>

