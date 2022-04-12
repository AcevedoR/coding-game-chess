run perfs: `go test -bench=. -cpuprofile cpu.prof ./ && go tool pprof -http :9402 cpu.prof`
test only one case: `go test -timeout 30s -run ^TestIllegalMoves/Illegal_move_3$ main/gochess`
