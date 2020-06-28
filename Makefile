bkname = ./backup/dplt-$(shell date +"%Y-%m-%d-%H-%M")
pid = $(shell cat .pid +"%T")
logfile = log/$(pid).log

default: build install

build:
	@mkdir -p bin
	@echo "STEP : BUILD"
	@go build -o bin/dplt
	@echo "Build successfully."

install:
	@echo "STEP : INSTALL"
	@go install
	@echo "Install successfully."

back-up:
	@echo "STEP : BACK UP"
	@mkdir -p ./backup
	@cp ./bin/dplt $(bkname)
	@echo "Done"

stop:
	@clear
	@echo "Kill PID = $(pid)"
	@kill $(pid)
	@echo "=> done.."
