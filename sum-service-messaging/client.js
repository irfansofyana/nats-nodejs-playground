require('dotenv').config()
const NATS = require('nats')

console.log(process.env.NATS_URL)

nc = NATS.connect({
    json: true,
    url: process.env.NATS_URL
})

module.exports = nc