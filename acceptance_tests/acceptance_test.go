package acceptance_test

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/helpers"
	"github.com/cloudfoundry-incubator/cf-test-helpers/workflowhelpers"
	"github.com/cloudfoundry/cf-acceptance-tests/helpers/config"
)

const (
	DB_CREATE_TIMEOUT = 30 * time.Minute
)

var (
	testConfig  config.CatsConfig
	httpClient  *http.Client
	testContext *workflowhelpers.ReproducibleTestSuiteSetup
)

func TestSuite(t *testing.T) {
	var err error
	RegisterFailHandler(Fail)

	testConfig, err = config.NewCatsConfig(os.Getenv("CONFIG"))
	if err != nil {
		t.Fatal(err)
	}

	httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: testConfig.GetSkipSSLValidation()},
		},
	}

	testContext = workflowhelpers.NewTestSuiteSetup(testConfig)

	BeforeSuite(func() {
		testContext.Setup()

		// FIXME this should be removed once these services are generally available.
		org := testContext.GetOrganizationName()
		workflowhelpers.AsUser(testContext.AdminUserContext(), testContext.ShortTimeout(), func() {
			enableServiceAccess := cf.Cf("enable-service-access", "aws-s3-bucket", "-o", org).Wait(testConfig.DefaultTimeoutDuration())
			Expect(enableServiceAccess).To(Exit(0))
			Expect(enableServiceAccess).To(Say("OK"))
		})
	})

	AfterSuite(func() {
		testContext.Teardown()
	})

	componentName := "Custom-Acceptance-Tests"
	if testConfig.GetArtifactsDirectory() != "" {
		helpers.EnableCFTrace(testConfig, componentName)
	}

	RunSpecs(t, componentName)
}

// quietCf is an equivelent of cf.Cf that doesn't send the output to
// GinkgoWriter. Used when you don't want the output, even in verbose mode (eg
// when polling the API)
func quietCf(program string, args ...string) *Session {
	command, err := Start(exec.Command(program, args...), nil, nil)
	Expect(err).NotTo(HaveOccurred())
	return command
}

func pollForAuditEventInDb(auditEvent cfclient.Event) {
	fmt.Fprint(GinkgoWriter, "Polling for service creation to complete")
	Eventually(func() *Buffer {
		fmt.Fprint(GinkgoWriter, ".")
		command := quietCf("cf", "service", dbInstanceName).Wait(testConfig.DefaultTimeoutDuration())
		Expect(command).To(Exit(0), fmt.Sprint("Error calling cf service creation phase: ", string(command.Out.Contents())))
		return command.Out
	}, DB_CREATE_TIMEOUT, 15*time.Second).Should(Say("create succeeded"))
	fmt.Fprint(GinkgoWriter, "done\n")
}
