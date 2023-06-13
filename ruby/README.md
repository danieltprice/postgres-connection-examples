# Ruby PostgreSQL Connection Example using Sequel

This example shows you how to connect to PostgreSQL using Sequel in Ruby programming language with SSL.

## Prerequisites

Ensure you have the following installed on your machine:

- Ruby. You can install Ruby on Ubuntu with the following command:

```bash
sudo apt-get install ruby-full
```

- Bundler:

```bash
sudo gem install bundler
```

## Usage

1. Clone this repository:

```bash
git clone https://github.com/<your_username>/postgresql-connection-examples.git
```
2. Navigate to the directory:

```bash
cd postgresql-connection-examples/ruby
```

3. In this directory, create a file called Gemfile and add the following lines to it:

```ruby
source 'https://rubygems.org'
gem 'pg'
gem 'sequel'
```

4. Set up a local gem installation directory for your project by creating a .bundle directory:

```bash
mkdir .bundle
```

5. Create a configuration file for bundler:

```bash
echo 'BUNDLE_PATH: vendor/bundle' > .bundle/config
```

This tells Bundler to install the gems in the `vendor/bundle` directory inside your project.

6. Run the bundle install command:

```bash
bundle install
```

5. Run the program:

```bash
ruby connect.rb
```

or 

```bash
ruby connect-ssl.rb
```

This will print the PostgreSQL version and the current timestamp.

Always remember to handle database credentials securely in your applications. In a real-world scenario, you should not hardcode your credentials in the code. Instead, consider using environment variables or secure configuration files to handle such sensitive data.