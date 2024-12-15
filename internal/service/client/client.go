package client

type service struct {
	cfg            Config
	repos          repository
	mailer         mailer
	sessionStorage sessionStorage
	oauth          oauthClient
}

func New(
	cfg Config,
	repos repository,
	storage sessionStorage,
	mailer mailer,
	oauth oauthClient,
) *service {
	return &service{
		cfg:            cfg,
		repos:          repos,
		oauth:          oauth,
		mailer:         mailer,
		sessionStorage: storage,
	}
}
