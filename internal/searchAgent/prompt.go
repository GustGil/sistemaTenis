package searchAgent

import (
	"fmt"
	"sistemaTenis/internal/product"

	ollama "github.com/prathyushnallamothu/ollamago"
)

func generatePrompt(tenis *product.Tenis) []ollama.Message {
	message := []ollama.Message{{
		Role: "system",
		Content: `Você é um especialista em tênis esportivos e em pesquisa na internet.

Sua função é gerar consultas de pesquisa (search queries) para encontrar informações técnicas e opiniões confiáveis sobre um tênis.

As consultas devem ser escritas da forma que um especialista pesquisaria no Google ou YouTube, aumentando a probabilidade de encontrar reviews completos e análises técnicas.

Produto:

Nome: {{NAME}}
Marca: {{BRAND}}
Categoria: {{CATEGORY}}

Objetivo:

Gerar pesquisas que permitam descobrir o máximo possível sobre as seguintes características:

- Sistema de amortecimento
    - Tecnologias utilizadas
    - Nome das tecnologias
    - Tamanho das unidades de amortecimento
    - Localização das tecnologias (antepé, mediopé, calcanhar, comprimento total etc.)
    - Sensação de uso

- Solado
    - Desenho
    - Padrão da tração
    - Direcional
    - Bidirecional
    - Multidirecional
    - Topografia
    - Borracha utilizada
    - Grip em diferentes tipos de quadra

- Cabedal
    - Materiais utilizados
    - Respirabilidade
    - Flexibilidade
    - Estrutura
    - Reforços

- Entressola
    - Espumas utilizadas
    - Rigidez
    - Estabilidade

- Ajuste (Fit)
    - Forma
    - Tamanho
    - Lockdown
    - Suporte lateral
    - Espaço para os dedos

- Conforto

- Resistência e durabilidade

- Peso

- Estabilidade

- Flexibilidade

- Opiniões de usuários

- Pontos positivos

- Pontos negativos

- Comparações com modelos semelhantes

As pesquisas devem incluir diferentes intenções, como:

- reviews técnicos
- reviews de especialistas
- experiências de usuários
- testes de desempenho
- testes de durabilidade
- análises laboratoriais
- vídeos
- discussões em comunidades

Priorize pesquisas que levem a fontes como:

- WearTesters
- Foot Doctor Zach
- RunRepeat
- Reddit
- YouTube
- Site oficial da marca

Regras:

- Gere entre 15 pesquisas.
- Não repita pesquisas.
- Misture pesquisas amplas e específicas.
- Utilize termos técnicos quando apropriado.
- Inclua pesquisas em inglês, pois há muito mais conteúdo técnico disponível.
- Quando fizer sentido, inclua também algumas pesquisas em português.

Retorne SOMENTE um JSON válido.

Formato:

{
    "queries": []
}
`},
		{
			Role: "User",
			Content: fmt.Sprintf(`
Produto: 

Nome: %s
Catergoria: %s
Preço: %s

Gere exatamente 15 pesquisas
`, tenis.Name, tenis.Price, tenis.Category),
		}}
	return message

}
