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
			Name:        "debug",
			Usage:       "Display debug output",
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


