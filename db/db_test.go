package db

import (
	"chat/models"
	"fmt"
	"testing"
)

func TestCreateMessages(t *testing.T)  {
	err := Connect()
	if err != nil {
		fmt.Println("Connection error: ", err)
		t.Error(err)
	}
	defer Disconnect()
	var message models.Message
	message.From = 1
	message.To = 0
	message.Text = "test10"
	id,err := NewMessage(message)
	fmt.Println(id)
	if err != nil {
		t.Error(err)
	}
	message.From = 0
	message.To = 0
	message.Text = "test00"
	id,err = NewMessage(message)
	fmt.Println(id)
	if err != nil {
		t.Error(err)
	}
	message.From = 1
	message.To = 2
	message.Text = "test12"
	id,err = NewMessage(message)
	fmt.Println(id)
	if err != nil {
		t.Error(err)
	}
	message.From = 2
	message.To = 1
	message.Text = "test21"
	id, err = NewMessage(message)
	fmt.Println(id)
	if err != nil {
		t.Error(err)
	}

	messages := GlobalChat10()
	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("GLOBALCHAT10")
	for _, mes := range messages {
		if mes.ID == 0 {
			t.Error("Не вернул ID-шник")
		}
		// fmt.Println("from"+string(mes.From)+" to"+string(mes.To)+" text:"+string(mes.Text))
	}
	// fmt.Println("")
	// fmt.Println("")
	if len(messages)==0 {
		t.Error("Не вернул сообщения")
	}

	messages = GlobalChatAll()
	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("GLOBALCHATALL")
	for _, mes := range messages {
		if mes.ID == 0 {
			t.Error("Не вернул ID-шник")
		}
		// fmt.Println("from"+string(mes.From)+" to"+string(mes.To)+" text:"+string(mes.Text))
	}
	// fmt.Println("")
	// fmt.Println("")
	if len(messages)==0 {
		t.Error("Не вернул сообщения")
	}

	messages = PrivateChatAll(1, 2)
	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("GLOBALCHATFROM1TO2")
	for _, mes := range messages {
		if mes.ID == 0 {
			t.Error("Не вернул ID-шник")
		}
		// fmt.Println("from"+string(mes.From)+" to"+string(mes.To)+" text:"+string(mes.Text))
	}
	// fmt.Println("")
	// fmt.Println("")
	if len(messages)==0 {
		t.Error("Не вернул сообщения")
	}
	var delID int
	for _, mes := range messages {
		if mes.ID == 0 {
			t.Error("Не вернул ID-шник")
		}
		delID = int(mes.ID)
	}
	lenMes := len(messages)
	DeleteMessage(delID)
	messages = PrivateChatAll(1, 2)
	if len(messages) +1 !=  lenMes{
		t.Error("Не удалил последнюю")
	}

	UpdateMessage(delID-1, "updated")
	
	CreateUser(1, "test_user")
	login, _ := GetLogin(1)
	if login != "test_user" {
		t.Error("Не создал юзера")
	}
	CreateUser(1, "test_user1")
	login, _ = GetLogin(1)
	if login != "test_user1" {
		t.Error("Не апдейтнул юзера")
	}

}

