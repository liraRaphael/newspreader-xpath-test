## Newspreader

### O Newspreader é uma API construida em GO com o intuito de trazer o conteudo de artigos de notícia de forma gratuita

***

<br/>

#### Nesse momento a API retorna apenas resultados de artigos do estadão utilizando o endpoint **GET** ```/paynot?url=https://example.com.br```

#### O resultado retornado contém: 
 - Autor;
 - Imagem do artigo;
 - Texto do artigo completo.

<br/>

### Atualmente o projeto está hospedado na ferramenta de hospedagem gratuita [Render](https://render.com) através do endereço https://newspreader.onrender.com

Através do `curl` abaixo, é possivel efetivar o teste do endpoint de onde você estiver:
```
curl -X GET -G 'https://newspreader.onrender.com/paynot' \
-d url=https://www.estadao.com.br/economia/americanas-antecipar-pagamento-dividas-trabalhistas-pequenos-fornecedores/
```

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