{
  description = "A flake that sets up the devShell";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    let supportedSystems = [ "aarch64-linux" "i686-linux" "x86_64-linux" ];
    in flake-utils.lib.eachSystem supportedSystems (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ ];
        };
        pp = ps: with ps; [ regex ];
      in {
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [ go (python312.withPackages pp) ];
        };
      });
}
