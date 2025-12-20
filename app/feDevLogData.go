package app

import (
	"context"
	"fmt"
	"google.golang.org/genai"
)

type FeDevLogDataHandler struct {
	Sender Sender
	APIKey string
}

func (dh *FeDevLogDataHandler) generateSummary(diff string) string {
	if diff == "" {
		return "No diff Provided"
	}
	promptCtx := "ringkas kode git diff berikut ini supaya bos saya yang orang awam bisa mengerti perubahan apa yang saya buat singkat saja 1 - 2 kalimat:\n\n"
	prompt := fmt.Sprintf("%s%s", promptCtx, diff)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: dh.APIKey,
	})
	if err != nil {
		return "Error Creating AI Client"
	}
	resp, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash", genai.Text(prompt), nil)
	if err != nil {
		return "Error Generating Summary"
	}
	summary := resp.Text()
	return summary
}

func (dh *FeDevLogDataHandler) formatMessage(data FeDevLogDataProcessed) string {
	fmtdMsg := fmt.Sprintf(`
PENGEMBANGAN FRONTEND LEASING DEV 💻

Commit Message: %s
Penulis: %s
Waktu Commit: %s
Ringkasan Perubahan:
%s
	`,
		data.CommitMsg,
		data.Author,
		data.TimeStamp,
		data.Summary,
	)
	return fmtdMsg
}

func (dh *FeDevLogDataHandler) Handle(data FeDevLogData) error {
	summary := dh.generateSummary(data.Diff)
	processedData := FeDevLogDataProcessed{
		CommitMsg: data.CommitMsg,
		Author:    data.Author,
		TimeStamp: data.TimeStamp,
		Summary:   summary,
	}
	msg := dh.formatMessage(processedData)
	return dh.Sender.Send(msg)
}

func NewFeDevLogDataHandler(sender Sender, apiKey string) *FeDevLogDataHandler {
	return &FeDevLogDataHandler{
		Sender: sender,
		APIKey: apiKey,
	}

}
