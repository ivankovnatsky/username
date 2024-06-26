{ buildGoModule
, fetchFromGitHub
}:

buildGoModule rec {
  pname = "username";
  version = "0.0.1";

  src = fetchFromGitHub {
    owner = "ivankovnatsky";
    repo = "username";
    rev = "v${version}";
    hash = "sha256-LEOQrsTTjJwlQP2BgtaQ9n0c82YBHiBGn/sfB41iMmY=";
  };

  vendorHash = null;
}
