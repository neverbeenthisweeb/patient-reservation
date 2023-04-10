image:
	docker build -t patientreservation:latest .

container:
	docker run --rm -p 4040:4040 patientreservation:latest

unittest:
	go test -v -failfast ./app