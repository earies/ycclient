PROGRAM=ycclient

VERSION=0.0.0
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(PROGRAM)

$(PROGRAM):
	go build ${LDFLAGS} -o ${PROGRAM}

.PHONY: clean
clean:
	if [ -f ${PROGRAM} ]; then rm ${PROGRAM}; fi
