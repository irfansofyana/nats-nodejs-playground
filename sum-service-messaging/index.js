const sum = require('./sum')
const generate = require('./generate-number')

setInterval(() => {
    generate()
}, 1000)

sum()