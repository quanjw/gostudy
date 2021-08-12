module gostudy

go 1.16

require rsc.io/quote v1.5.2

require (
	github.com/myuser/calculator v0.0.0
	github.com/rs/zerolog v1.23.0 // indirect
)

replace github.com/myuser/calculator => ../calculator
