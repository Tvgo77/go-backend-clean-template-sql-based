package setup

type Env struct {
	TimeoutSeconds int
	DBpassword string
	TokenSecret string
	RunMigration bool
	TestMode bool
}

func NewEnv() *Env {
	return &Env{
		TimeoutSeconds: 2,
		DBpassword: "postgres",
		TokenSecret: "secret",
		RunMigration: true,
		TestMode: true,
	}
}