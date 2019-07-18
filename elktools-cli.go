package main

import (
	"github.com/disaster37/elktools/elasticsearch"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/altsrc"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {

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
			Name:        "elasticsearch-url",
			Usage:       "The elasticsearch URL",
			EnvVar:      "ELASTICSEARCH_URL",
			Destination: &elktools_elasticsearch.ElasticsearchUrl,
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:        "elasticsearch-user",
			Usage:       "The Elasticsearch user",
			EnvVar:      "ELASTICSEARCH_USER",
			Destination: &elktools_elasticsearch.User,
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:        "elasticsearch-password",
			Usage:       "The Elasticsearch password",
			EnvVar:      "ELASTICSEARCH_PASSWORD",
			Destination: &elktools_elasticsearch.Password,
		}),
		cli.BoolFlag{
			Name:        "self-signed-certificate",
			Usage:       "Disable the TLS certificate check",
			Destination: &elktools_elasticsearch.DisableVerifySSL,
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Display debug output",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "create-or-update-lifecycle-policy",
			Usage: "Create or update lifecycle policy",
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
			Name:  "save-lifecycle-policy",
			Usage: "Save lifecycle policy on file",
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
			Name:  "delete-lifecycle-policy",
			Usage: "Delete lifecycle policy",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "lifecycle-policy-id",
					Usage: "The lifecycle policy name",
				},
			},
			Action: elktools_elasticsearch.DeleteILMPolicy,
		},
		{
			Name:  "save-all-lifecycle-policies",
			Usage: "Save all lifecycle policies on folder provided",
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
			Name:  "create-or-update-all-lifecycle-policies",
			Usage: "Create or update all lifecycle policies from folder provided",
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
			Name:  "get-lifecycle-policy-status",
			Usage: "Get the lifecycle policy on Elasticsearch index",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "elasticsearch-index",
					Usage: "The Elasticsearch index where you should to get lifecycle policy status",
				},
			},
			Action: elktools_elasticsearch.GetStatusILMPolicy,
		},
		{
			Name:  "create-or-update-indice-template",
			Usage: "Create or update indice template",
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
			Name:  "create-or-update-all-indice-templates",
			Usage: "Create or update all indice templates",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "indice-template-path",
					Usage: "The indice templates base path",
				},
			},
			Action: elktools_elasticsearch.CreateAllIndiceTemplates,
		},
		{
			Name:  "delete-indice-template",
			Usage: "Delete indice template",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "indice-template-id",
					Usage: "The indice template ID to delete",
				},
			},
			Action: elktools_elasticsearch.DeleteIndiceTemplate,
		},
		{
			Name:  "save-indice-template",
			Usage: "Save indice template on file",
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
			Name:  "create-indice",
			Usage: "Create new indice with settings",
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
