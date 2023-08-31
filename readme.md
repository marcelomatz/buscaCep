# BuscaCep
Este projeto é uma aplicação web simples para buscar informações de CEPs no Brasil. 

A aplicação é dividida em duas partes: um backend que é desenvolvido com Go (Golang) e um frontend simples criado com HTML.

O backend é um servidor Go que realiza consultas a um serviço de CEP para obter informações relacionadas a um determinado CEP.

O frontend é uma interface web simples que permite que o usuário insira um CEP e solicite informações sobre ele. 

Em seguida, ele apresenta esses dados em uma tabela na página.

### Como isso funciona
O serviço obtém informações de CEP de um endpoint REST da ViaCep, um serviço de CEP brasileiro. As informações do CEP são devolvidas ao cliente no formato JSON.

### Configurando e rodando o projeto:

1. Clone o repositório git clone https://github.com/marcelomatz/buscaCep.git
2. Navegue até a pasta do projeto cd buscaCep.
3. Execute go run main.go para iniciar o servidor.
4. Abra seu navegador e vá para http://localhost:8080 para ver a aplicação em funcionamento.

#### Pré-requisitos
Esse projeto é escrito na linguagem de programação Go. Assim, para rodar localmente o código, é necessário ter o Go instalado no seu sistema e um editor de texto de sua preferência.

### Contribuindo
Se você tem alguma sugestão ou identificou algum bug, por favor, crie uma issue no nosso repositório.
Todas as Pull Requests devem ser feitas para a branch develop.

### Licença
Este projeto está licenciado sob a licença MIT - consulte o arquivo LICENSE.md para obter detalhes.

### Agradecimentos
* [ViaCep](https://viacep.com.br/) - Serviço público de CEP brasileiro
* [Golang](https://golang.org/) - Linguagem de programação utilizada
* [Goland](https://www.jetbrains.com/pt-br/go/) - IDE utilizada

### Autor
* **Marcelo Matz** - [Github](https://github.com/marcelomatz/) - [LinkedIn](https://www.linkedin.com/in/marcelo-matzembacher-750b51276/)
* Email - eumesmo@marcelomatz.com