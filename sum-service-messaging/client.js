const NATS = require('nats')

nc = NATS.connect({
    json: true
})

module.exports = nc