const texto1 = 'De longe avistei o fogo e uma pessoa gritando: FOGOOOOOO!'

const texto2 = 'There is a big fog in NYC'

const regex = /fogo+/gi

console.log(texto1.match(regex))
console.log(texto2.match(regex))

const texto3 = 'Os números: 0123456789'
console.log(texto3.match(/\d/g))
console.log(texto3.match(/\d+/g))