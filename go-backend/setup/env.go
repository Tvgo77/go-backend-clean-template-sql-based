package setup

type Env struct {
	TimeoutSeconds int
	DBpassword string
	TokenSecret string
	RunMigration bool
}

func NewEnv() *Env {
	return &Env{
		TimeoutSeconds: 2,
		DBpassword: "0123456789sqa9876543210",
		TokenSecret: "secret",
		RunMigration: false,
	}
}