package main

import (
	//"fmt"
	"context"
	"fmt"

	"github.com/ForgeCloud/saas/go/common/pkg/cluster"
	"github.com/ForgeCloud/saas/go/common/pkg/logging"
	"github.com/ForgeCloud/saas/go/services/customer/org-engine/pkg/keystoresecrets"
)

//idm keystore is fr-platform/secrets/idm/keystore.jceks
//userstore keystore is fr-platform/secrets/ds/
func main() {
	name := "idm_keystore_secret"
	ctx := context.Background()
	log := logging.Recorder
	gcpService := keystoresecrets.NewSecretsService(ctx, "fr-wq1sxjgtkmzm0npumdluxv4rv2p")
	safe, err := gcpService.CheckBackup(name)
	if err != nil {
		log.Infof("Error: %s", err)
	}
	if !safe {
		log.Infof("Safe: %t", safe)
		k8sClient, err := cluster.NewKubernetesClient()
		if err != nil {
			log.Infof("K8S Client error: %s", err)
		}
		idmKeystore, err := keystoresecrets.GetIDMKeystore(k8sClient)
		//log.Infof("IDM Keystore: %s", string(idmKeystore))
		if err != nil {
			log.Infof("Error getting idm secret: %s", err)
		}

		err = gcpService.BackupKeystore(name, idmKeystore)
		if err != nil {
			log.Infof("Create error: %s", err)
		}
	}
}
