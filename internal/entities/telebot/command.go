package telebot

type Command struct {
	Name        string
	Description string
	Action      func(input interface{}) string
}
