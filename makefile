.PHONY: build

tag = shafinhasnat/rebis

build:
	go build -ldflags "-s -w" -o ./dist/rebis_$(VERSION) .
	docker build -t $(tag):$(VERSION) .
	docker tag $(tag):$(VERSION) $(tag):latest

tag:
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)

push:
	docker push $(tag):$(VERSION)
	docker push $(tag):latest
