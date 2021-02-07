const nc = require('../client')

const getRndInteger = (min, max) => {
    return Math.floor(Math.random() * (max - min) ) + min;
}

const generate = () => {
    const A = getRndInteger(1,1000)
    const B = getRndInteger(1,1000)
    
    console.log(`Publish two numbers: ${A} and ${B}`)

    nc.publish('sum-service-messaging', {
        'A': A,
        'B': B
    })
}

module.exports = generate