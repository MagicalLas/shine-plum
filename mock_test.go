package shine_plum_test

// UseCase is us
type UseCase struct{}

type Creative struct {
	Image string
	Icon  string
}

type AD struct {
	Title       string
	Description string
	Creatives   []*Creative
}

// GetAd is get ad. las
//aka wonho
func (u *UseCase) GetAd(lineitemID string, adnID int) *AD {
	return &AD{
		Title:       "tt",
		Description: "dd",
		Creatives: []*Creative{
			{
				Image: "http://magical.dev",
			},
			{
				Icon: "http://magical.dev",
			},
		},
	}
}

type User struct {
	IFA string
	IP  string
}

type ListAdCommand struct {
	LineitemID string
	ADNID      int
	UnitID     int
	User       User
}

func (u *UseCase) ListAd(command *ListAdCommand) []*AD {
	return []*AD{
		u.GetAd(command.LineitemID, command.ADNID),
	}
}

