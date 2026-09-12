module example

go 1.26

require (
	solod.dev v0.4.0
	solod.dev/sodium v0.0.0-00010101000000-000000000000
)

replace solod.dev/sodium => ..
