

```bash
(git diff; time2 go test -v -bench=. -run=Bench -benchtime=1ns) |& tee -a results-192-cores.txt;
(git diff; time2 go test -v -bench=. -run=Bench -benchtime=9x) |& tee -a results-192-cores.txt;
(git diff; time2 go test -v -bench=. -run=Bench) |& tee -a results-192-cores.txt;
```
