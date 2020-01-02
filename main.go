package main

import (
	"os"
	"sort"

	elktools_elasticsearch "github.com/disaster37/elktools/v7/elasticsearch"
	elktools_kibana "github.com/disaster37/elktools/v7/kibana"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
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
	app.Version = "develop"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "Load configuration from `FILE`",
		},
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "url",
			Usage:  "The elasticsearch or kibana URL",
			EnvVar: "ELKTOOLS_URL",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "user",
			Usage:  "The  user",
			EnvVar: "ELKTOOLS_USER",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "password",
			Usage:  "The password",
			EnvVar: "ELKTOOLS_PASSWORD",
		}),
		cli.BoolFlag{
			Name:  "self-signed-certificate",
			Usage: "Disable the TLS certificate check",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Display debug output",
		},
	}
	app.Commands = []cli.Command{
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
				cli.StringFlag{
					Name:  "lifecycle-policy-id",
					Usage: "The lifecycle policy name",
				},
				cli.StringFlag{
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
				cli.StringFlag{
					Name:  "lifecycle-policy-id",
					Usage: "The lifecycle policy name",
				},
				cli.StringFlag{
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
				cli.StringFlag{
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
				cli.StringFlag{
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
				cli.StringFlag{
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
				cli.StringFlag{
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
				cli.StringFlag{
					Name:  "indice-template-id",
					Usage: "The indice template name",
				},
				cli.StringFlag{
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
				cli.StringFlag{
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
				cli.StringFlag{
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
				cli.StringFlag{
					Name:  "indice-template-id",
					Usage: "The indice template ID to save",
				},
				cli.StringFlag{
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
				cli.StringFlag{
					Name:  "indice-name",
					Usage: "The indice name to create",
				},
				cli.StringFlag{
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
				cli.StringFlag{
					Name:  "file-path",
					Usage: "The file path to export all dashboards",
				},
				cli.StringFlag{
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
				cli.StringFlag{
					Name:  "file-path",
					Usage: "The file path to load all dashboards",
				},
				cli.StringFlag{
					Name:  "user-space",
					Usage: "The Kibana user space where write dashboards",
					Value: "default",
				},
			},
			Action: elktools_kibana.ImportDashboards,
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
