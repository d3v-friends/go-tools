## 용어 설명

1. protoc -> *.proto -> 컴파일 -> 타겟코드(go, js, java ...)
2. protoc-gen-go -> protoc 에 사용되는 golang 플러그인, message 를 번역해준다
3. protoc-gen-go-grpc -> protoc 에 사용되는 golang 플러그인, service 를 번역해준다

## 사용방법

1. protoc 설치
   https://github.com/protocolbuffers/protobuf/releases
2. 들어가서 바이너리 다운받아서 $HOME/sdk/bin 에 추가
	- 경로는 예제이며 필요에 따라 변경하여 사용한다
3. protobuf 인스톨 하기
	* $GOPATH/bin 에 protoc extension 이 설치된다

   ~~~bash
   go get google.golang.org/protobuf/cmd
   go install google.golang.org/protobuf/cmd/protoc-gen-go
   
   go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
   ~~~

4. cli 사용하기

   ~~~bash
   # 설치
   go install github.com/d3v-friends/go-tools/fnGrpc/go-grpc
   
   # 사용
   go-grpc --config grpc.yml
   ~~~


## 프로파일링
~~~bash
go get github.com/google/pprof
~~~
