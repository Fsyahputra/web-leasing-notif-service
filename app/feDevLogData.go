package app

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type FeDevLogNotifier struct {
	Sender Sender
	APIKey string
}

type FeDevLogger struct {
	sender Sender
}

func (fl *FeDevLogger) formatMessage(data FeDevLogDataRaw) string {
	fmtdMsg := fmt.Sprintf(`
	FRONTEND WEB LEASING DEV LOG 💻
	
	Commit Message: %s
	Penulis: %s
	Waktu Commit: %s
	`,
		data.CommitMsg,
		data.Author,
		data.TimeStamp,
	)
	return fmtdMsg

}

func NewFeDevLogger(sender Sender) *FeDevLogger {
	return &FeDevLogger{
		sender: sender,
	}
}

func (fl *FeDevLogger) Handle(data FeDevLogData) error {
	raw := FeDevLogDataRaw{
		CommitMsg: data.CommitMsg,
		Author:    data.Author,
		TimeStamp: data.TimeStamp,
	}
	fmtMsg := fl.formatMessage(raw)
	return fl.sender.Send(fmtMsg)
}

func (dh *FeDevLogNotifier) generateSummary(diff string) string {
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

func (dh *FeDevLogNotifier) formatMessage(data FeDevLogDataProcessed) string {
	fmtdMsg := fmt.Sprintf(`
	PENGEMBANGAN FRONTEND WEB LEASING DEV 💻
	
	Commit Message: %s
	Penulis: %s
	Wjktu Commit: %s
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

func (dh *FeDevLogNotifier) Handle(data FeDevLogData) error {
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

func NewFeDevLogNotifier(sender Sender, apiKey string) *FeDevLogNotifier {
	return &FeDevLogNotifier{
		Sender: sender,
		APIKey: apiKey,
	}

}
