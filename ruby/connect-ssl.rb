require 'bundler/setup'
require 'sequel'

DB = Sequel.connect('postgres://<user>:<password>@<hostname>/<dbname>?sslmode=verify-full&sslrootcert=/etc/ssl/certs/ca-certificates.crt')

version = DB.fetch('SELECT version()').first[:version]
puts "PostgreSQL version: #{version}"

current_time = DB.fetch('SELECT NOW()').first[:now]
puts "Current timestamp: #{current_time}"
