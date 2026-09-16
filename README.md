<div align="center">

# StudyGolang

**Laboratório pessoal de estudos em Go** — anotações práticas, algoritmos e resoluções comentadas do LeetCode.

[![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![LeetCode](https://img.shields.io/badge/LeetCode-2%20resolvidos-FFA116?style=for-the-badge&logo=leetcode&logoColor=white)](https://leetcode.com/)
[![Status](https://img.shields.io/badge/status-em%20evolução-2ea44f?style=for-the-badge)](#roadmap)

</div>

---

## Sobre

Este repositório documenta minha jornada de aprendizado em **Go**. A prioridade não é a solução mais curta ou mais "esperta" — é entender **por que** ela funciona.

Por isso cada resolução traz:

- **Comentários em português**, escritos como material de revisão
- **Complexidade explícita** de tempo e espaço
- **Foco nos idiomas da linguagem** — slices, ponteiros, `switch` sem condição, limites numéricos

---

## Exercícios resolvidos

| # | Problema | Dificuldade | Técnica | Tempo | Espaço |
|:---:|---|:---:|---|:---:|:---:|
| 8 | [String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/) | ![Medium](https://img.shields.io/badge/Medium-FFA116?style=flat-square) | Parsing manual + saturação | `O(n)` | `O(1)` |
| 11 | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) | ![Medium](https://img.shields.io/badge/Medium-FFA116?style=flat-square) | Dois ponteiros | `O(n)` | `O(1)` |

<details>
<summary><b>Notas de implementação</b></summary>

<br>

**8 · String to Integer (atoi)** — [`main.go`](main.go)

O ponto delicado não é o parsing, é o **overflow**. A verificação acontece _dentro_ do laço, a cada dígito acumulado, comparando contra `math.MaxInt32` / `math.MinInt32` antes que o valor estoure. Em uma plataforma de 64 bits o `int` de Go aguenta o valor intermediário, o que dá a folga necessária para detectar o estouro e saturar no limite — comportamento exigido pelo enunciado.

**11 · Container With Most Water** — [`Leetcode/ContainerWithMostWater/main.go`](Leetcode/ContainerWithMostWater/main.go)

A força bruta seria `O(n²)`. A solução usa **dois ponteiros** nas extremidades e sempre move o da **menor** barra para dentro. A intuição: a área é limitada pela barra mais baixa, então mover a mais alta só pode reduzir a largura sem nunca aumentar a altura útil — logo, nenhuma solução melhor é descartada. Isso derruba a busca para `O(n)`.

</details>

---

## Estrutura

```
StudyGolang/
├── go.mod                              # módulo raiz — module estudos
├── main.go                             # bancada de trabalho / exercício da vez
└── Leetcode/
    └── ContainerWithMostWater/
        ├── go.mod                      # módulo independente
        └── main.go
```

> [!NOTE]
> Cada exercício vive em sua própria pasta com um `package main` isolado. Isso evita colisão entre funções de mesmo nome e permite executar qualquer solução sem tocar nas outras.

---

## Como executar

**Pré-requisito:** [Go 1.24+](https://go.dev/dl/)

```bash
# clonar
git clone https://github.com/KaioRBraga/StudyGolang.git
cd StudyGolang

# rodar o exercício da raiz
go run main.go

# rodar um exercício específico
cd Leetcode/ContainerWithMostWater && go run main.go
```

> [!WARNING]
> Como as pastas de exercício têm `go.mod` próprio, elas são **submódulos**. Um `go build ./...` na raiz **não** as alcança — compile dentro da pasta do exercício.

<details>
<summary><b>Compilar tudo de uma vez</b></summary>

<br>

```bash
# percorre a raiz e cada submódulo
go build ./... && \
find . -name go.mod -mindepth 2 -execdir go build ./... \;
```

</details>

---

## Roadmap

- [x] Configurar o módulo base e a estrutura de pastas
- [x] Two Pointers — _Container With Most Water_
- [x] Manipulação de strings — _String to Integer_
- [ ] Adicionar testes com `testing` e table-driven tests
- [ ] Benchmarks com `go test -bench`
- [ ] Sliding Window
- [ ] Hash Maps e frequência de caracteres
- [ ] Concorrência: goroutines, channels e `sync`

---

<div align="center">

**[Kaio Rodrigues Braga](https://github.com/KaioRBraga)**

_Aprender em público é a melhor forma de manter o ritmo._

</div>
