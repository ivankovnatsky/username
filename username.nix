{ buildGoModule
, fetchFromGitHub
, scowl
, makeWrapper
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

  nativeBuildInputs = [ makeWrapper ];

  postInstall = ''
    wrapProgram $out/bin/username \
      --set WORD_FILE ${scowl}/share/dict/wamerican.50
  '';
}
