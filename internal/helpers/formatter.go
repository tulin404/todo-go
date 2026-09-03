package helpers

func SplitRigid(s string, chunkMax int) []string {
	var chunks []string
	runes := []rune(s) // Convertido para runes para evitar quebrar caracteres multibyte ao meio

	for len(runes) > 0 {
		if len(runes) < chunkMax {
			chunkMax = len(runes)
		}
		chunks = append(chunks, string(runes[:chunkMax]))
		runes = runes[chunkMax:]
	}
	return chunks
}
