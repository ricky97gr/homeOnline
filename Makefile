GinLogMode=""


SystemName="社区系统"
GoVersion=$(shell go version | awk '{print $$3, $$4}')
BuildTime=$(shell date "+%Y-%m-%d %H:%M:%S")
GitCommitID=$(shell git rev-parse HEAD)
Version="0.0.1_base"


VER=debug


ifeq ($(VER),debug)
	GinLogMode="debug"
else
	GinLogMode="release"
endif


define start-build=
	@echo $(1)
	@make $(1)
endef


LDFLAGS="\
 -X 'github.com/ricky97gr/homeOnline/pkg/bininfo.GoVersion=$(GoVersion)'\
 -X 'github.com/ricky97gr/homeOnline/pkg/bininfo.SystemName=$(SystemName)'\
 -X 'github.com/ricky97gr/homeOnline/pkg/bininfo.CommitID=$(GitCommitID)'\
 -X 'github.com/ricky97gr/homeOnline/pkg/bininfo.BuildTime=$(BuildTime)'\
 -X 'github.com/ricky97gr/homeOnline/pkg/bininfo.Version=$(Version)'\
 -X 'github.com/ricky97gr/homeOnline/pkg/bininfo.GinLogMode=${GinLogMode}'\
"



.PHONY:
prepare:
	@echo "\n\n\n" >> /etc/hosts
	@echo "192.168.0.200 mysql.test.com" >> /etc/hosts
	@echo "192.168.0.200 redis.test.com" >> /etc/hosts
	@echo "192.168.0.200 mongo.test.com" >> /etc/hosts
	
	@echo "10.182.34.112 mysql.test.com" >> /etc/hosts
	@echo "10.182.34.112 redis.test.com" >> /etc/hosts
	@echo "10.182.34.112 mongo.test.com" >> /etc/hosts
	@echo "system prepare ready"
	@cp ./internal/conf/config.yaml /root/tmp/.config.yaml
	@echo "prepare db env"
	@bash ./scripts/prepare.sh


.PHONY:
messagegrpc:
	@protoc --go_out=./internal/grpcserver/proto/message/ ./internal/grpcserver/proto/message/message.proto
	@protoc --go-grpc_out=./internal/grpcserver/proto/message/ ./internal/grpcserver/proto/message/message.proto
	@echo "message grpc protobuf generate successfully"


.PHONY:
loggrpc:
	@protoc --go_out=./internal/grpcserver/proto/log/ ./internal/grpcserver/proto/log/log.proto
	@protoc --go-grpc_out=./internal/grpcserver/proto/log/ ./internal/grpcserver/proto/log/log.proto
	@echo "log grpc protobuf generate successfully"


.PHONY:
noticegrpc:
	@protoc --go_out=./internal/grpcserver/proto/station_notice/ ./internal/grpcserver/proto/station_notice/station_notice.proto
	@protoc --go-grpc_out=./internal/grpcserver/proto/station_notice/ ./internal/grpcserver/proto/station_notice/station_notice.proto
	@echo "station_notice grpc protobuf generate successfully"
	

.PHONY:
release:
	@echo "start to build release version"
	@rm -rf bin
	@mkdir bin
	@go build -ldflags ${LDFLAGS} -o bin/family_webservice cmd/webservice/main.go
	@echo "build family_webservice successfully!"
	

.PHONY:
debug:
	@echo "start to build debug version"
	@rm -rf bin
	@mkdir bin
	@go build -ldflags ${LDFLAGS} -o bin/family_webservice cmd/webservice/main.go
	@echo "build family_webservice successfully!"


.PHONY:
run:
ifeq ($(VER), debug)
	@$(call start-build,debug)
else
	@$(call start-build,release)
endif


