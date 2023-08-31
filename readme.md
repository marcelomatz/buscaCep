# BuscaCep
BuscaCep é um microserviço simples e responsivo que permite pesquisar e recuperar informações de um determinado CEP para endereços no Brasil.

### Como isso funciona
O serviço obtém informações de CEP de um endpoint REST da ViaCep, um serviço de CEP brasileiro. As informações do CEP são devolvidas ao cliente no formato JSON.

### Código de Exemplo
```go
package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", BuscaCepHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func BuscaCepHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" || len(cepParam) != 8 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	viaCep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(viaCep)
	if err != nil {
		return
	}

}

func BuscaCep(cep string) (*ViaCep, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var viaCep ViaCep
	err = json.Unmarshal(body, &viaCep)
	if err != nil {
		return nil, err
	}
	return &viaCep, nil
}
```
O serviço recebe um CEP através das query params em um GET Request e retorna um JSON com as informações deste CEP.

#### Exemplo de retorno da requisição
URL de requisição: http://localhost:8080/?cep=01001000

Resposta:
```json
{
    "cep": "01001-000",
    "logradouro": "Praça da Sé",
    "complemento": "lado ímpar",
    "bairro": "Sé",
    "localidade": "São Paulo",
    "uf": "SP",
    "ibge": "3550308",
    "gia": "1004",
    "ddd": "11",
    "siafi": "7107"
}
```

### Como usar
Esta seção fornecerá informações sobre como instalar e executar o projeto.

#### Pré-requisitos
Esse projeto é escrito na linguagem de programação Go. Assim, para rodar localmente o código, é necessário ter o Go instalado no seu sistema e um editor de texto de sua preferência.

#### Instalando
Primeiro, copie este projeto para o seu ambiente local usando o comando 'git clone' com o link do repositório deste projeto. Se você não deseja usar o git, também pode baixar o projeto como arquivo zip do github.
Depois de clonar (ou baixar) o repositório para sua máquina local, navegue até o diretório do projeto.

#### Executando o projeto
Dentro do diretório do projeto, você pode iniciar o servidor de aplicação utilizando o seguinte comando:
```sh
go run main.go
```
Este comando irá iniciar o servidor na porta 8080. Agora, você pode acessar o serviço em http://localhost:8080/ através de seu navegador ou qualquer cliente HTTP, como o Postman.

#### Testando o serviço
Para testar o serviço, você pode utilizar a seguinte URL: http://localhost:8080/?cep=01001000
Lembre-se de substituir 01001000 com o CEP do qual você deseja obter informações.
Com esses passos, você deve ser capaz de instalar e executar o projeto BuscaCep no seu ambiente local de desenvolvimento.

### Versão
Essa é a versão 0.0.1 do BuscaCep, a primeira implementação do serviço. Esta implementação tem fins acadêmicos e seu objetivo é apenas demonstrar o funcionamento do serviço e não deve ser utilizada em produção.

### Contribuindo
Estamos sempre procurando contribuições para melhorar o nosso serviço. Se você tem alguma sugestão ou identificou algum bug, por favor, crie uma issue no nosso repositório.
Todas as Pull Requests devem ser feitas para a branch develop.

### Licença
Este projeto está licenciado sob a licença MIT - consulte o arquivo LICENSE.md para obter detalhes.

### Agradecimentos
* [ViaCep](https://viacep.com.br/) - Serviço público de CEP brasileiro
* [Golang](https://golang.org/) - Linguagem de programação utilizada
* [Goland](https://www.jetbrains.com/pt-br/go/) - IDE utilizada

### Autor
* **Marcelo Matz** - [Github](https://github.com/marcelomatz/) - [LinkedIn](https://www.linkedin.com/in/marcelo-matzembacher-750b51276/) - 
* Email - eumesmo@marcelomatz.com