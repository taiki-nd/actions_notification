package core

const (
	EnviromentLocal = "local"
	EnviromentDev   = "dev"
	EnviromentStg   = "stg"
	EnviromentCiCd  = "cicd"
	EnviromentProd  = "prod"
)

type Enviroment struct {
	Enviroment string
}

func NewEnviroment(env string) *Enviroment {
	return &Enviroment{
		Enviroment: env,
	}
}

func (e *Enviroment) IsLocal() bool {
	return e.Enviroment == EnviromentLocal
}

func (e *Enviroment) IsDev() bool {
	return e.Enviroment == EnviromentDev
}

func (e *Enviroment) IsStg() bool {
	return e.Enviroment == EnviromentStg
}

func (e *Enviroment) IsCiCd() bool {
	return e.Enviroment == EnviromentCiCd
}

func (e *Enviroment) IsProd() bool {
	return e.Enviroment == EnviromentProd
}
