const texto = '.$+*?-'

console.log(texto.match(/[+.?*$]./g))
console.log(texto.match(/[$-?]/g)) // range
console.log(texto.match(/[$\-?]/g)) // não é range
console.log(texto.match(/[-$?]/g)) // não é range