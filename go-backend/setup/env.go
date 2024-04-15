package setup

type Env struct {
	TimeoutSeconds int
	TokenSecret string
	RunMigration bool
}

func NewEnv() Env {
	return Env{};
}