package middleware

var processingStatus = make(map[string]string)

func (m *middleware) CreateMenu() Menu {
	return NewMenu(m)
}

func (m *middleware) Begin(msg string, id int64, menu Menu) {
	menu.Main(msg, id)
}

func (m *middleware) Request(msg string, id int64, menu Menu) {
	switch processingStatus[StatusOn] {
	case m.strCfg.Status.ProcessingStatus.Organization:
		menu.Organization(msg, id)
	case m.strCfg.Status.ProcessingStatus.Address:
		menu.Address(msg, id)
	case m.strCfg.Status.ProcessingStatus.Name:
		menu.Name(msg, id)
	case m.strCfg.Status.ProcessingStatus.Intermediate:
		menu.Intermediate(msg, id)
	case m.strCfg.Menu.Button.Back:
		menu.Back(id)
	case m.strCfg.Status.ProcessingStatus.View:
		menu.View(msg, id)
	case m.strCfg.Status.ProcessingStatus.Overview:
		menu.Overview(msg, id)
	case m.strCfg.Status.ProcessingStatus.Region:
		menu.Region(msg, id)
	case m.strCfg.Status.ProcessingStatus.Code:
		menu.Code(msg, id)
	}
}
