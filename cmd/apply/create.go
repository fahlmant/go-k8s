package main

import (
	"flag"
	"fmt"
    "log"
    "io/ioutil"
    "reflect"
	"path/filepath"
    //"k8s.io/apimachinery/pkg/api/errors"
    //"github.com/pkg/errors"
	apiv1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
    rbacv1 "k8s.io/api/rbac/v1"
    //"github.com/hashicorp/vault/api"
)

var clusterFlag = flag.String("cluster", "", "Specify the name of the cluster")
var configFlag = flag.String("fileconfig", "", "Specify the k8s object config file")

func init() {
    flag.StringVar(clusterFlag, "c", "", "Specify the name of the cluster")
    flag.StringVar(configFlag, "f", "", "Specify the k8s object config file")
}

func loadKubeConfig() (*kubernetes.Clientset) {

    //Build the path to the kubeconfig
    home := homedir.HomeDir()
    var kubeconfig string = filepath.Join(home, ".kube", "config")

    //Use the config to create the client
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        panic(err)
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err)
    }

    return clientset;
}

func checkErr(err error) {

    if err != nil {
        log.Fatal(fmt.Sprintf("Error: %s", err))
    }
}

func main() {

    flag.Parse()

    clientset := loadKubeConfig()

    //Decode yaml file
    yamlFile, _ := ioutil.ReadFile(*configFlag)
    decode := scheme.Codecs.UniversalDeserializer().Decode
    obj, _, err := decode([]byte(yamlFile), nil, nil)
    checkErr(err)

    //
    switch o := obj.(type) {
        case *apiv1.Namespace:
            fmt.Println("Creating namespace...")
            _, err = clientset.CoreV1().Namespaces().Update(o)
            checkErr(err)
        case *apiv1.ServiceAccount:
            fmt.Println("Creating serviceaccount...")
            fmt.Println(o.Namespace)
            _, err = clientset.CoreV1().ServiceAccounts(o.Namespace).Create(o)
            checkErr(err)
        case *rbacv1.ClusterRole:
            fmt.Println("Creating ClusterRole...")
            _, err = clientset.RbacV1().ClusterRoles().Create(o)
            checkErr(err)
        case *rbacv1.ClusterRoleBinding:
            fmt.Println("Creating ClusterRoleBinding...")
            _, err = clientset.RbacV1().ClusterRoleBindings().Create(o)
            checkErr(err)
        default:
            fmt.Println("K8s object not currently handled")
            fmt.Println(reflect.TypeOf(o))
            fmt.Println(o)
    }
}
