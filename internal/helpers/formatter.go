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

func DoneFormat(done bool) string {
	if done {
		return "\033[32m✓\033[0m"
	}
	return "\033[31m×\033[0m"
}
