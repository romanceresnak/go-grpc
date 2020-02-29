# Generate .pb go files based on the .proto file
START=${pwd}

protoc:
	protoc -I pb/v1/ \
	--go_out= plugins=grpc:pb	\
	--gogrpcmock)out=:pb	\
	  pb/v1/*proto

install:
	go get -u \
			github.com/golang/protobuf/proto	\
			github.com/golang/protobuf/protoc-gen-go	\
			google.golang.org/grpc	\
			github.com/gogo/protobuf/protoc-gen-gogoslick	\
			github.com/gogo/protobuf/gogoproto	\
			github.com/DATA-dog/go-sqlmock	\
			github.com/onsi/ginkgo/ginkgo	\
			github.com/onsi/gomega/...	\
			github.com/SafetyCulture/s12-proto/protobuf/protoc-gen-gogrpcmock	\
			gopkg.in/go-playground/validator.v9
		go install github.com/SafetyCulture/s12-proto/protobuf/protoc-gen-gogrpcmock
		go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

clean:
	rm ./pb/**/*.pb.go

test:
	ginko -r -failFast

# migrate create -dir migrations -ext sql create_users_table
# migrate -path ./migrations -database mysql://root:Geqnnvpgv7@/grpc -verbose up