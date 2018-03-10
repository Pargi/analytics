# Pargi Analytics

This repository contains a simple backend service, built on the Serverless framework, which exposes an endpoint for gathering analytics data from Pargi. As the goal of this service is super simple, it does not currently offer any test coverage for its features or in-depth documentation.

## API

The API for the service is documented in [API.md](/API.md) as well as the further commented [Paw document](/API.paw) (for reference, [Paw](https://paw.cloud) is a tool for testing APIs).

### API Key

As part of the deployment, Serverless creates all the necessary resources for the service, including an API key that is used to protect the endpoints. This API key should be embedded in the app in order to enable analytics gathering. The key (as well as the endpoint) used in the App Store version of Pargi is not exposed in any of the repositories.

## Development

The service is built on top of [Serverless](https://serverless.com) using the Go language. Refer to Serverless docs for how to deploy/develop/maintain the service. The following commands should enable deploying a dev stage of the service.

```
dep ensure
make
serverless deploy
```

These assume a functioning install of [Go](https://golang.org), [dep](https://github.com/golang/dep) and [Serverless](https://serverless.com). All of these can be installed with these commands

```
npm i -g serverless
brew update
brew install go
brew install dep
```
