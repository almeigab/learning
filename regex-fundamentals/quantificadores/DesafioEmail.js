const texto = `Os e-mails dos convidados são:
    - fulano@cod3r.com.br
    - xico@gmail.com
    - gabriel_almeida1403@hotmail.com`

console.log(texto.match(/\w+@\w+.\w+(.\w+)?/g))