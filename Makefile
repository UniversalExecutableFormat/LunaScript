.PHONY: all compiler runtime

all:
	@python3 scripts/make_all.py

compiler:
	@python3 scripts/make_compiler.py

runtime:
	@python3 scripts/make_runtime.py