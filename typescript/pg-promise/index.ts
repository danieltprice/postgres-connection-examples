import * as pgPromise from 'pg-promise';

const pgp = pgPromise();

const db = pgp({
    host: 'ep-billowing-sun-767748.us-west-2.aws.neon.tech',
    port: 5432,
    database: 'neondb',
    user: '<user>',
    password: '<password>',
    ssl: {
        rejectUnauthorized: false,
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




