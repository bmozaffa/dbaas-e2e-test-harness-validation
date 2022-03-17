NAME=validate-e2e
TAG=quay.io/bmozaffa/$(NAME)
VER=v1.0

all: clean compile build-image

compile:
	GOOS=linux go build -o validate-e2e app.go

build-image:
	podman build -t $(TAG) -t $(TAG):$(VER) .

run:
	podman run --name=$(NAME) $(TAG)

clean:
	-podman stop $(NAME)
	-podman rm $(NAME)

push:
	podman push $(TAG)
	podman push $(TAG):$(VER)
