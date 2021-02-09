require('dotenv').config()
const NATS = require('nats')

nc = NATS.connect({
    json: true,
    url: process.env.NATS_URI
})

module.exports = nc