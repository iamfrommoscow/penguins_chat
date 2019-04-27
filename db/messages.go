package db

import 
(
	"chat/models"
	"fmt"
)

const insertIntoMessages = `
INSERT INTO messages (fromUser, toUser, message)`

const unauthorizedToGlobalChat = `
VALUES (
	0,
	0,
	$1
);`

const authorizedToGlobalChat = `
VALUES (
	$1,
	0,
	$2
);`

const userToUser = `
VALUES (
	$1,
	$2,
	$3
);`

func NewMessage(message models.Message) error {
	query := insertIntoMessages
	if message.To == 0{
		if message.From == 0 {
			query += unauthorizedToGlobalChat
			_, err := Exec(query, message.Text)
			if err != nil {
				fmt.Println(err)
				return err
			}
			return nil
		}
		query += authorizedToGlobalChat
		_, err := Exec(query, message.From, message.Text)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
	query += userToUser
	_, err := Exec(query, message.From, message.To, message.Text)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

const select10GlobalMessages= `
SELECT fromUser, message
FROM messages
WHERE toUser = 0
ORDER BY id DESC LIMIT 10;
`

func GlobalChat10() []models.Message  {
	var messages []models.Message
	rows, err := Query(select10GlobalMessages)
	if err != nil {
		fmt.Println(err)
		return messages
	}
	defer rows.Close()
	messages = RowsToMessages(rows)
	return messages
}

const selectGlobalMessages = `
SELECT fromUser, message
FROM messages
WHERE toUser = 0
ORDER BY id DESC;
`

func GlobalChatAll() []models.Message {
	var messages []models.Message
	rows, err := Query(selectGlobalMessages)
	if err != nil {
		fmt.Println(err)
		return messages
	}
	defer rows.Close()
	messages = RowsToMessages(rows)
	return messages
}

const privateChatMessages = `
SELECT fromUser, message
FROM messages
WHERE (
  fromUser = $1 AND toUser = $2
  ) OR (
  toUser = $1 AND fromUser = $2
  )
ORDER BY id DESC;
`

func PrivateChatAll(from int, to int) []models.Message {
	var messages []models.Message
	rows, err := Query(privateChatMessages, from, to)
	if err != nil {
		fmt.Println(err)
		return messages
	}
	defer rows.Close()
	messages = RowsToMessages(rows)
	return messages
}
