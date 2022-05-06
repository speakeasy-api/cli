package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"

	"github.com/speakeasy-api/parser/services/parser"

	openapi "github.com/getkin/kin-openapi/openapi3" //for reading OpenAPI schemas
	"github.com/speakeasy-api/parser/apipackage"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

const (
	searchDirFlag        = "dir"
	excludeFlag          = "exclude"
	generalInfoFlag      = "generalInfo"
	propertyStrategyFlag = "propertyStrategy"
	outputFlag           = "output"
	outputTypesFlag      = "outputTypes"
	parseVendorFlag      = "parseVendor"
	parseDependencyFlag  = "parseDependency"
	markdownFilesFlag    = "markdownFiles"
	codeExampleFilesFlag = "codeExampleFiles"
	parseInternalFlag    = "parseInternal"
	generatedTimeFlag    = "generatedTime"
	parseDepthFlag       = "parseDepth"
	instanceNameFlag     = "instanceName"
	overridesFileFlag    = "overridesFile"
)

var initFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    generalInfoFlag,
		Aliases: []string{"g"},
		Value:   "main.go",
		Usage:   "Go file path in which 'OpenAPI general API Info' is written",
	},
	&cli.StringFlag{
		Name:    searchDirFlag,
		Aliases: []string{"d"},
		Value:   "./",
		Usage:   "Directories you want to parse,comma separated and general-info file must be in the first one",
	},
	&cli.StringFlag{
		Name:  excludeFlag,
		Usage: "Exclude directories and files when searching, comma separated",
	},
	&cli.StringFlag{
		Name:    propertyStrategyFlag,
		Aliases: []string{"p"},
		Value:   parser.CamelCase,
		Usage:   "Property Naming Strategy like " + parser.SnakeCase + "," + parser.CamelCase + "," + parser.PascalCase,
	},
	&cli.StringFlag{
		Name:    outputFlag,
		Aliases: []string{"o"},
		Value:   "./docs",
		Usage:   "Output directory for all the generated files(opeanapi.json, opeanapi.yaml)",
	},
	&cli.StringFlag{
		Name:    outputTypesFlag,
		Aliases: []string{"ot"},
		Value:   "go,json,yaml",
		Usage:   "Output types of generated files (opeanapi.json, opeanapi.yaml) like go,json,yaml",
	},
	&cli.BoolFlag{
		Name:  parseVendorFlag,
		Usage: "Parse go files in 'vendor' folder, disabled by default",
	},
	&cli.BoolFlag{
		Name:    parseDependencyFlag,
		Aliases: []string{"pd"},
		Usage:   "Parse go files inside dependency folder, disabled by default",
	},
	&cli.StringFlag{
		Name:    markdownFilesFlag,
		Aliases: []string{"md"},
		Value:   "",
		Usage:   "Parse folder containing markdown files to use as description, disabled by default",
	},
	&cli.StringFlag{
		Name:    codeExampleFilesFlag,
		Aliases: []string{"cef"},
		Value:   "",
		Usage:   "Parse folder containing code example files to use for the x-codeSamples extension, disabled by default",
	},
	&cli.BoolFlag{
		Name:  parseInternalFlag,
		Usage: "Parse go files in internal packages, disabled by default",
	},
	&cli.BoolFlag{
		Name:  generatedTimeFlag,
		Usage: "Generate timestamp at the top of docs.go, disabled by default",
	},
	&cli.IntFlag{
		Name:  parseDepthFlag,
		Value: 100,
		Usage: "Dependency parse depth",
	},
	&cli.StringFlag{
		Name:  instanceNameFlag,
		Value: "",
		Usage: "This parameter can be used to name different schema(openapi) document instances. It is optional.",
	},
	&cli.StringFlag{
		Name:  overridesFileFlag,
		Value: apipackage.DefaultOverridesFile,
		Usage: "File to read global type overrides from.",
	},
}

func initAction(c *cli.Context) error {
	strategy := c.String(propertyStrategyFlag)

	switch strategy {
	case parser.CamelCase, parser.SnakeCase, parser.PascalCase:
	default:
		return fmt.Errorf("not supported %s propertyStrategy", strategy)
	}

	outputTypes := strings.Split(c.String(outputTypesFlag), ",")
	if len(outputTypes) == 0 {
		return fmt.Errorf("no output types specified")
	}

	return apipackage.New().Build(&apipackage.Config{
		SearchDir:           c.String(searchDirFlag),
		Excludes:            c.String(excludeFlag),
		MainAPIFile:         c.String(generalInfoFlag),
		PropNamingStrategy:  strategy,
		OutputDir:           c.String(outputFlag),
		OutputTypes:         outputTypes,
		ParseVendor:         c.Bool(parseVendorFlag),
		ParseDependency:     c.Bool(parseDependencyFlag),
		MarkdownFilesDir:    c.String(markdownFilesFlag),
		ParseInternal:       c.Bool(parseInternalFlag),
		GeneratedTime:       c.Bool(generatedTimeFlag),
		CodeExampleFilesDir: c.String(codeExampleFilesFlag),
		ParseDepth:          c.Int(parseDepthFlag),
		InstanceName:        c.String(instanceNameFlag),
		OverridesFile:       c.String(overridesFileFlag),
	})
}

func exitNormally(diffEmpty bool) {
	if false && !diffEmpty {
		os.Exit(1)
	}
	os.Exit(0)
}

func printYAML(output interface{}) error {
	if reflect.ValueOf(output).IsNil() {
		return nil
	}

	bytes, err := yaml.Marshal(output)
	if err != nil {
		return err
	}
	fmt.Printf("%s", bytes)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Version = "v0.1.0-alpha"
	app.Usage = "Automatically track the state of your API, Generate artifacts like OpenAPI schemas and more."
	app.Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create initial state and default schema artifacts",
			Action:  initAction,
			Flags:   initFlags,
		},
		{
			Name:    "fmt",
			Aliases: []string{"f"},
			Usage:   "format Annotations",
			Action: func(c *cli.Context) error {
				searchDir := c.String(searchDirFlag)
				excludeDir := c.String(excludeFlag)
				mainFile := c.String(generalInfoFlag)

				return parser.NewFormat().Build(&parser.Config{
					SearchDir: searchDir,
					Excludes:  excludeDir,
					MainFile:  mainFile,
				})
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    searchDirFlag,
					Aliases: []string{"d"},
					Value:   "./",
					Usage:   "Directories you want to parse,comma separated and general-info file must be in the first one",
				},
				&cli.StringFlag{
					Name:  excludeFlag,
					Usage: "Exclude directories and files when searching, comma separated",
				},
				&cli.StringFlag{
					Name:    generalInfoFlag,
					Aliases: []string{"g"},
					Value:   "main.go",
					Usage:   "Go file path in which 'openapi general API Info' is written",
				},
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update API version",
			Action: func(c *cli.Context) error {
				// schemaDir := c.String(schemaDirFlag)
				// schemaDir := "/Users/HenrySwaffield/code/parser/cmd/api/docs/openapi.yaml" //json also there...

				doc, err := openapi.NewLoader().LoadFromFile("./petStoreSchema.json")

				if err != nil {
					log.Fatal(err)
				}

				//doc.Components.Responses

				log.Print(doc)

				log.Println(doc.Paths) //paths.

				//messing around with k8:

				var kubeconfig *string
				if home := homedir.HomeDir(); home != "" {
					kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
				} else {
					kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
				}
				flag.Parse()

				config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
				if err != nil {
					panic(err)
				}
				clientset, err := kubernetes.NewForConfig(config)
				if err != nil {
					panic(err)
				}

				deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

				list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})

				log.Println(list) // this is a hefty chunk of text, but it shows my toy services, which is good.

				/*
					prints:
						[]Deployment{Deployment{ObjectMeta:{toy-go-server  default  7e44a0c9-6d77-4231-afcf-3f113c15f9dc 2110548 11 2022-05-05 15:26:44 -0700 PDT <nil> <nil> map[app:toy-go-service component:server] map[deployment.kubernetes.io/revision:5 kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{},"labels":{"app":"toy-go-service","component":"server"},"name":"toy-go-server","namespace":"default"},"spec":{"replicas":1,"selector":{"matchLabels":{"app":"toy-go-service","component":"server"}},"template":{"metadata":{"labels":{"app":"toy-go-service","component":"server"}}, ... and so on
				*/

				ingressClient := clientset.NetworkingV1().Ingresses("default")

				log.Println(ingressClient.List(context.TODO(), metav1.ListOptions{}))

				/*
									prints:
										[]Ingress{Ingress{ObjectMeta:{toy-go-service  default  15e20126-6bac-4912-89bd-7a0db04c2b78 2111037 2 2022-05-05 18:25:37 -0700 PDT <nil> <nil> map[] map[ingress.kubernetes.io/backends:{"k8s1-c3a86ceb-default-toy-go-service-8080-c12d7085":"HEALTHY","k8s1-c3a86ceb-default-toy-go-service-v3-8080-8ff93136":"HEALTHY"} ingress.kubernetes.io/forwarding-rule:k8s2-fr-ct54ecng-default-toy-go-service-k1mkdcr6 ingress.kubernetes.io/target-proxy:k8s2-tp-ct54ecng-default-toy-go-service-k1mkdcr6 ingress.kubernetes.io/url-map:k8s2-um-ct54ecng-default-toy-go-service-k1mkdcr6 kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"networking.k8s.io/v1","kind":"Ingress","metadata":{"annotations":{},"name":"toy-go-service","namespace":"default"},"spec":{"defaultBackend":{"service":{"name":"toy-go-service","port":{"number":8080}}},"rules":[{"http":{"paths":[{"backend":{"service":{"name":"toy-go-service-v3","port":{"number":8080}}},"path":"/pet","pathType":"ImplementationSpecific"},{"backend":{"service":{"name":"toy-go-service","port":{"number":8080}}},"path":"/bike","pathType":"ImplementationSpecific"}]}}]},"status":{"loadBalancer":{}}}
					] [] [networking.gke.io/ingress-finalizer-V2]  [{kubectl-client-side-apply Update networking.k8s.io/v1 2022-05-05 18:25:37 -0700 PDT FieldsV1 {"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}}},"f:spec":{"f:defaultBackend":{".":{},"f:service":{".":{},"f:name":{},"f:port":{".":{},"f:number":{}}}},"f:rules":{}}} } {glbc Update networking.k8s.io/v1 2022-05-05 18:26:59 -0700 PDT FieldsV1 {"f:metadata":{"f:annotations":{"f:ingress.kubernetes.io/backends":{},"f:ingress.kubernetes.io/forwarding-rule":{},"f:ingress.kubernetes.io/target-proxy":{},"f:ingress.kubernetes.io/url-map":{}},"f:finalizers":{".":{},"v:\"networking.gke.io/ingress-finalizer-V2\"":{}}},"f:status":{"f:loadBalancer":{"f:ingress":{}}}} status}]},Spec:IngressSpec{DefaultBackend:&IngressBackend{Resource:nil,Service:&IngressServiceBackend{Name:toy-go-service,Port:ServiceBackendPort{Name:,Number:8080,},},},TLS:[]IngressTLS{},Rules:[]IngressRule{IngressRule{Host:,IngressRuleValue:IngressRuleValue{HTTP:&HTTPIngressRuleValue{Paths:[]HTTPIngressPath{HTTPIngressPath{Path:/pet,Backend:IngressBackend{Resource:nil,Service:&IngressServiceBackend{Name:toy-go-service-v3,Port:ServiceBackendPort{Name:,Number:8080,},},},PathType:*ImplementationSpecific,},HTTPIngressPath{Path:/bike,Backend:IngressBackend{Resource:nil,Service:&IngressServiceBackend{Name:toy-go-service,Port:ServiceBackendPort{Name:,Number:8080,},},},PathType:*ImplementationSpecific,},},},},},},IngressClassName:nil,},Status:IngressStatus{LoadBalancer:{[{34.149.66.48  []}]},},},},} <nil> ...
				*/

				// next steps:
				// loop through paths ^ and ensure that all those paths are added to the ingress that corresponds to this
				// package
				// also have all of those paths route to the same backend service

				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
