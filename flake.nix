{
  description = "Convert markdown to pukiwiki notation";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    flake-compat.url = "github:edolstra/flake-compat";
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  nixConfig = {
    extra-substituters = [
      "https://nix-cache.myuu.dev/packages"
      "https://nix-community.cachix.org"
      "https://cache.numtide.com"
    ];

    extra-trusted-public-keys = [
      "nix-cache.myuu.dev-1:2lAuxMiua4hEYRgGu3JXHafpZrHprrQi+TmrQIKJ6+E="
      "nix-community.cachix.org-1:mB9FSh9qf2dCimDSUo8Zy7bkq5CX+/rkCWyvRCYg3Fs="
      "niks3.numtide.com-1:DTx8wZduET09hRmMtKdQDxNNthLQETkc/yaX7M4qK0g="
    ];
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      treefmt-nix,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages = {
          md2puki = pkgs.buildGoModule (finalAttrs: {
            pname = "md2puki";
            version = "0.3.3";

            src = ./.;

            subPackages = [ "cmd/md2puki" ];

            vendorHash = "sha256-j2XenbE5d8JlJW3eRrFUz4arYYQBtdGIr0sHyKi37a4=";

            meta = {
              description = "Markdown to Pukiwiki notation converter";
              homepage = "https://github.com/Nanamiiiii/md2puki";
              mainProgram = "md2puki";
            };
          });

          default = self.packages.${system}.md2puki;
        };

        apps = {
          md2puki = {
            type = "app";
            program = "${self.packages.${system}.md2puki}/bin/md2puki";
          };

          default = self.apps.${system}.md2puki;
        };

        devShells = {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
            ];
          };
        };

        formatter = treefmt-nix.lib.mkWrapper pkgs {
          projectRootFile = "flake.nix";
          programs.nixfmt.enable = true;
          programs.gofmt.enable = true;
        };
      }
    );
}
