package lib

import (
	"CHGPU_T_BOT/lib/back"
	"CHGPU_T_BOT/lib/domain"
	"CHGPU_T_BOT/lib/educationlevels"
	"CHGPU_T_BOT/lib/mainmenu"
	"CHGPU_T_BOT/lib/specialist"
)

var MessageMap = map[string]domain.MessageHandler{
	domain.ApplicationTypes: mainmenu.SendApplicationTypesMessage,
	domain.ApplicationDocs:  mainmenu.SendApplicationDocsMessage,
	domain.Contacts:         mainmenu.SendContactsMessage,
	domain.EducationLevels:  mainmenu.SendEducationLevelsMessage,
	domain.BachelorMenu:     educationlevels.SendBachelorMessage,
	domain.MastersMenu:      educationlevels.SendMastersMessage,
	domain.PostGradMenu:     educationlevels.SendPostGradMessage,
	domain.SecondaryMenu:    educationlevels.SendSecondaryMessage,
	domain.EntranceRule:     educationlevels.EntranceRules,
	domain.MinPoints:        educationlevels.MinimalPoints,
	domain.List:             educationlevels.List,
	domain.IndividualList:   educationlevels.IndividualList,
	domain.Exception:        educationlevels.Exception,
	domain.Periods:          educationlevels.Periods,
	domain.BackMessage:      educationlevels.HandleBackMessage,
	domain.BackToLevels:     back.HandleBackMessage,
	domain.AskSpecialist:    specialist.SendSpecialistMessage,
}
