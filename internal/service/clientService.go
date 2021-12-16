package service

func (s *service) SendMessage(chatId int64, msg string) {
	s.tg.SendMessage(chatId, msg)
}

func (s *service) TGMenu(chatId int64, msg string) {
	s.tg.Menu(chatId, msg)
}
