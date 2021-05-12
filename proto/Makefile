# See README.txt.
go:     add_person_go     list_people_go
gotest: add_person_gotest list_people_gotest

clean:
	rm -f add_person_cpp list_people_cpp add_person_java list_people_java add_person_python list_people_python
	rm -f javac_middleman AddPerson*.class ListPeople*.class com/example/tutorial/*.class
	rm -f protoc_middleman addressbook.pb.cc addressbook.pb.h addressbook_pb2.py com/example/tutorial/AddressBookProtos.java
	rm -f *.pyc
	rm -f protoc_middleman_go tutorial/*.pb.go add_person_go list_people_go go.mod go.sum
	rm -f protoc_middleman_dart dart_tutorial/*.pb*.dart
	rmdir dart_tutorial 2>/dev/null || true
	rmdir tutorial 2>/dev/null || true
	rmdir com/example/tutorial 2>/dev/null || true
	rmdir com/example 2>/dev/null || true
	rmdir com 2>/dev/null || true

protoc_middleman_go: addressbook.proto
	mkdir -p tutorial # make directory for go package
	protoc --go_out=tutorial ./addressbook.proto

go_mod:
	go mod init proto
	go mod tidy

add_person_go: add_person.go protoc_middleman_go go_mod
	go build -o add_person_go add_person.go

add_person_gotest: add_person_test.go add_person_go go_mod
	go test add_person.go add_person_test.go

list_people_go: list_people.go protoc_middleman_go go_mod
	go build -o list_people_go list_people.go

list_people_gotest: list_people.go list_people_go go_mod
	go test list_people.go list_people_test.go