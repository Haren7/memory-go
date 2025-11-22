package summarizer

import (
	"context"
	"strings"

	textrank "github.com/DavidBelicza/TextRank"
	"github.com/DavidBelicza/TextRank/rank"
)

type ServiceInterface interface {
	Summarize(ctx context.Context, text string) (string, error)
}

type TextRankService struct {
}

func NewTextRankService() ServiceInterface {
	return &TextRankService{}
}

func (r *TextRankService) Summarize(ctx context.Context, text string) (string, error) {
	tr := textrank.NewTextRank()

	// Use default language and rule for parsing
	language := textrank.NewDefaultLanguage()
	rule := textrank.NewDefaultRule()

	// Populate the text
	tr.Populate(text, language, rule)

	// Use default algorithm for ranking
	algorithm := textrank.NewDefaultAlgorithm()
	tr.Ranking(algorithm)

	// Get the top-ranked sentences for the summary
	// rank.ByQty (0) means select by quantity/importance
	// Typically use 3-5 sentences or ~30% of original sentences
	ranks := tr.GetRankData()
	sentences := rank.FindSentences(ranks, rank.ByQty, 5)

	// Extract the sentence text and join them
	var summaryParts []string
	for _, sentence := range sentences {
		summaryParts = append(summaryParts, sentence.Value)
	}

	summary := strings.Join(summaryParts, " ")
	return summary, nil
}

type NoOpService struct {
}

func NewNoOpService() ServiceInterface {
	return &NoOpService{}
}

func (r *NoOpService) Summarize(ctx context.Context, text string) (string, error) {
	return text, nil
}
