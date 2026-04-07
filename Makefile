default:
	go build -o username .

nix-update:
	nix flake update username --flake $$(ghq list -p nix-config) --commit-lock-file
