for x in **/*.proto;
do
  echo "start gen proto ,path:$x"
  protoc --proto_path=. --go_out=paths=source_relative:. "$x";
done
