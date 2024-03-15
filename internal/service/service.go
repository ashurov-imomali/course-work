package service

type Service struct {
	Repo ReposMeths
}

func GetService(repo ReposMeths) SrvMeths {
	return Service{Repo: repo}
}
