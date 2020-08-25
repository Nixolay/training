
.PHONY: install-pre-commit
install-pre-commit:
	python3 -m pip install --user --upgrade pip ; python3 -m pip install --user pre-commit gitlint
	pre-commit install -f ; pre-commit install --hook-type commit-msg
