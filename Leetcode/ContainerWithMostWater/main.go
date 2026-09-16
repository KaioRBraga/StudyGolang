package main

func maxArea(height []int) int {
	maxAgua := 0
	esquerda := 0
	direita := len(height) - 1

	for esquerda < direita {
		// A altura do container é limitada pela menor barra
		alturaMinima := height[esquerda]
		if height[direita] < alturaMinima {
			alturaMinima = height[direita]
		}

		// A largura é a distância entre os dois ponteiros
		largura := direita - esquerda

		// Calcula a área atual
		areaAtual := alturaMinima * largura

		// Atualiza a maior área encontrada até agora
		switch {
		case areaAtual > maxAgua:
			maxAgua = areaAtual
		}

		switch {
		case height[esquerda] < height[direita]:
			esquerda++
		default:
			direita--
		}
	}

	return maxAgua
}

func main() {

}
