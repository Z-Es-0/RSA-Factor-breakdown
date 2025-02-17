package factorbreakdown

import (
	"context"
	"math/big"
	"sync"
)

type interval struct {
	begin big.Int
	end   big.Int
}

func findfactor(ctx context.Context, xp interval, n big.Int) big.Int {
	i := new(big.Int).Set(&xp.begin)
	two := big.NewInt(2)
	if new(big.Int).Mod(i, two).Cmp(big.NewInt(0)) == 0 {
		i.Add(i, big.NewInt(1))
	}
	for i.Cmp(&xp.end) <= 0 {

		select {
		case <-ctx.Done():
			return *big.NewInt(0)
		default:
		}

		if new(big.Int).Mod(&n, i).Cmp(big.NewInt(0)) == 0 {
			println("Factor found:", i.String())
			return *i
		}

		i.Add(i, two)
	}
	return *big.NewInt(0)
}

func worker(ctx context.Context, ports chan interval, wg *sync.WaitGroup, n big.Int, ans chan big.Int, cancel context.CancelFunc) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case p, ok := <-ports:
			if !ok {
				return
			}
			factor := findfactor(ctx, p, n)

			if factor.Cmp(big.NewInt(0)) != 0 {

				cancel()

				select {

				case ans <- factor:
				default:
				}
				return
			}
		}
	}
}

/*
BuildFactory 启动 WorkerSize 个协程，并将任务分配到 ports 通道中。

WorkerSize 为并发执行的 goroutine 数量，
n 为待分解的数。
*/
func BuildFactory(WorkerSize int, n big.Int) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	ports := make(chan interval, WorkerSize)
	ans := make(chan big.Int, 1)

	for i := 0; i < WorkerSize; i++ {
		wg.Add(1)
		go worker(ctx, ports, &wg, n, ans, cancel)
	}

	span := new(big.Int).SetInt64(1000000000)
	sqrt := new(big.Int).Sqrt(&n)
	right := new(big.Int).SetInt64(2)

	// 分配任务

	for i := new(big.Int).SetInt64(2); i.Cmp(sqrt) < 0; i.Add(i, span) {

		select {
		case <-ctx.Done():
			goto DONE
		default:
		}

		right.Add(right, span)
		if right.Cmp(sqrt) > 0 {
			right.Set(sqrt)
		}

		ports <- interval{
			begin: *new(big.Int).Set(i),
			end:   *new(big.Int).Set(right),
		}
	}
DONE:

	close(ports)

	wg.Wait()

	select {
	case factor := <-ans:
		q := new(big.Int).Div(&n, &factor)
		println("p and q are", factor.String(), "and", q.String())
	default:
		println("No factor found")
	}
}
