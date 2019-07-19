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

### Create or update one ILM policy from file

It's permit to create or update Life cycle policy on Elasticsearch from file.

You need to set the following parameters:
- **--lifecycle-policy-id**: The unique name of lifecycle policy you should to create or update
- **--lifecycle-policy-file**: The full path of file that contain the lifecycle policy.


Sample of lifecycle policy call `sample-ilm.json`:
```json
{
  "policy": {
    "phases": {
      "hot": {
        "actions": {
          "rollover": {
            "max_size": "30GB",
            "max_age": "1d"
          },
          "set_priority" : {
            "priority": 100
          }
        }
      },
      "warm": {
        "actions": {
          "forcemerge": {
            "max_num_segments": 1
          },
          "shrink": {
              "number_of_shards": 1
          },
          "set_priority" : {
            "priority": 50
          },
          "readonly": {}
        }
      },
      "delete": {
        "min_age": "30d",
        "actions": {
          "delete": {}
        }
      }
    }
  }
}


```

Sample of command:
```bash
elktools_linux_amd64 --elasticsearch-url https://elasticsearch.company.com --elasticsearch-user elastic --elasticsearch-password changeme --self-signed-certificate create-or-update-lifecycle-policy --lifecycle-policy-id test --lifecycle-policy-file sample-ilm.json 
```

### Create or updates multiples ILM policy from folder

It permit to create or update multiple lifecycle policy from folder where each files contains lifecycle policy.

> The lifecycle policy unique name is based on the file name.

You need to set the following parameters:
- **--lifecycle-policy-base-path**: The full path that contains lifecycle policies files.

Sample of command:
```bash
elktools_linux_amd64 --elasticsearch-url https://elasticsearch.company.com --elasticsearch-user elastic --elasticsearch-password changeme --self-signed-certificate create-or-update-all-lifecycle-policies --lifecycle-policy-base-path ilm-policies/
```

### Save ILM policy on file

It permit to save existing lifecycle policy from Elasticsearch on file.

You need to set the following parameters:
- **--lifecycle-policy-id**: The unique name of lifecycle policy you should to save
- **--lifecycle-policy-file**: The full path of file where you should to save the lifecycle policy

Sample of command:
```bash
elktools_linux_amd64 --elasticsearch-url https://elasticsearch.company.com --elasticsearch-user elastic --elasticsearch-password changeme --self-signed-certificate save-lifecycle-policy --lifecycle-policy-id test --lifecycle-policy-file backup-ilm.json 
```

### Save all ILM policies on files

It permit to save all existing lifecycle policies from Elasticsearch in files. Each lifecycle policy is store in its own file.

> The file name is base on the unique name of the lifecycle policy.

You need to set the following parameters:
- **--lifecycle-policy-base-path**: The full path where store the lifecycle policies files.

Sample of command:
```bash
elktools_linux_amd64 --elasticsearch-url https://elasticsearch.company.com --elasticsearch-user elastic --elasticsearch-password changeme --self-signed-certificate save-all-lifecycle-policies --lifecycle-policy-base-path ilm-policies/
```

### Delete ILM policy

It permit to delete existing lifecycle policy on Elasticsearch.

You need to set the following parameters:
- **--lifecycle-policy-id**: The unique name of lifecycle policy you should to delete

Sample of command:
```bash
elktools_linux_amd64 --elasticsearch-url https://elasticsearch.company.com --elasticsearch-user elastic --elasticsearch-password changeme --self-signed-certificate delete-lifecycle-policy --lifecycle-policy-id test
```

### Get the ILM policy state on indice

It permit to get the actual state of lifecycle policy on particular Elasticsearch indice.

You need to set the following parameters:
- **--elasticsearch-index**: The Elasticsearch indice name where you should to get the lifecycle policy state.

Sample of command:
```bash
elktools_linux_amd64 --elasticsearch-url https://elasticsearch.company.com --elasticsearch-user elastic --elasticsearch-password changeme --self-signed-certificate get-lifecycle-policy-status --elasticsearch-index logstash-2019.01.01
```