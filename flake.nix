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
            version = "0.2.0";

            src = ./.;

            subPackages = [ "cmd/md2puki" ];

            vendorHash = "sha256-P+itOTHU9mtHxZ5SI/Z0kH/yCN/bVBBNMt8UKYT7WDQ=";

            meta = {
              description = "Markdown to Pukiwiki notation converter";
              homepage = "https://github.com/Nanamiiiii/md2puki";
              mainProgram = "md2puki";
            };
          });
        };

        apps.md2puki = {
          type = "app";
          program = "${self.packages.${system}.md2puki}/bin/md2puki";
        };

        formatter = treefmt-nix.lib.mkWrapper pkgs {
          projectRootFile = "flake.nix";
          programs.nixfmt.enable = true;
          programs.gofmt.enable = true;
        };
      }
    );
}
