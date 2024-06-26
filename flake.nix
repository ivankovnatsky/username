{
  description = "Generate random meaningful username.";

  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system}; in
      rec {
        packages = flake-utils.lib.flattenTree {
          username = pkgs.callPackage ./username.nix { };
        };
        defaultPackage = packages.username;
        devShell = pkgs.mkShell {
          buildInputs = [ packages.username ];
        };
      }
    );
}
