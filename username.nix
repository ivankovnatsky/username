{ buildGoModule
}:

buildGoModule {
  pname = "username";
  version = "0.0.7";

  src = ./.;

  vendorHash = null;
}
