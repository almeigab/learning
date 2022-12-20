const texto = 'ABC [abc] a-c 1234'

console.log(texto.match(/[a-c]/g))
console.log(texto.match(/a-c/g))
console.log(texto.match(/[A-z]/g))
console.log(texto.match(/[A-Za-z]/g))

// erro de ordem inválida
// console.log(texto.match(/[a-Z]/g))
// console.log(texto.match(/[4-1]/g))