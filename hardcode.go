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

var prcrd = `
	apiVersion: policy.kubernetes.io/v1alpha1
	kind: PolicyReport
	metadata:
`

//program to "modify" the code.
//follows the structure seen in var. But can have different data populating the fields.
type Policy struct {
    Name string `yaml:"name,omitempty"` 
    Labels string `yaml:"labels,omitempty"`
	Annotations struct{
		A_name string `yaml:"a_name,omitempty"`
		Category string `yaml:"category,omitempty"`
		P_version string `yaml:"p_version,omitempty"`
	} `yaml:"annotations,omitempty"`
	Summary struct{
		Pass int `yaml:"pass,omitempty"`
		Fail int `yaml:"fail,omitempty"`
		Warn int `yaml:"warn,omitempty"`
		Info int `yaml:"info,omitempty"`
		Error int `yaml:"eror,omitempty"`
		Skip int `yaml:"skip,omitempty"`
	} `yaml:"summary,omitempty"`	
}

func prcrdFields(	
    name string,
    labels string,
    a_name string, 
    category string,
    p_version string, 
    pass int,
    fail int,
    warn int,
    info int,
    eror int,
    skip int
    ) (*yaml.Node, error) {

    app := Policy{
        Name: name,
        Labels: labels,
	Annotation: struct {
	     A_name string `yaml:"a_name,omitempty"`
	     Category string `yaml:"category,omitempty"`
	     P_version string `yaml:"p_version,omitempty"`    
	}{a_name, category, p_version}, 
	Summary struct{
	     Pass int `yaml:"pass,omitempty"`
	     Fail int `yaml:"fail,omitempty"`
	     Warn int `yaml:"warn,omitempty"`
	     Info int `yaml:"info,omitempty"`
	     Error int `yaml:"eror,omitempty"`
	     Skip int `yaml:"skip,omitempty"`
       }{pass, fail, warn, info, eror, skip},
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



func main() {
 a := 5
	
//modify yaml file
 yamlNode := yaml.Node{}
	
    err := yaml.Unmarshal([]byte(sourceYaml), &yamlNode)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    newApp, err := newApplicationNode("test", a, "Kubernetes Policies",
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
