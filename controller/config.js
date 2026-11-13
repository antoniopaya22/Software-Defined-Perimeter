// SDP Controller configuration.
// Secrets and connection details come from environment variables (see .env.example)
// so nothing sensitive is committed to the repository.

const env = (name, fallback) => {
    const value = process.env[name];
    if (value === undefined || value === '') {
        if (fallback === undefined) {
            throw new Error('Missing required environment variable ' + name);
        }
        return fallback;
    }
    return value;
};

const certDir = env('CERT_DIR', '/certs');
const controllerId = env('CONTROLLER_SDP_ID', '2');

module.exports = {
    // print debug statements
    'debug': env('CONTROLLER_DEBUG', 'false') === 'true',

    'serverPort': parseInt(env('CONTROLLER_PORT', '5000'), 10),
    'maxConnections': 100,

    // milliseconds, 0 indicates no timeout
    // this is controller's way of noticing a lost connection
    'socketTimeout': 30000,

    // false indicates the server should disconnect
    // after a successful credential update
    'keepClientsConnected': true,

    // Legacy access requests let the SPA packet choose the port and NAT
    // target, which is not secure. Keep disabled.
    'allowLegacyAccessRequests': false,

    // Controller TLS identity, issued by genCredentials.js for its own SDP ID
    'serverCert': certDir + '/' + controllerId + '.crt',
    'serverKey': certDir + '/' + controllerId + '.key',
    // keys produced by genCredentials.js are not encrypted
    'serverKeyPassword': '',
    'serverKeyPasswordRequired': false,

    // CA used to sign every member certificate (created by entrypoint.sh)
    'caCert': certDir + '/ca.crt',
    'caKey': certDir + '/ca.key',
    'caKeyPassword': env('CA_KEY_PASSWORD'),
    'caKeyPasswordRequired': true,

    // how many days new certificates should be good for
    'daysToExpiration': parseInt(env('CERT_DAYS', '31'), 10),

    // SPA key lengths in bytes (genCredentials.js accepts 64 to 256)
    'encryptionKeyLen': 64,
    'hmacKeyLen': 128,

    // database options
    'dbHost': env('DB_HOST', 'sdp-db'),
    'dbUser': env('DB_USER', 'sdp'),
    'dbPasswordRequired': true,
    'dbPassword': env('DB_PASSWORD'),
    'dbName': env('DB_NAME', 'sdp'),

    // if any of these are exceeded, the controller
    // disconnects from the client
    'maxDataTransmitTries': 3,
    'maxCredentialMakerTries': 3,
    'maxBadMessages': 3,

    // retry interval (milliseconds) for database failures
    'databaseRetryInterval': 5000,
    'databaseMaxRetries': 5,

    // interval (milliseconds) to check database for changes
    // that require sending updates to gateways
    'databaseMonitorInterval': 3000,
};
