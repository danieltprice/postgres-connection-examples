# PHP PostgreSQL Connection Example using PDO

This example shows you how to connect to PostgreSQL using PDO in PHP programming language.

## Prerequisites

Ensure you have the following installed on your machine:

- PHP (https://www.php.net/downloads.php)
- PDO Extension for PHP (This usually comes with PHP 5.1.0 and later)
- Composer (https://getcomposer.org/download/)

## Usage

1. Clone this repository:

```bash
git clone https://github.com/<your_username>/postgresql-connection-examples.git
```

2. Navigate to the directory:

```bash
cd postgresql-connection-examples/php-pdo
```

3. Create a `composer.json` file with the following content:

```php
json
Copy code
{
    "require": {
        "ext-pdo": "*"
    }
}
```

4. Run the following command to install the dependencies:

```
composer install
```

5. Create a new file named `index.php` and add the following code to it:

```php
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
```
6. Run the PHP script:

```bash
php index.php
```

This will print the PostgreSQL version and the current timestamp.

Remember to replace <hostname>, <dbname>, <user>, and <password> in the `$dsn` string with your own PostgreSQL credentials.

Always remember to handle database credentials securely in your applications. In a real-world scenario, you should not hardcode your credentials in the code. Instead, consider using environment variables or secure configuration files to handle such sensitive data.

