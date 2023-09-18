package main

import (
	"fmt"
	"os"
	"sort"

	elktools_elasticsearch "github.com/disaster37/elktools/v8/elasticsearch"
	elktools_kibana "github.com/disaster37/elktools/v8/kibana"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	version string
	commit  string
)

func run(args []string) error {

	// Logger setting
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	formatter.ForceFormatting = true
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)

	// CLI settings
	app := cli.NewApp()
	app.Usage = "Manage ELK on cli interface"
	app.Version = fmt.Sprintf("%s-%s", version, commit)
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Usage: "Load configuration from `FILE`",
		},
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "url",
			Usage:   "The elasticsearch or kibana URL",
			EnvVars: []string{"ELKTOOLS_URL"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "user",
			Usage:   "The  user",
			EnvVars: []string{"ELKTOOLS_USER"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "password",
			Usage:   "The password",
			EnvVars: []string{"ELKTOOLS_PASSWORD"},
		}),
		&cli.BoolFlag{
			Name:  "self-signed-certificate",
			Usage: "Disable the TLS certificate check",
		},
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "Display debug output",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:     "check-connexion-elasticsearch",
			Usage:    "Check the elasticsearch connexion",
			Category: "Check",
			Action:   elktools_elasticsearch.CheckConnexion,
		},
		{
			Name:     "create-or-update-lifecycle-policy",
			Usage:    "Create or update lifecycle policy",
			Category: "ILM policy actions",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "lifecycle-policy-id",
					Usage: "The lifecycle policy name",
				},
				&cli.StringFlag{
					Name:  "lifecycle-policy-file",
					Usage: "The full path of lifecycle policy file",
				},
			},
			Action: elktools_elasticsearch.CreateILMPolicy,
		},
		{
			Name:     "save-lifecycle-policy",
			Usage:    "Save lifecycle policy on file",
			Category: "ILM policy actions",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "lifecycle-policy-id",
					Usage: "The lifecycle policy name",
				},
				&cli.StringFlag{
					Name:  "lifecycle-policy-file",
					Usage: "The full path of lifecycle policy file",
				},
			},
			Action: elktools_elasticsearch.SaveILMPolicy,
		},
		{
			Name:     "delete-lifecycle-policy",
			Usage:    "Delete lifecycle policy",
			Category: "ILM policy actions",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "lifecycle-policy-id",
					Usage: "The lifecycle policy name",
				},
			},
			Action: elktools_elasticsearch.DeleteILMPolicy,
		},
		{
			Name:     "save-all-lifecycle-policies",
			Usage:    "Save all lifecycle policies on folder provided",
			Category: "ILM policy actions",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "lifecycle-policy-base-path",
					Usage: "The base path to save all lifecycle policy",
					Value: "lifecycle_policy",
				},
			},
			Action: elktools_elasticsearch.SaveAllILMPolicies,
		},
		{
			Name:     "create-or-update-all-lifecycle-policies",
			Usage:    "Create or update all lifecycle policies from folder provided",
			Category: "ILM policy actions",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "lifecycle-policy-base-path",
					Usage: "The base path to load all lifecycle policy",
					Value: "lifecycle_policy",
				},
			},
			Action: elktools_elasticsearch.CreateAllILMPolicies,
		},
		{
			Name:     "get-lifecycle-policy-status",
			Usage:    "Get the lifecycle policy on Elasticsearch index",
			Category: "ILM policy actions",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "elasticsearch-index",
					Usage: "The Elasticsearch index where you should to get lifecycle policy status",
				},
			},
			Action: elktools_elasticsearch.GetStatusILMPolicy,
		},
		{
			Name:     "create-or-update-indice-template",
			Usage:    "Create or update indice template",
			Category: "Indice template actions",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "indice-template-id",
					Usage: "The indice template name",
				},
				&cli.StringFlag{
					Name:  "indice-template-file",
					Usage: "The full path of indice template file",
				},
			},
			Action: elktools_elasticsearch.CreateIndiceTemplate,
		},
		{
			Name:     "create-or-update-all-indice-templates",
			Usage:    "Create or update all indice templates",
			Category: "Indice template actions",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "indice-template-path",
					Usage: "The indice templates base path",
				},
			},
			Action: elktools_elasticsearch.CreateAllIndiceTemplates,
		},
		{
			Name:     "delete-indice-template",
			Usage:    "Delete indice template",
			Category: "Indice template actions",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "indice-template-id",
					Usage: "The indice template ID to delete",
				},
			},
			Action: elktools_elasticsearch.DeleteIndiceTemplate,
		},
		{
			Name:     "save-indice-template",
			Usage:    "Save indice template on file",
			Category: "Indice template actions",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "indice-template-id",
					Usage: "The indice template ID to save",
				},
				&cli.StringFlag{
					Name:  "indice-template-file",
					Usage: "The indice template file to store it",
				},
			},
			Action: elktools_elasticsearch.SaveIndiceTemplate,
		},
		{
			Name:     "create-indice",
			Usage:    "Create new indice with settings",
			Category: "Indice actions",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "indice-name",
					Usage: "The indice name to create",
				},
				&cli.StringFlag{
					Name:  "indice-setting-file",
					Usage: "The indice setting file",
				},
			},
			Action: elktools_elasticsearch.CreateIndice,
		},
		{
			Name:     "check-connexion-kibana",
			Usage:    "Check the kibana connexion",
			Category: "Check",
			Action:   elktools_kibana.CheckConnexion,
		},
		{
			Name:     "export-all-dashboards",
			Usage:    "Export all dashboards and all references in ndjson file",
			Category: "Kibana dashboard",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "file-path",
					Usage: "The file path to export all dashboards",
				},
				&cli.StringFlag{
					Name:  "user-space",
					Usage: "The Kibana user space where export dashboards",
					Value: "default",
				},
			},
			Action: elktools_kibana.ExportDashboards,
		},
		{
			Name:     "import-all-dashboards",
			Usage:    "Import all dashboards and all references from ndjson file",
			Category: "Kibana dashboard",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "file-path",
					Usage: "The file path to load all dashboards",
				},
				&cli.StringFlag{
					Name:  "user-space",
					Usage: "The Kibana user space where write dashboards",
					Value: "default",
				},
			},
			Action: elktools_kibana.ImportDashboards,
		},
		{
			Name:     "check-elasticsearch-status",
			Usage:    "Check the elasticsearch status",
			Category: "Check",
			Action:   elktools_elasticsearch.CheckClusterStatus,
		},
		{
			Name:     "check-node-online",
			Usage:    "Check the node is online on Elasticsearch cluster",
			Category: "Check",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "node-name",
					Usage:    "The node name",
					Required: true,
				},
				&cli.StringSliceFlag{
					Name:     "labels",
					Usage:    "The labels to check the node name",
					Required: false,
				},
			},
			Action: elktools_elasticsearch.CheckNodeOnline,
		},
		{
			Name:     "check-number-nodes",
			Usage:    "Check there are a number of node in cluster",
			Category: "Check",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:     "number-nodes",
					Usage:    "The number of node expected",
					Required: true,
				},
			},
			Action: elktools_elasticsearch.CheckExpectedNumberNodes,
		},
		{
			Name:     "disable-routing-allocation",
			Usage:    "Disable routing allocation on Elasticsearch cluster",
			Category: "Downtime",
			Action:   elktools_elasticsearch.ClusterDisableRoutingAllocation,
		},
		{
			Name:     "enable-routing-allocation",
			Usage:    "Enable routing allocation on Elasticsearch cluster",
			Category: "Downtime",
			Action:   elktools_elasticsearch.ClusterEnableRoutingAllocation,
		},
		{
			Name:     "enable-ml-upgrade",
			Usage:    "Enable upgrade mode on ML",
			Category: "Downtime",
			Action:   elktools_elasticsearch.EnableMlUpgradeMode,
		},
		{
			Name:     "disable-ml-upgrade",
			Usage:    "Disable upgrade mode on ML",
			Category: "Downtime",
			Action:   elktools_elasticsearch.DisableMlUpgradeMode,
		},
		{
			Name:     "stop-watcher-service",
			Usage:    "Stop watcher service",
			Category: "Downtime",
			Action:   elktools_elasticsearch.StopWatcherService,
		},
		{
			Name:     "start-watcher-service",
			Usage:    "Start watcher service",
			Category: "Downtime",
			Action:   elktools_elasticsearch.StartWatcherService,
		},
		{
			Name:     "stop-ilm-service",
			Usage:    "Stop ILM service",
			Category: "Downtime",
			Action:   elktools_elasticsearch.StopILMService,
		},
		{
			Name:     "start-ilm-service",
			Usage:    "Start ILM service",
			Category: "Downtime",
			Action:   elktools_elasticsearch.StartILMService,
		},
		{
			Name:     "stop-slm-service",
			Usage:    "Stop SLM service",
			Category: "Downtime",
			Action:   elktools_elasticsearch.StopSLMService,
		},
		{
			Name:     "start-slm-service",
			Usage:    "Start SLM service",
			Category: "Downtime",
			Action:   elktools_elasticsearch.StartSLMService,
		},
		{
			Name:     "export-data",
			Usage:    "Export data from query to file",
			Category: "Export",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "from",
					Usage: "From time to export data",
					Value: "now-24h",
				},
				&cli.StringFlag{
					Name:  "to",
					Usage: "To time to export data",
					Value: "now",
				},
				&cli.StringFlag{
					Name:  "date-field",
					Usage: "The date field to range over",
					Value: "@timestamp",
				},
				&cli.StringFlag{
					Name:  "index",
					Usage: "The index to export data",
					Value: "_all",
				},
				&cli.StringFlag{
					Name:  "query",
					Usage: "To query to export data",
				},
				&cli.StringSliceFlag{
					Name:  "fields",
					Usage: "Fields to extracts",
					Value: cli.NewStringSlice("log.original"),
				},
				&cli.StringFlag{
					Name:  "separator",
					Usage: "The separator to concatain field when extract multi fields",
					Value: "|",
				},
				&cli.StringFlag{
					Name:  "split-file-field",
					Usage: "The field to use to split data into multi files",
					Value: "host.name",
				},
				&cli.StringFlag{
					Name:  "path",
					Usage: "The root path to create extracted files",
				},
			},
			Action: elktools_elasticsearch.ExportDataToFiles,
		},
	}

	app.Before = func(c *cli.Context) error {

		if c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}

		if c.String("config") != "" {
			before := altsrc.InitInputSourceWithContext(app.Flags, altsrc.NewYamlSourceFromFlagFunc("config"))
			return before(c)
		}
		return nil
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(args)
	return err
}

func main() {
	err := run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
