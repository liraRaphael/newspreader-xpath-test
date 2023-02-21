## Newspreader

### O Newspreader é um serviço construido em GO com o intuito de trazer o conteudo de artigos de notícia de forma gratuita. O serviço entrega as informações em formato API REST e página web.

***

<br/>

#### Nesse momento o serviço retorna resultados de artigos de determinados portais de notícia em formato API REST utilizando o endpoint **GET** ```/api/paynot?url=https://example.com.br```

#### E também retorna uma página web com o conteúdo extraído do link acessando direto pela rota  ```/paynot?url=https://example.com.br```

### Os portais atualmente disponíveis pra consulta são:
 - Estadão
 - Folha de São Paulo
 - CNN Brasil
 - Correio Braziliense

### Planejamento de disponibilização pra novos portais:
 - [ ] R7
 - [ ] Terra
 - [ ] UOL
 - [ ] G1
 - [ ] Exame
 - [ ] InfoMoney
 - [ ] Metropoles
 - [ ] New York Times
 - [ ] BBC News

#### O resultado retornado contém: 
 - Autor;
 - Mídia do artigo;
 - Tipo da mídia do artigo (Vídeo/Imagem);
 - Titulo do artigo;
 - Texto do artigo completo.

<br/>

### Atualmente o projeto está hospedado na ferramenta de hospedagem gratuita [Render](https://render.com) através do endereço https://newspreader.onrender.com

Através do `curl` abaixo, é possivel efetivar o teste do endpoint de onde você estiver:
```
curl -X GET -G 'https://newspreader.onrender.com/api/paynot' \
-d url=https://www.estadao.com.br/economia/americanas-antecipar-pagamento-dividas-trabalhistas-pequenos-fornecedores/
```
<br/>
E através do link a seguir é possível visualizar a página web com o conteúdo:
https://newspreader.onrender.com/paynot?url=https://www.estadao.com.br/economia/americanas-antecipar-pagamento-dividas-trabalhistas-pequenos-fornecedores/

<br/>

### O **Newspreader** foi realizado como projeto de estudo da linguagem GO. Como base de aprendizado, segue abaixo os sites, artigos e/ou docs que auxiliaram no processo:
 - [GO by example - URL Parsing](https://gobyexample.com/url-parsing)
 - [GO by example - If-else](https://gobyexample.com/if-else)
 - [Fiber](https://docs.gofiber.io)
 - [Fiber - Error Handling](https://docs.gofiber.io/guide/error-handling/)
 - [Geeks for geeks - GO operators](https://www.geeksforgeeks.org/go-operators/)
 - [GO Colly - Getting started](http://go-colly.org/docs/introduction/start/)
 - [GO Colly - Web Scraping tips](https://go-colly.org/articles/scraping_tips/)
 - [Scraping Bee - Web Scraping With Go](https://www.scrapingbee.com/blog/web-scraping-go/)