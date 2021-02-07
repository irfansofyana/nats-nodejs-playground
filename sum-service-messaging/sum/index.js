const nc = require('../client')

const sum = () => {
    nc.subscribe('sum-service-messaging', (msg) => {
        const A = msg.A;
        const B = msg.B;
        
        console.log(`Received message: ${A} + ${B} is ${A+B}`)
    })
}

module.exports = sum;