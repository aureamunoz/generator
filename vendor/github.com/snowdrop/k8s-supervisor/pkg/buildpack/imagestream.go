package buildpack

import (
	"github.com/ghodss/yaml"
	log "github.com/sirupsen/logrus"

	imagev1 "github.com/openshift/api/image/v1"
	imageclientsetv1 "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	restclient "k8s.io/client-go/rest"

	"github.com/snowdrop/k8s-supervisor/pkg/buildpack/types"
	"github.com/snowdrop/k8s-supervisor/pkg/common/oc"
	"strings"
)

var defaultImages = []types.Image{
	*CreateTypeImage(true, "dev-s2i", "latest", "quay.io/snowdrop/spring-boot-s2i", false),
	*CreateTypeImage(true, "copy-supervisord", "latest", "quay.io/snowdrop/supervisord", true),
}

func CreateDefaultImageStreams(config *restclient.Config, appConfig types.Application) {
	CreateImageStreamTemplate(config, appConfig, defaultImages)
}

func CreateImageStreamTemplate(config *restclient.Config, appConfig types.Application, images []types.Image) {
	imageClient := getImageClient(config)

	appCfg := appConfig
	for _, img := range images {

		appCfg.Image = img

		// first check that the image stream hasn't already been created
		if oc.Exists("imagestream", img.Name) {
			log.Infof("'%s' ImageStream already exists, skipping", img.Name)
		} else {
			// Parse ImageStream Template
			tName := strings.Join([]string{builderPath,"imagestream"},"/")
			var b = ParseTemplate(tName, appCfg)

			// Create ImageStream struct using the generated ImageStream string
			img := imagev1.ImageStream{}
			errYamlParsing := yaml.Unmarshal(b.Bytes(), &img)
			if errYamlParsing != nil {
				panic(errYamlParsing)
			}

			_, errImages := imageClient.ImageStreams(appConfig.Namespace).Create(&img)
			if errImages != nil {
				log.Fatalf("Unable to create ImageStream: %s", errImages.Error())
			}
		}
	}
}

func getImageClient(config *restclient.Config) *imageclientsetv1.ImageV1Client {
	imageClient, err := imageclientsetv1.NewForConfig(config)
	if err != nil {
		log.Fatal("Couldn't get ImageV1Client: %s", err)
	}
	return imageClient
}

func DeleteDefaultImageStreams(config *restclient.Config, appConfig types.Application) {
	for _, img := range defaultImages {
		// first check that the image stream hasn't already been created
		if oc.Exists("imagestream", img.Name) {
			client := getImageClient(config)
			err := client.ImageStreams(appConfig.Namespace).Delete(img.Name, deleteOptions)
			if err != nil {
				log.Fatalf("Unable to delete ImageStream: %s", img.Name)
			}
		}
	}
}

func CreateTypeImage(dockerImage bool, name string, tag string, repo string, annotationCmd bool) *types.Image {
	return &types.Image{
		DockerImage:    dockerImage,
		Name:           name,
		Repo:           repo,
		AnnotationCmds: annotationCmd,
		Tag:            tag,
	}
}
