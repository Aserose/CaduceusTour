package middleware

type Menu interface {
	Main(msg string, id int64)
	Region(msg string, id int64)
	Organization(msg string, id int64)
	Address(msg string, id int64)
	Name(msg string, id int64)
	Intermediate(msg string, id int64)
	View(msg string, id int64)
	Overview(msg string, id int64)
	Comparing(id int64)
	Code(msg string, id int64)
	Back(id int64)
}

const (
	AllRegions = "РФ"
	StatusOn   = "on"
	compareOn  = "compare"
)

type menu struct {
	middleware *middleware
}

func NewMenu(middleware *middleware) Menu {
	return &menu{
		middleware: middleware,
	}
}

func (m menu) Main(msg string, id int64) {
	switch msg {
	case "/start":
		m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.Menu)

	case m.middleware.strCfg.Menu.Button.Region:
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Region
		m.middleware.service.TGMenu(id, msg)

	case m.middleware.strCfg.Menu.Button.Organization:
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Organization
		m.middleware.service.TGMenu(id, msg)

	case m.middleware.strCfg.Menu.Button.Back:
		m.Back(id)
	}
}

func (m menu) Region(msg string, id int64) {
	switch msg {
	case AllRegions:
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.View
		m.middleware.getDataAboutRegion("", id)
		m.middleware.service.Next(id)
		m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.View)

	case m.middleware.strCfg.Menu.Button.Code:
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Code
		m.middleware.service.SendMessage(id, m.middleware.strCfg.Response.MsgToUser.EnterRegion)

	case m.middleware.strCfg.Menu.Button.Back:
		m.Back(id)
	}
}
func (m menu) Organization(msg string, id int64) {
	switch msg {
	case m.middleware.strCfg.Menu.Button.Back:
		m.Back(id)

	case m.middleware.strCfg.Menu.Button.Name:
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Name
		m.middleware.service.SendMessage(id, m.middleware.strCfg.Response.MsgToUser.EnterName)
	case m.middleware.strCfg.Menu.Button.Address:

		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Address
		m.middleware.service.SendMessage(id, m.middleware.strCfg.Response.MsgToUser.EnterAddress)
	}
}

func (m menu) Address(msg string, id int64) {
	switch msg {
	case m.middleware.strCfg.Menu.Button.Back:
		m.Back(id)

	case m.middleware.strCfg.Menu.Button.Name:
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Main
		m.Organization(msg, id)

	case m.middleware.strCfg.Menu.Button.Address:
		m.Organization(msg, id)

	default:
		if m.middleware.getDataAboutOrganization(m.middleware.getOrganizationNameByAddress(msg), id) == "not found" {
			processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Organization
			m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.Organization)
		} else {
			m.AfterFoundOrganization(id)
		}
	}
}

func (m menu) Name(msg string, id int64) {
	switch msg {
	case m.middleware.strCfg.Menu.Button.Back:
		m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.Menu)

	case m.middleware.strCfg.Menu.Button.Address:
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Main
		m.Organization(msg, id)

	case m.middleware.strCfg.Menu.Button.Name:
		m.Organization(msg, id)

	default:
		if m.middleware.getDataAboutOrganization(msg, id) == m.middleware.strCfg.Status.StatusSearch.None {
			processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Organization
			m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.Organization)
		} else {
			m.AfterFoundOrganization(id)
		}
	}
}

func (m menu) AfterFoundOrganization(id int64) {
	switch processingStatus[compareOn] {
	case " ":
		m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.Intermediate)
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Intermediate

	case "":
		m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.Intermediate)
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Intermediate

	case m.middleware.strCfg.Status.ProcessingStatus.Organization:
		m.Comparing(id)
	}
}

func (m menu) Intermediate(msg string, id int64) {
	switch msg {
	case m.middleware.strCfg.Menu.Button.Overview:
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Overview
		m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.Overview)
		m.middleware.service.General(id)

	case m.middleware.strCfg.Menu.Button.List:
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.View
		m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.View)
		m.middleware.service.Next(id)
	}
}

func (m menu) View(msg string, id int64) {
	switch msg {
	case m.middleware.strCfg.Menu.Button.NextRow:
		m.middleware.service.NextRow(id)

	case m.middleware.strCfg.Menu.Button.Next:
		m.middleware.service.Next(id)

	case m.middleware.strCfg.Menu.Button.Load:
		m.middleware.service.CreateDocument(id)

	case m.middleware.strCfg.Menu.Button.Overview:
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Overview
		m.middleware.service.ToOverview()
		m.middleware.service.General(id)
		m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.Overview)

	case m.middleware.strCfg.Menu.Button.Back:
		m.Back(id)
	}
}

func (m menu) Overview(msg string, id int64) {
	switch msg {
	case m.middleware.strCfg.Menu.Button.Compare:
		processingStatus[compareOn] = m.middleware.strCfg.Status.ProcessingStatus.Organization
		m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.Organization)
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Organization
		m.Organization(msg, id)
	case m.middleware.strCfg.Menu.Button.Back:
		m.Back(id)

	case m.middleware.strCfg.Menu.Button.List:
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.View
		m.middleware.service.ToView()
		m.middleware.service.Next(id)
		m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.View)
	}
}

func (m menu) Comparing(id int64) {
	m.middleware.service.Compare(id)
	m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.Overview)
	processingStatus[compareOn] = ""
	processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Overview
}

func (m menu) Code(msg string, id int64) {
	switch msg {
	case m.middleware.strCfg.Menu.Button.Back:
		m.Back(id)

	case AllRegions:
		m.Region(msg, id)

	case m.middleware.strCfg.Menu.Button.Code:
		m.Region(msg, id)

	default:
		if m.middleware.getDataAboutRegion(msg, id) == m.middleware.strCfg.Status.StatusSearch.None {
			break
		}
		processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.View
		m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.View)
		m.middleware.service.Next(id)
	}
}

func (m menu) Back(id int64) {
	processingStatus[StatusOn] = m.middleware.strCfg.Status.ProcessingStatus.Main
	processingStatus[compareOn] = " "
	m.middleware.service.TGMenu(id, m.middleware.strCfg.Menu.Button.Menu)
}
