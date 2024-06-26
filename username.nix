{ buildGoModule
, fetchFromGitHub
}:

buildGoModule rec {
  pname = "username";
  version = "0.0.6";

  src = fetchFromGitHub {
    owner = "ivankovnatsky";
    repo = "username";
    rev = "v${version}";
    hash = "sha256-Yno1wleUg8mlazhwzF1J9G8sRN/eKqXYBQWG6g5UhGk=";
  };

  vendorHash = null;
}
