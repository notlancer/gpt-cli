package bootstrap

type App struct {
	Env *Env
}

func CreateApp() App {
	app := &App{}
	app.Env = NewEnv()

	return *app
}
