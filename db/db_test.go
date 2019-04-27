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
	err = NewMessage(message)
	if err != nil {
		t.Error(err)
	}
	message.From = 0
	message.To = 0
	message.Text = "test00"
	err = NewMessage(message)
	if err != nil {
		t.Error(err)
	}
	message.From = 1
	message.To = 2
	message.Text = "test12"
	err = NewMessage(message)
	if err != nil {
		t.Error(err)
	}
	message.From = 2
	message.To = 1
	message.Text = "test21"
	err = NewMessage(message)
	if err != nil {
		t.Error(err)
	}

	messages := GlobalChat10()
	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("GLOBALCHAT10")
	// for _, mes := range messages {
	// 	fmt.Println(mes.From)
	// 	fmt.Println("from"+string(mes.From)+" to"+string(mes.To)+" text:"+string(mes.Text))
	// }
	// fmt.Println("")
	// fmt.Println("")
	if len(messages)==0 {
		t.Error("Не вернул сообщения")
	}

	messages = GlobalChatAll()
	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("GLOBALCHATALL")
	// for _, mes := range messages {
	// 	fmt.Println("from"+string(mes.From)+" to"+string(mes.To)+" text:"+string(mes.Text))
	// }
	// fmt.Println("")
	// fmt.Println("")
	if len(messages)==0 {
		t.Error("Не вернул сообщения")
	}

	messages = PrivateChatAll(1, 2)
	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("GLOBALCHATFROM1TO2")
	// for _, mes := range messages {
	// 	fmt.Println(mes.From)
	// 	fmt.Println("from"+string(mes.From)+" to"+string(mes.To)+" text:"+string(mes.Text))
	// }
	// fmt.Println("")
	// fmt.Println("")
	if len(messages)==0 {
		t.Error("Не вернул сообщения")
	}



}

