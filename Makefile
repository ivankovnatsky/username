default:
	go build -o username .

update-nix-config:
	@set -euo pipefail; \
	branch=$$(git rev-parse --abbrev-ref HEAD); \
	if [ "$$branch" = "HEAD" ]; then \
		echo "Error: detached HEAD state, must be on main branch"; exit 1; \
	elif [ "$$branch" != "main" ]; then \
		echo "Error: must be on main branch (currently on $$branch)"; exit 1; \
	fi; \
	git push; \
	cd "$$(ghq root)/github.com/ivankovnatsky/nix-config"; \
	NIX_CONFIG="access-tokens = github.com=$$(gh auth token)" nix flake update username --commit-lock-file
