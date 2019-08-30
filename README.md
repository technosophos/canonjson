# CanonJSON: A Tool for Formatting JSON into Canonical JSON

This is a simple tool for reformatting a JSON document into [Canonical JSON](http://wiki.laptop.org/go/Canonical_JSON).
This tool relies on Docker's [implementation of Canonical JSON](https://github.com/docker/go/tree/master/canonical/json), which is used by [Notary](https://github.com/theupdateframework/notary) and [CNAB](https://github.com/deislabs/cnab-spec).

## Building

1. Clone this repo
2. Run `dep ensure`
3. Run `go build ./canonjson.go`
4. Copy `canonjson` to your desired location

## Running

The first argument is the input JSON file, and the second is an optional output file.

```console
$ canonjson example.json
{"credentials":{"hostkey":{"env":"HOST_KEY","path":"/etc/hostkey.txt"},"image_token":{"env":"AZ_IMAGE_TOKEN"},"kubeconfig":{"path":"/home/.kube/config"}},"description":"An example 'thick' helloworld Cloud-Native Application Bundle","images":[{"description":"helloworld microservice","digest":"sha256:bbbbbbbbbbbb...","image":"technosophos/helloworld:0.1.2","mediaType":"application/vnd.docker.distribution.manifest.v2+json","platform":{"architecture":"amd64","os":"linux"},"size":1337}],"invocationImages":[{"digest":"sha256:aaaaaaaaaaaa...","image":"technosophos/helloworld:1.2.3","imageType":"docker","mediaType":"application/vnd.docker.distribution.manifest.v2+json","platform":{"architecture":"amd64","os":"linux"},"size":1337}],"name":"helloworld","parameters":{"backend_port":{"defaultValue":80,"maxValue":10240,"metadata":{"description":"The port that the backend will listen on"},"minValue":10,"type":"int"}},"schemaVersion":"v1.0.0-WD","version":"1.0.0"}

$ canonjson example.json canonical.json
$ cat canonical.json
{"credentials":{"hostkey":{"env":"HOST_KEY","path":"/etc/hostkey.txt"},"image_token":{"env":"AZ_IMAGE_TOKEN"},"kubeconfig":{"path":"/home/.kube/config"}},"description":"An example 'thick' helloworld Cloud-Native Application Bundle","images":[{"description":"helloworld microservice","digest":"sha256:bbbbbbbbbbbb...","image":"technosophos/helloworld:0.1.2","mediaType":"application/vnd.docker.distribution.manifest.v2+json","platform":{"architecture":"amd64","os":"linux"},"size":1337}],"invocationImages":[{"digest":"sha256:aaaaaaaaaaaa...","image":"technosophos/helloworld:1.2.3","imageType":"docker","mediaType":"application/vnd.docker.distribution.manifest.v2+json","platform":{"architecture":"amd64","os":"linux"},"size":1337}],"name":"helloworld","parameters":{"backend_port":{"defaultValue":80,"maxValue":10240,"metadata":{"description":"The port that the backend will listen on"},"minValue":10,"type":"int"}},"schemaVersion":"v1.0.0-WD","version":"1.0.0"}
```
