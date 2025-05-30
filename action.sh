#!/usr/bin/env bash
set -Eeuo pipefail

# Detect platform
if [[ "$(uname -s)" != "Linux" ]]; then
  echo "This action currently supports Ubuntu-linux runners only."; exit 1
fi

# Install terrafetch
echo "::group::Install terrafetch"
ver="${INPUT_TERRAFETCH_VERSION}"
if [[ "$ver" == "latest" ]]; then ver=""; fi
curl -1sLf https://dl.cloudsmith.io/public/rosesecurity/terrafetch/setup.deb.sh \
 | sudo -E bash
sudo apt-get -qq update
sudo apt-get -qq install -y "terrafetch${ver:+=$ver}"
echo "::endgroup::"

# Run terrafetch
echo "::group::Run terrafetch"
set +e
TERRAFETCH_OUTPUT="$(terrafetch -d "$INPUT_TERRAFORM_DIRECTORY")"
ret=$?
set -e
echo "terrafetch finished with exit code $ret"
echo "terrafetch-return-code=$ret" >>"$GITHUB_OUTPUT"
[[ $ret -ne 0 ]] && exit "$ret"        # fail the step/job on error
echo "::endgroup::"

# Inject output into README
outfile="$INPUT_OUTPUT_FILE"
tmp=$(mktemp)
awk -v block="$TERRAFETCH_OUTPUT" '
  BEGIN         {inside=0}
  /<!-- TERRAFETCH:START -->/ {print; print block; inside=1; next}
  /<!-- TERRAFETCH:END -->/   {inside=0}
  !inside        {print}
' "$outfile" >"$tmp"
mv "$tmp" "$outfile"

# Commit when file changes
echo "::group::Commit & push (if needed)"
git config --global user.email "41898282+github-actions[bot]@users.noreply.github.com"
git config --global user.name  "github-actions[bot]"

if git diff --quiet "$outfile"; then
  echo "No changes detected – skipping commit."
else
  git add "$outfile"
  git commit -m "chore: update Terrafetch section"
  git push "https://${INPUT_GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git" HEAD:$(git rev-parse --abbrev-ref HEAD)
fi
echo "::endgroup::"
