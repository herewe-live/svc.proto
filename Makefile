PROJECT=svc.proto
PREFIX=$(shell pwd)
VERSION=$(shell git describe --match 'v[0-9]*'  --always)
DEFAULT_BRANCH=$(shell git symbolic-ref --short -q HEAD)

ifndef OS
	OS=linux
	unameOut=$(shell uname -s)
	ifeq ($(unameOut),Darwin)
		OS=darwin
	endif

	ifeq ($(OSTYPE),win32)
		OS=windows
	endif
endif

ifndef ARCH
	ARCH=amd64
	unameOut=$(shell uname -m)
	ifeq ($(unameOut),i386)
		ARCH=386
	endif

	ifeq ($(unameOut),i686)
		ARCH=386
	endif
endif

ifndef BRANCH
	BRANCH=$(DEFAULT_BRANCH)
endif

ifdef CI_COMMIT_REF_SLUG
	BRANCH=$(CI_COMMIT_REF_SLUG)
endif

ifndef GO
	GO=go
endif

ifndef GOFMT
	GOFMT=gofmt
endif

ifndef PROTOC
	PROTOC=protoc
endif

ifndef GIT
	GIT=git
endif

SOURCE_DIR=$(PREFIX)
DOCS_DIR=$(PREFIX)/docs

.PHONY: all summary micro build doc
.DEFAULT: all

# Targets
all: summary build doc

summary:
	@printf "\033[1;37m  == \033[1;32m$(PROJECT) \033[1;33m$(VERSION) \033[1;37m==\033[0m\n"
	@printf "    Platform    : \033[1;37m$(shell uname -sr)\033[0m\n"
	@printf "    Target OS   : \033[1;37m$(OS)\033[0m\n"
	@printf "    Target Arch : \033[1;37m$(ARCH)\033[0m\n"
	@printf "    Go          : \033[1;37m`$(GO) version`\033[0m\n"
	@printf "    Git         : \033[1;37m$(shell git version)\033[0m\n"
	@printf "    Branch      : \033[1;37m$(BRANCH)\033[0m\n"
	@echo

build:
	@printf "\033[1;36m  Compiling protos ...\033[0m\n"
	@for f in $(shell find . -name '*.proto') ; do \
		printf "    Compiling <$${f}>\n"  && \
		$(PROTOC) --go_out=. --go-grpc_out=. $${f} ; \
	done
	@echo

grpc:
	@printf "\033[1;36m  Upgrading gRPC compiler toolchains ...\033[0m\n"
	@GOOS=$(OS) GOARCH=$(ARCH) $(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@GOOS=$(OS) GOARCH=$(ARCH) $(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@GOOS=$(OS) GOARCH=$(ARCH) $(GO) install github.com/envoyproxy/protoc-gen-validate@latest
	@GOOS=$(OS) GOARCH=$(ARCH) $(GO) install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	@echo

doc:
	@printf "\033[1;36m  Generating document ...\033[0m\n"
	@mkdir -p $(PREFIX)/doc
	@$(PROTOC) --doc_opt=markdown,doc.md --doc_out=./doc $(shell find . -name '*.proto')
	@echo
