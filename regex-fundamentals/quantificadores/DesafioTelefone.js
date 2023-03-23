const texto = `Lista telefônica:
    - (11) 98756-1212
    - 98756-1212
    - (11) 8756-1212
    - 8756-1212`

console.log(texto.match(/(\(\d{2}\) )?\d{4,5}-\d{4}/g))