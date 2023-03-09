const texto1 = 'De longe avistei o fogo e uma pessoa gritando: FOGOOOOOO!'

const texto2 = 'There is a big fog in NYC'

const regex = /fogo?/gi

console.log(texto1.match(regex))
console.log(texto2.match(regex))