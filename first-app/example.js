const NATS = require('nats')
nc = NATS.connect()

const send = (msg) => {
    console.log(`${msg} is sent!`)
    nc.publish('example', msg);
}

const received = () => {
    nc.subscribe('example', msg => {
        console.log(`${msg} is received!`)
    })
}

setInterval(() => {
    send()
    received()
}, 1000)