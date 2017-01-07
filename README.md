[Website](https://appscode.com) • [Slack](https://slack.appscode.com) • [Forum](https://discuss.appscode.com) • [Twitter](https://twitter.com/AppsCodeHQ)

# appstream
3 dimensions of Apps API

## Build Instructions
```sh
# dev build
./hack/make.py

# Install/Update dependency (needs glide)
glide slow
```

## Test api
Run the server, by running:
```sh
appstream --v=10
```

Now open the following url:
##### Docker
http://127.0.0.1:50066/api/apps/v1beta1/metadata/docker?name=appscode/kubed
