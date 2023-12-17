module github.com/multiverse-os/banner

go 1.19

replace (
	github.com/multiverse-os/banner/fonts/big v0.1.0 => ./fonts/big
	github.com/multiverse-os/banner/fonts/giant v0.1.0 => ./fonts/giant
)

require (
	github.com/multiverse-os/banner/fonts/big v0.1.0
	github.com/multiverse-os/banner/fonts/giant v0.1.0
)
