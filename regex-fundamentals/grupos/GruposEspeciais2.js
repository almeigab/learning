const texto = 'supermercado superação hiperMERCADO Mercado';

console.log(texto.match(/(?:super)[\wÀ-ú]+/gi));

console.log(texto.match(/(?<=super)[\wÀ-ú]+/gi));

console.log(texto.match(/(?<!super)mercado/gi));
