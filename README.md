# elktools
A cli tools to help manage ELK

## Contribute

You PR are always welcome. Please use develop branch to do PR (git flow pattern)
Don't forget to add test if you add some functionalities.

To build, you can use the following command line:
```sh
make build
```

To lauch golang test, you can use the folowing command line:
```sh
make test
```

## CLI

### Global options

The following parameters are available for all commands line :
- **--elasticsearch-url**: The Elasticsearch URL. For exemple https://elasticsearch.company.com. Alternatively you can use environment variable `ELASTICSEARCH_URL`.
- **--elasticsearch-user**: The login to connect on Elasticsearch. Alternatively you can use environment variable `ELASTICSEARCH_USER`.
- **--elasticsearch-password**: The password to connect on Elasticsearch. Alternatively you can use environment variable `ELASTICSEARCH_PASSWORD`.
- **--self-signed-certificate**: Disable the check of server SSL certificate 
- **--debug**: Enable the debug mode
- **--help**: Display help for the current command


You can set also this parameters on yaml file (one or all) and use the parameters `--config` with the path of your Yaml file.
```yaml
---
elasticsearch-url: https://elasticsearch.company.com
elasticsearch-user: elastic
elasticsearch-password: changeme
```
