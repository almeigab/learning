const texto = `CPF dos aprovados:
    - 600.567.890-12
    - 765.998.345-23
    - 395.659.342-65`

console.log(texto.match(/\d{3}\.\d{3}\.\d{3}-\d{2}/g))
console.log(texto.match(/(\d{3}\.){2}\d{3}-\d{2}/g))
