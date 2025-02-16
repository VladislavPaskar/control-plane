package provisioning

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/broker"
	provisionerAutomock "github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/provisioner/automock"
	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/ptr"
	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/storage"
	"github.com/kyma-project/control-plane/components/provisioner/pkg/gqlschema"
	"github.com/pivotal-cf/brokerapi/v8/domain"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateRuntimeWithoutKyma_Run(t *testing.T) {
	// given
	log := logrus.New()
	memoryStorage := storage.NewMemoryStorage()

	operation := fixOperationCreateRuntime(t, broker.GCPPlanID, "europe-west3")
	operation.ShootDomain = "kyma.org"
	err := memoryStorage.Operations().InsertProvisioningOperation(operation)
	assert.NoError(t, err)

	err = memoryStorage.Instances().Insert(fixInstance())
	assert.NoError(t, err)

	administrator := ""
	provisionerInput := gqlschema.ProvisionRuntimeInput{
		RuntimeInput: &gqlschema.RuntimeInput{
			Name:        "dummy",
			Description: nil,
			Labels: &gqlschema.Labels{
				"broker_instance_id":   instanceID,
				"global_subaccount_id": subAccountID,
				"operator_grafanaUrl":  "https://grafana.kyma.org",
			},
		},
		ClusterConfig: &gqlschema.ClusterConfigInput{
			GardenerConfig: &gqlschema.GardenerConfigInput{
				Name:                                shootName,
				KubernetesVersion:                   k8sVersion,
				DiskType:                            ptr.String("pd-standard"),
				VolumeSizeGb:                        ptr.Integer(30),
				MachineType:                         "n2-standard-8",
				Region:                              "europe-west3",
				Provider:                            "gcp",
				Purpose:                             &shootPurpose,
				LicenceType:                         nil,
				WorkerCidr:                          "10.250.0.0/19",
				AutoScalerMin:                       3,
				AutoScalerMax:                       4,
				MaxSurge:                            4,
				MaxUnavailable:                      1,
				TargetSecret:                        "",
				EnableKubernetesVersionAutoUpdate:   ptr.Bool(autoUpdateKubernetesVersion),
				EnableMachineImageVersionAutoUpdate: ptr.Bool(autoUpdateMachineImageVersion),
				ProviderSpecificConfig: &gqlschema.ProviderSpecificInput{
					GcpConfig: &gqlschema.GCPProviderConfigInput{
						Zones: []string{"europe-west3-b", "europe-west3-c"},
					},
				},
				Seed: nil,
				OidcConfig: &gqlschema.OIDCConfigInput{
					ClientID:       "9bd05ed7-a930-44e6-8c79-e6defeb7dec9",
					GroupsClaim:    "groups",
					IssuerURL:      "https://kymatest.accounts400.ondemand.com",
					SigningAlgs:    []string{"RS256"},
					UsernameClaim:  "sub",
					UsernamePrefix: "-",
				},
			},
			Administrators: []string{administrator},
		},
		KymaConfig: nil,
	}

	provisionerClient := &provisionerAutomock.Client{}
	provisionerClient.On("ProvisionRuntime", globalAccountID, subAccountID, mock.MatchedBy(
		func(input gqlschema.ProvisionRuntimeInput) bool {
			return reflect.DeepEqual(input.RuntimeInput.Labels, provisionerInput.RuntimeInput.Labels) &&
				input.KymaConfig == nil && reflect.DeepEqual(input.ClusterConfig, provisionerInput.ClusterConfig)
		},
	)).Return(gqlschema.OperationStatus{
		ID:        ptr.String(provisionerOperationID),
		Operation: "",
		State:     "",
		Message:   nil,
		RuntimeID: ptr.String(runtimeID),
	}, nil)

	step := NewCreateRuntimeWithoutKymaStep(memoryStorage.Operations(), memoryStorage.RuntimeStates(), memoryStorage.Instances(), provisionerClient)

	// when
	entry := log.WithFields(logrus.Fields{"step": "TEST"})
	operation, repeat, err := step.Run(operation, entry)

	// then
	assert.NoError(t, err)
	assert.Zero(t, repeat)
	assert.Equal(t, provisionerOperationID, operation.ProvisionerOperationID)

	instance, err := memoryStorage.Instances().GetByID(operation.InstanceID)
	assert.NoError(t, err)
	assert.Equal(t, instance.RuntimeID, runtimeID)
}

func TestCreateRuntimeWithoutKymaStep_RunWithBadRequestError(t *testing.T) {
	// given
	log := logrus.New()
	memoryStorage := storage.NewMemoryStorage()

	operation := fixOperationCreateRuntime(t, broker.AzurePlanID, "westeurope")
	err := memoryStorage.Operations().InsertProvisioningOperation(operation)
	assert.NoError(t, err)

	err = memoryStorage.Instances().Insert(fixInstance())
	assert.NoError(t, err)

	provisionerClient := &provisionerAutomock.Client{}
	provisionerClient.On("ProvisionRuntime", globalAccountID, subAccountID, mock.Anything).Return(gqlschema.OperationStatus{}, fmt.Errorf("some permanent error"))

	step := NewCreateRuntimeWithoutKymaStep(memoryStorage.Operations(), memoryStorage.RuntimeStates(), memoryStorage.Instances(), provisionerClient)

	// when
	entry := log.WithFields(logrus.Fields{"step": "TEST"})
	operation, _, err = step.Run(operation, entry)

	// then
	assert.Equal(t, domain.Failed, operation.State)
}
