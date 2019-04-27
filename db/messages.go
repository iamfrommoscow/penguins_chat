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
)
RETURNING id`

const authorizedToGlobalChat = `
VALUES (
	$1,
	0,
	$2
)
RETURNING id`

const userToUser = `
VALUES (
	$1,
	$2,
	$3
)
RETURNING id`


func NewMessage(message models.Message) (int, error) {
	query := insertIntoMessages
	var id int
	if message.To == 0{
		if message.From == 0 {
			query += unauthorizedToGlobalChat
			err := connection.QueryRow(query, message.Text).Scan(&id)
			if err != nil {
				fmt.Println(err)
				return -1, err
			}
			return id,nil
		}
		query += authorizedToGlobalChat
		err := connection.QueryRow(query, message.From, message.Text).Scan(&id)
		if err != nil {
			fmt.Println(err)
			return -1,err
		}
		return id, nil
	}
	query += userToUser
	err := connection.QueryRow(query, message.From, message.To, message.Text).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return -1,err
	}
	return id,nil
}

const select10GlobalMessages= `
SELECT id, fromUser, message
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
SELECT id, fromUser, message
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
SELECT id, fromUser, message
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

const deleteMessageByID = `
DELETE FROM messages
WHERE id = $1;`

func DeleteMessage(id int) error {
	_, err := Exec(deleteMessageByID, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

const updateMessageByID = `
UPDATE messages
SET message = $2
WHERE id = $1;`

func UpdateMessage(id int, updatedMessage string) error {

	_, err := Exec(updateMessageByID, id, updatedMessage)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
