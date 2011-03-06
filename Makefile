all: maps.6 test

test : maps_test.6
	6l -otest maps_test.6 
maps_test.6 : maps_test.go
	6g maps_test.go
maps.6 : maps.go
	6g maps.go
clean:
	rm *.6 test

