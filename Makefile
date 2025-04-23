# Makefile template: https://gist.github.com/grihabor/4a750b9d82c9aa55d5276bd5503829be

MAKE               := make --no-print-directory

DESCRIBE           := $(shell git describe --match "v*" --always --tags)
DESCRIBE_PARTS     := $(subst -, ,$(DESCRIBE))

VERSION_TAG        := $(word 1,$(DESCRIBE_PARTS))
COMMITS_SINCE_TAG  := $(word 2,$(DESCRIBE_PARTS))

VERSION            := $(subst v,,$(VERSION_TAG))
VERSION_PARTS      := $(subst ., ,$(VERSION))

MAJOR              := $(word 1,$(VERSION_PARTS))
MINOR              := $(word 2,$(VERSION_PARTS))
MICRO              := $(word 3,$(VERSION_PARTS))

NEXT_MAJOR         := $(shell echo $$(($(MAJOR)+1)))
NEXT_MINOR         := $(shell echo $$(($(MINOR)+1)))
NEXT_MICRO          = $(shell echo $$(($(MICRO)+$(COMMITS_SINCE_TAG))))

ifeq ($(strip $(COMMITS_SINCE_TAG)),)
CURRENT_VERSION_MICRO := $(MAJOR).$(MINOR).$(MICRO)
CURRENT_VERSION_MINOR := $(CURRENT_VERSION_MICRO)
CURRENT_VERSION_MAJOR := $(CURRENT_VERSION_MICRO)
else
CURRENT_VERSION_MICRO := $(MAJOR).$(MINOR).$(NEXT_MICRO)
CURRENT_VERSION_MINOR := $(MAJOR).$(NEXT_MINOR).0
CURRENT_VERSION_MAJOR := $(NEXT_MAJOR).0.0
endif

DATE                = $(shell date +'%d.%m.%Y')
TIME                = $(shell date +'%H:%M:%S')
COMMIT             := $(shell git rev-parse HEAD)
AUTHOR             := $(firstword $(subst @, ,$(shell git show --format="%aE" $(COMMIT))))
BRANCH_NAME        := $(shell git rev-parse --abbrev-ref HEAD)

TAG_MESSAGE         = "$(TIME) $(DATE) $(AUTHOR) $(BRANCH_NAME)"
COMMIT_MESSAGE     := $(shell git log --format=%B -n 1 $(COMMIT))

CURRENT_TAG_MICRO  := "v$(CURRENT_VERSION_MICRO)"
CURRENT_TAG_MINOR  := "v$(CURRENT_VERSION_MINOR)"
CURRENT_TAG_MAJOR  := "v$(CURRENT_VERSION_MAJOR)"

# --- Version commands ---
.PHONY: version version-micro version-minor version-majo

version:
	@$(MAKE) version-micro

version-micro:
	@echo "$(CURRENT_VERSION_MICRO)"

version-minor:
	@echo "$(CURRENT_VERSION_MINOR)"

version-major:
	@echo "$(CURRENT_VERSION_MAJOR)"

# --- Tag commands ---
.PHONY: tag-micro tag-minor tag-major

tag-micro:
	@echo "$(CURRENT_TAG_MICRO)"

tag-minor:
	@echo "$(CURRENT_TAG_MINOR)"

tag-major:
	@echo "$(CURRENT_TAG_MAJOR)"

# -- Meta info ---
.PHONY: tag-message commit-message

tag-message:
	@echo "$(TAG_MESSAGE)"

commit-message:
	@echo "$(COMMIT_MESSAGE)"


# -- Project --
.PHONY: run tmp

# requires wgo: https://github.com/bokwoon95/wgo
run:
	wgo run ./tests/main.go
