package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "log"
    
   "gopkg.in/yaml.v3"
)

//part of select specific yaml file.
//reads specific information about yaml file. 
type YamlConfig struct {
     Hits int `yaml:"hits"`
    Time int `yaml:"time"`
    
}

//program to "modify" the code.
//follows the structure seen in var. But can have different data populating the fields.
type Policy struct {
    Control string `yaml:"control,omitempty"` 
    Id string `yaml:"id,omitempty"`
    Text string `yaml:"text,omitempty"`
    Checks string `yaml:"checks,omitempty"` 
    Group struct{
    	 Id string `yaml:"id,omitempty"` 
	 Text string `yaml:"text,omitempty"` 
    } `yaml:"group,omitempty"
}

func prcrdFields(	
    control string,
    id string,
    text string, 
    checks string,
    comment string) (*yaml.Node, error) {

    app := Policy{
        Control: control,
        Id: id,
	Text: text,
	Checks: checks, 
	Group: struct{
    	    Id string `yaml:"id,omitempty"`
	    Text string `yaml:"text,omitempty"` 
	} {id, text},
    }
    marshalledApp, err := yaml.Marshal(&app)
    if err != nil {
        return nil, err
    }

    node := yaml.Node{}
    if err := yaml.Unmarshal(marshalledApp, &node); err != nil {
        return nil, err
    }
    node.Content[0].HeadComment = comment
    return &node, nil
}

var prcrd = `
	apiVersion: policy.kubernetes.io/v1alpha1
	kind: PolicyReport
	metadata:
	
#  First app
  - name: app1
    kind: nodejs
    path: app1
    exec:
      platforms: k8s
      builder: test
`
)


func main() {

//modify yaml file
 yamlNode := yaml.Node{}
	
    err := yaml.Unmarshal([]byte(sourceYaml), &yamlNode)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    newApp, err := newApplicationNode("test", "5", "Kubernetes Policies",
         "5.1", "Service")
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    appIdx := -1
    for i, k := range yamlNode.Content[0].Content {
        if k.Value == "applications" {
            appIdx = i + 1
            break
        }
    }

yamlNode.Content[0].Content[appIdx].Content = append(
        yamlNode.Content[0].Content[appIdx].Content, newApp.Content[0])

    out, err := yaml.Marshal(&yamlNode)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(out))
    

//----------------

//Select specific yaml file
    fmt.Println("Parsing YAML file")

    var fileName string
    flag.StringVar(&fileName, "f", "", "YAML file to parse.")
    flag.Parse()

    if fileName == "" {
        fmt.Println("Please provide yaml file by using -f option")
        return
    }

    yamlFile, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Printf("Error reading YAML file: %s\n", err)
        return
    }

    var yamlConfig YamlConfig
    err = yaml.Unmarshal(yamlFile, &yamlConfig)
    if err != nil {
        fmt.Printf("Error parsing YAML file: %s\n", err)
    }

    fmt.Printf("Result: %v\n", yamlConfig)
}
