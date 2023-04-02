package setting

type Settings struct {
	Themes
}

func New() Settings {
	return Settings{
		Themes: Theme(),
	}
}
