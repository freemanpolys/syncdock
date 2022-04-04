# Syncdock

Syncdock cli pull an image docker from some registry and push it to a local docker registry.
One use case it to handle easily docker images in airgaped environment.
## Installation

Download a binary in the [release](https://github.com/freemanpolys/syncdock/releases) page and make it executable.

## Configuration

Configure the local repository url.

```bash
syncdock config -r mylocal.repo
```

## Usage

```shell
# Get the command help
syncdock  --help

# Pull the nginx image and push it to mylocal.repo registry
syncdock  -i nginx

# Pull the docker image with full url is usefull for non dockerhub repo
syncdock  -f docker.elastic.co/eck/eck-operator:2.1.0  -i eck/eck-operator:2.1.0

# Fetch image from local repo and tag it to another image
syncdock  fetch -i nginx -t nama/nginx
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
Apache 2.0