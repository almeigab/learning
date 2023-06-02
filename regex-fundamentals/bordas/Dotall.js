const texto = 'Romário era um excelente jogador\n, mas hoje é um político questionador';

console.log(texto.match(/^r.*r$/gi));
console.log(texto.match(/^r[\w\W]*r$/gi));