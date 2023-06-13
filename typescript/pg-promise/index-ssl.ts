import * as pgPromise from 'pg-promise';
import * as fs from 'fs';

const pgp = pgPromise();

const db = pgp({
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
    .then(result => {
        console.log(result);
    })
    .catch(error => {
        console.log('ERROR:', error);
    })
    .finally(() => {
        // Close the connection after we're done
        pgp.end();
    });
