# elktools
A cli tools to help manage ELK

## Contribute

You PR are always welcome. Please use the righ branch to do PR:
 - 7.x for Elasticsearch 7.x
 - 6.x for Elasticsearch 6.x
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
- **--url**: The Elasticsearch or Kibana URL. For exemple https://elasticsearch.company.com. Alternatively you can use environment variable `ELASTICSEARCH_URL`.
- **--user**: The login to connect on Elasticsearch. Alternatively you can use environment variable `ELASTICSEARCH_USER`.
- **--password**: The password to connect on Elasticsearch. Alternatively you can use environment variable `ELASTICSEARCH_PASSWORD`.
- **--self-signed-certificate**: Disable the check of server SSL certificate
- **--debug**: Enable the debug mode
- **--help**: Display help for the current command


You can set also this parameters on yaml file (one or all) and use the parameters `--config` with the path of your Yaml file.
```yaml
---
url: https://elasticsearch.company.com
user: elastic
password: changeme
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
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate create-or-update-lifecycle-policy --lifecycle-policy-id test --lifecycle-policy-file sample-ilm.json
```

### Create or updates multiples ILM policy from folder

It permit to create or update multiple lifecycle policy from folder where each files contains lifecycle policy.

> The lifecycle policy unique name is based on the file name.

You need to set the following parameters:
- **--lifecycle-policy-base-path**: The full path that contains lifecycle policies files.

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate create-or-update-all-lifecycle-policies --lifecycle-policy-base-path ilm-policies/
```

### Save ILM policy on file

It permit to save existing lifecycle policy from Elasticsearch on file.

You need to set the following parameters:
- **--lifecycle-policy-id**: The unique name of lifecycle policy you should to save
- **--lifecycle-policy-file**: The full path of file where you should to save the lifecycle policy

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate save-lifecycle-policy --lifecycle-policy-id test --lifecycle-policy-file backup-ilm.json
```

### Save all ILM policies on files

It permit to save all existing lifecycle policies from Elasticsearch in files. Each lifecycle policy is store in its own file.

> The file name is base on the unique name of the lifecycle policy.

You need to set the following parameters:
- **--lifecycle-policy-base-path**: The full path where store the lifecycle policies files.

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate save-all-lifecycle-policies --lifecycle-policy-base-path ilm-policies/
```

### Delete ILM policy

It permit to delete existing lifecycle policy on Elasticsearch.

You need to set the following parameters:
- **--lifecycle-policy-id**: The unique name of lifecycle policy you should to delete

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate delete-lifecycle-policy --lifecycle-policy-id test
```

### Get the ILM policy state on indice

It permit to get the actual state of lifecycle policy on particular Elasticsearch indice.

You need to set the following parameters:
- **--elasticsearch-index**: The Elasticsearch indice name where you should to get the lifecycle policy state.

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate get-lifecycle-policy-status --elasticsearch-index logstash-2019.01.01
```

### Create or update indice template from file

It permit to create or update indice template on Elasticsearch from file.

You need to set the following parameters:
- **--indice-template-id**: The unique indice template name you should to create or update
- **--indice-template-file**: The full path of file that contain the indice template to create or update.

Sample of indice template call `sample-indice-template.json`:
```json
{
  "index_patterns": [
    "logstash-*"
  ],
  "order": 2,
  "settings": {
    "index.lifecycle.name": "policy-logstash-log",
    "index.refresh_interval": "5s"
  }
}
```

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate create-or-update-indice-template --indice-template-id logstash-log --indice-template-file sample-indice-template.json
```

### Create or update all indice templates from folder

It permit to create or update all indice templates contain in folder where each file contain one indice template.

> The unique name of indice template is based on file name.

You need to set the following parameters:
- **--indice-template-path**: The full path of folder that contain the indice templates to create or update

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate create-or-update-all-indice-templates --indice-template-path indice-templates/
```

### Delete one indice template

It permit to delete one indice template in Elasticsearch.

You need to set the following parameters:
- **--indice-template-id**: The unique name of indice template you should to delete

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate delete-indice-template --indice-template-id logstash-log
```

### Save indice template in file

It permit to save one existing indice template from Elasticsearch to file.

You need to set the following parameters:
- **--indice-template-id**: The unique name of indice template you should to save
- **--indice-template-file**: The full path of file where you should to store the indice template

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate save-indice-template --indice-template-id logstash-log --indice-template-file template-backup.json
```

### Create indice from file

It permit to create new indice on Elasticsearch. It usefull for exemple create rolleover indice.

You need to set the following parameters:
- **indice-name**: The indice name you should to create
- **indice-setting-file**: The full path of file that contain the indice settings


Sample of indice setting call `indice-setting.json`:
```json
{
  "settings": {
    "number_of_shards": "2",
    "number_of_replicas": "1"
  },
  "aliases": {
    "logstash-log-alias": {
      "is_write_index": true
    }
  }
}
```

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate create-indice --indice-name logstash-log-000001 --indice-setting-file indice-setting.json
```

### Disable shard allocation

It permit to disable shard allocation. It usefull when reboot or upgrade nodes.

There are no parameter

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate disable-routing-allocation
```

### Enable shard allocation

It permit to enable shard allocation. It usefull when reboot or upgrade nodes.

There are no parameter

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate enable-routing-allocation
```

### Stop task for machine learning

It permit to temporarily stop the tasks associated with active machine leaning jobs and datafeeds. It usefull when reboot or upgrade nodes.

There are no parameter

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate enable-ml-upgrade
```

### Start task for machine learning

It permit to start the tasks associated with active machine leaning jobs and datafeeds. It usefull when reboot or upgrade nodes.

There are no parameter

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate disable-ml-upgrade
```

### Stop Watcher service

It permit to stop watcher service. It usefull when reboot or upgrade nodes.

There are no parameter

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate stop-watcher-service
```

### Start Watcher service

It permit to start watcher service. It usefull when reboot or upgrade nodes.

There are no parameter

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate start-watcher-service
```

### Stop ILM service

It permit to stop Index Lifecycle Management service. It usefull when reboot or upgrade nodes.

There are no parameter

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate stop-ilm-service
```

### Start ILM service

It permit to start Index Lifecycle Management service. It usefull when reboot or upgrade nodes.

There are no parameter

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate start-ilm-service
```

### Stop SLM service

It permit to stop Snapshot Lifecycle Management service. It usefull when reboot or upgrade nodes.

There are no parameter

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate stop-slm-service
```

### Start SLM service

It permit to start Snapshot Lifecycle Management service. It usefull when reboot or upgrade nodes.

There are no parameter

Sample of command:
```bash
elktools_linux_amd64 --url https://elasticsearch.company.com --user elastic --password changeme --self-signed-certificate start-slm-service
```

### Export all Kibana dashboards

It's permit to export all Kibana dashboards using Kibana API.

You need to set the following parameters:
- **--file-path**: The full path where store exported dashboard and references.
- **--user-space**: The kibana user space where to retrive dashboards. Default to `default`.

Sample of command:
```bash
elktools_linux_amd64 --url https://kibana.company.com --user elastic --password changeme --self-signed-certificate export-all-dashboards --file-path export.ndjson --user-space defaut
```





### Import all Kibana dashboards

It's permit to import all Kibana dashboards using Kibana API.

You need to set the following parameters:
- **--file-path**: The full path where load dashboards and references.
- **--user-space**: The kibana user space where to load dashboards. Default to `default`.



Sample of command:
```bash
elktools_linux_amd64 --url https://kibana.company.com --user elastic --password changeme --self-signed-certificate import-all-dashboards --file-path export.ndjson --user-space defaut
```