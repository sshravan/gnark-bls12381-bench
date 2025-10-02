

```bash
time go test -v -bench=. -run=Bench -benchtime=9x -timeout 30m # Performs TEN, not NINE, runs
time go test -v -bench=. -run=Bench -benchtime=1ns -timeout 30m # Performs EXACTLY single run
```
