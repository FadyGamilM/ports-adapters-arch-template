package ports

type GrpcPort interface {
	Run()
	PerformDeposite()
	PerformWithdraw()
}
