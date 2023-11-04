camera:
	go run ./demo/showvideo/main.go 1

cmdlife:
	go run ./demo/cmdlife -gens=3 -pop=40 -height=10 -width=10

badgelife:
	tinygo flash -size short -target=gobadge ./demo/badgelife/

cubetest:
	tinygo flash -size short -target=itsybitsy-m4 -opt=2 ./demo/cubetest/
