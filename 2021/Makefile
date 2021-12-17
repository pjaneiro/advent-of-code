.DEFAULT_GOAL := all

all: FORCE
	@ echo "Golang" && cd golang && go run main.go
	@ echo
	@ echo "Haskell" && cd haskell && runhaskell main.hs
	@ echo
	@ echo "Lisp" && cd lisp && clisp main.lisp

go: FORCE
	@ cd golang && go run main.go

haskell: FORCE
	@ cd haskell && runhaskell main.hs

lisp: FORCE
	@ cd lisp && clisp main.lisp

FORCE:
