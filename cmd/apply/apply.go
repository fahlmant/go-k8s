package main

import (
	"flag"
	//"fmt"
    //"time"
    //"github.com/ghodss/yaml"
    //"k8s.io/apimachinery/pkg/api/errors"
    //"github.com/pkg/errors"
	//appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func main() {

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
    nsSpec := &apiv1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "ns"}}
    _, err = clientset.CoreV1().Namespaces().Create(nsSpec)
    if err != nil {
        panic(err)
    }
}


