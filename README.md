# Nginx Log Generator

A tiny Go utility to generate a large amount realistic-looking Nginx logs quickly. It was written to aid in testing logging pipelines and other such tools, and demoing them in Kubernetes.

Most of the heavy lifting is done by the amazing [gofakeit](https://github.com/brianvoe/gofakeit) library, with some extra work to skew the results towards typical values.

## Usage

The most important step is to set the desired rate in the `RATE` environment variable. The simplest way to do this is the following:

```sh
$ # Will generate 10 entries per second
$ RATE=10 ./nginx-log-generator
```

The reason this being an environment variable is so it's easier to run via Docker as well:

```sh
$ docker pull apassionatechie/nginx-log-generator
$ docker run -e "RATE=10" apassionatechie/nginx-log-generator
```

The format of the logs is as follows:
```sh
ip, time, httpMethod, path, httpVersion, statusCode, responseTime, upstream_ip, port, bodyBytesSent, referrer, userAgent, ssl_protocol, content_type, host
```
