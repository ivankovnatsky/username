{ buildGoModule
, fetchFromGitHub
}:

buildGoModule rec {
  pname = "username";
  version = "0.0.7";

  src = fetchFromGitHub {
    owner = "ivankovnatsky";
    repo = "username";
    rev = "v${version}";
    hash = "";
  };

  vendorHash = null;
}
