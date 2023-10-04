# FizzBuzz em Golang: Deploy para Azure Web App

Este é um projeto simples que implementa o famoso jogo FizzBuzz em Golang. O principal objetivo deste projeto é fornecer uma base para praticar o processo de implantação de aplicativos Golang em contêineres no Azure Web App.

## Descrição

O jogo FizzBuzz é uma atividade de programação simples, onde os números de 1 a N são iterados, e para cada número, certas regras são aplicadas. Se o número for divisível por 3, ele é substituído por "Fizz", se for divisível por 5, é substituído por "Buzz" e se for divisível por ambos, é substituído por "FizzBuzz". Caso contrário, o número permanece inalterado.

Este projeto é uma implementação mínima do FizzBuzz como uma aplicação Golang, com um único endpoint HTTP para jogar FizzBuzz. No entanto, seu objetivo principal é servir como um exemplo para a implantação em um ambiente Azure Web App.

## Como Implantar no Azure Web App

1. **Construir e Enviar a Imagem Docker para Azure Container Registry (ACR)**

    ```bash
    docker build -t felippegaede.azurecr.io/golang-fizzbuzz:v1 .
    ```

    ```bash
    docker login felippegaede.azurecr.io
    ```

    Envie a imagem para o Azure Container Registry (ACR):

    ```bash
    docker push felippegaede.azurecr.io/golang-fizzbuzz:v1
    ```

2. **Configurar o Azure Web App**

    - Abra o Azure Portal e acesse o serviço Azure Web App que você deseja implantar.
    
    - Vá para "Deployment Center" e siga as instruções para configurar a integração com seu ACR. Isso permitirá que o Web App Service puxe automaticamente a imagem Docker do ACR e a implante.
    
    - Configure manualmente quaisquer outras configurações necessárias, como variáveis de ambiente, se aplicável.
    
    - Inicie a implantação e aguarde até que o aplicativo esteja em execução.

**Aviso:**

Lembre-se de que este projeto é puramente didático e destinado a fins de aprendizado. O projeto SERÁ aprimorado para incluir um processo automatizado de CI/CD, como o GitHub Actions, para simplificar o processo de implantação.
