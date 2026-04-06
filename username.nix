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
    hash = "sha256-zhW/7Wo9NTH6VTY0XumlSl0W762yznsybvGtZ2hJfOg=";
  };

  vendorHash = null;
}
