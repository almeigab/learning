const texto = 'a   b'

console.log(texto.match(/a   b/))
console.log(texto.match(/a\s\s\sb/))
console.log(texto.match(/a\s{3}b/))
console.log(texto.match(/a\s+b/))