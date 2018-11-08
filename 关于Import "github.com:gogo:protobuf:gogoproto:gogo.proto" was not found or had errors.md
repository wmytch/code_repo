# 关于`Import "github.com/gogo/protobuf/gogoproto/gogo.proto" was not found or had errors`  
实际上这个错误并不是找不到`gogo.proto`,而是找不到其中`import`的`google/protobuf/descriptor.proto`.  
所以需要指定这个文件所在的目录，比方说`$GOPATH/src/google/protobuf/`。  
执行命令：`protoc -I$GOPATH/src/ --go_out=. abci-test/types/types.proto`  
在`$GOPATH/src/`中存在目录`google/protobuf/`.  
另外这个命令应该在`abci-test`的上一级目录执行，也就是`$GOPATH/src`(假定`abci-test/`是`$GOPATH/src/`的子目录),这样就会在`$GOPATH/src/abci-test/types/`下生成相应的`types.pb.go`文件,不要问我为什么。
