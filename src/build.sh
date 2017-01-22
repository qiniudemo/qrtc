DIR=$(cd ../; pwd)
export GOPATH=$GOPATH:$DIR
go build -o ../qrtc/qrtc main.go
