package setup

type Env struct {
	TimeoutSeconds int
}

func NewEnv() Env {
	return Env{};
}