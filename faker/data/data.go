package data

// Data consists of the main set of fake information
var Data = map[string]map[string][]string{
	"person": Person,
	//	"contact":  Contact,
	//	"address":  Address,
	//	"company":  Company,
	//	"lorem":    Lorem,
	"internet": Internet,
	//	"file":     Files,
	"job":      Jobs,
	"color":    Colors,
	"computer": Computer,
	"gender":   Genders,
	"areacode": AreaCodes,
	"phone":    Phones,
	"email":    Emails,
	//	"payment":  Payment,
	//	"hipster":  Hipster,
	//	"beer":     Beer,
	//	"hacker":   Hacker,
	//	"currency": Currency,
}

// IntData consists of the main set of fake information (integer only)
var IntData = map[string]map[string][]int{
	"status_code": StatusCodes,
}
