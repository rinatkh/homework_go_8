package analysis

// Cost разделяет дополнительную память операции и суммарную память структуры.
// Время обозначайте O(1), O(n), amortized O(1), expected O(1) по docs/07_analysis.md.
type Cost struct{ Time, Extra, Total string }

func Analyze(operation string, knownHandle bool) Cost { // TODO: 07.1.
	return Cost{}
}

// SelectStructure выбирает каноническую метку по ЯВНЫМ ограничениям сценария.
// Это учебные контракты, а не утверждение об одном лучшем решении для всех задач.
func SelectStructure(scenario string) string { // TODO: 07.2.
	return "TODO"
}
