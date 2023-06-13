"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var pgPromise = require("pg-promise");
var fs = require("fs");
var pgp = pgPromise();
var db = pgp({
    host: 'ep-billowing-sun-767748.us-west-2.aws.neon.tech',
    port: 5432,
    database: '<dbname>',
    user: '<user>>',
    password: '<password>',
    ssl: {
        ca: fs.readFileSync('/path/to/root.crt').toString(),
        rejectUnauthorized: true,
    },
});
db.one('SELECT version(), CURRENT_TIMESTAMP')
    .then(function (result) {
    console.log(result);
})
    .catch(function (error) {
    console.log('ERROR:', error);
})
    .finally(function () {
    // Close the connection after we're done
    pgp.end();
});
