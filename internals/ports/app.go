// contains the ports for the application layer
package ports

type ApiPort interface {
	PerformDeposite(float32) (float32, error)
	PerformWithdraw(float32) (float32, error)
}
