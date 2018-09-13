cd $GOPATH/src/github.com/lvxin1986/golang-test/mock
mockgen -destination=mocks/mock_doer.go -package=mocks github.com/lvxin1986/golang-test/mock/doer Doer
go test -v github.com/lvxin1986/golang-test/mock/user