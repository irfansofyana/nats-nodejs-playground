const nats = require('nats')
const nc = nats.connect()

//this code is retrieved from https://www.npmjs.com/package/nats

// Simple Publisher
nc.publish('foo', 'hello-world')

// Simple Subscriber
nc.subscribe('foo', (msg) => {
    console.log(`Received message: ${msg}`)
})

// unscribe
const sid = nc.subscribe('foo', function (msg) {})
nc.unsubscribe(sid)

// request handle any replies to the request subject and publishes the request with an optional payload
// This usage allows you to collect responses from multiple services
nc.request('help', 'my request', { timeout: 10000 }, (msg) => {
    if (msg instanceof nats.NatsError && msg.code === nats.REQ_TIMEOUT) {
        console.log('request timed out')
    } else {
        console.log('Got a response for help: ' + msg)
    }
})

// subscribe to "help" and publish it again
nc.subscribe('help', function (request, replyTo) {
    console.log(request) // output: "my request"
    nc.publish(replyTo, 'I can help!')
})


// close connection
// nc.close()