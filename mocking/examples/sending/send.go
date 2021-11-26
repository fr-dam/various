//go:generate mockery --name=SendFunc
package sending

type SendFunc func(data string) (int, error)