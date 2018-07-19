## Notes


The rpc package could return a pointer, which would make the code look cleaner when there is an error as
we could return nil, instead of struct{}