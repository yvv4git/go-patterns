package builder

import "fmt"

func sendEmail(message Message) string {
	return fmt.Sprintf(
		"Processing email send ...\nTo:%s,\nSubject:%s\n%s\nRegards,\n%s\n",
		message.to,
		message.subject,
		message.body,
		message.from,
	)
}
