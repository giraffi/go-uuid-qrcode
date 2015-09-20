gom exec gox -output "dist/qruuid_{{.OS}}_{{.Arch}}/qruuid"
find dist/* -type d -exec zip -r {}.zip {} \;
# ghr --username giraffi --token $GITHUB_TOKEN --replace --prerelease localtest dist/

