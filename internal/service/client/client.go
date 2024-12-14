package client

type service struct {
	cfg            Config
	repos          repository
	sessionStorage sessionStorage
}

func New(
	cfg Config,
	repos repository,
) *service {
	return &service{
		cfg:   cfg,
		repos: repos,
	}
}
