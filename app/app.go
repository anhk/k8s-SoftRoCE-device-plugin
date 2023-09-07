package app

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() error {
	return nil
}
