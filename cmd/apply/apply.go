package main

import (
	"flag"
	"fmt"
    "log"
    "io/ioutil"
    "reflect"
    //"time"
    //"github.com/ghodss/yaml"
    //"k8s.io/apimachinery/pkg/api/errors"
    //"github.com/pkg/errors"
	apiv1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

var clusterFlag = flag.String("cluster", "", "Specify the name of the cluster")
var configFlag = flag.String("config", "", "Specify the k8s object config file")
/*func apply() error {


}*/


func init() {
    flag.StringVar(clusterFlag, "c", "", "Specify the name of the cluster")
    flag.StringVar(configFlag, "f", "", "Specify the k8s object config file")
}

func decodeFile(*string) {

    yamlFile, _ := ioutil.ReadFile("test-namespace.yaml")
    decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode([]byte(yamlFile), nil, nil)
    if err != nil {
	    log.Fatal(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
	}
    fmt.Printf("%++v\n\n", obj.GetObjectKind())
    fmt.Printf("%++v\n\n", obj)
    /*switch o := obj.(type) {
        case *apiv1.Namespace:
             _, err = clientset.CoreV1().Namespaces().Create(o)
        default:
            fmt.Printf("bar")
    }*/

}


func main() {

  //  processFile()
    home := homedir.HomeDir()
	var kubeconfig string = filepath.Join(home, ".kube", "config")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
    }
    /*nsSpec := &apiv1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "ns"}}
    _, err = clientset.CoreV1().Namespaces().Create(nsSpec)
    if err != nil {
        panic(err)
    }*/
    yamlFile, _ := ioutil.ReadFile(*configFlag)
    decode := scheme.Codecs.UniversalDeserializer().Decode
    obj, _, err := decode([]byte(yamlFile), nil, nil)
    if err != nil {
        log.Fatal(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
    }
    fmt.Printf("%++v\n\n", obj.GetObjectKind())
    fmt.Printf("%++v\n\n", obj)
    switch o := obj.(type) {
        case *apiv1.Namespace:
            fmt.Printf("Foo\n")
             _, err = clientset.CoreV1().Namespaces().Create(o)
        default:
            fmt.Printf("bar")
    }
    fmt.Println(reflect.TypeOf(obj))
}


