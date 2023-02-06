package controllers

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/rand"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/mgrote/personal-iot/api/v1alpha1"
)

var _ = Describe("Power outlet controller", func() {
	Context("Power outlet controller resource test", func() {

		testName := "test-outlet-controller-test" + rand.String(4)

		ctx := context.Background()

		ns := corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name:      testName,
				Namespace: testName,
			},
		}

		typedNs := types.NamespacedName{Name: testName, Namespace: testName}

		BeforeEach(func() {
			By("Create Namespace to perform test")
			err := k8sClient.Create(ctx, &ns)
			Expect(err).To(Not(HaveOccurred()))
		})

		AfterEach(func() {
			By("Delete Namespace to clean up test")
			_ = k8sClient.Delete(ctx, &ns)
		})

		It("should create power outlet resource in namespace", func() {

			By("Non existent resource 'Poweroutlet' is expected")
			powerOutlet := v1alpha1.Poweroutlet{}
			err := k8sClient.Get(ctx, typedNs, &powerOutlet)
			Expect(errors.IsNotFound(err)).To(BeTrue())

			powerOutletList := &v1alpha1.PoweroutletList{}
			//err = k8sClient.List(ctx, powerOutletList, &client.ListOptions{Namespace: testName})
			Expect(k8sClient.List(ctx, powerOutletList, &client.ListOptions{Namespace: testName})).Should(Succeed())

			By("Create a power outlet resource")
			powerOutlet.Namespace = testName
			// TODO talk:
			// first: leave outlet name empty
			// second: set outlet to wrong name powerOutlet.Name = "Test-Power-Outlet"
			powerOutlet.Name = "poweroutlet"
			err = k8sClient.Create(ctx, &powerOutlet)
			Expect(err).ToNot(HaveOccurred())

			By("Power outlet object should be found.")
			Expect(k8sClient.List(ctx, powerOutletList, &client.ListOptions{Namespace: testName})).Should(Succeed())
			Expect(len(powerOutletList.Items)).To(BeIdenticalTo(1))

			// TODO Questions to explain: How is the test local kubeapi reached, how is the test local etcd inspected?

			objKey := client.ObjectKeyFromObject(&powerOutlet)
			err = k8sClient.Get(ctx, objKey, &powerOutlet)
			Expect(err).ShouldNot(HaveOccurred())

			By("Delete the created power outlet resource")
			err = k8sClient.Delete(ctx, &powerOutlet)
			Expect(err).ToNot(HaveOccurred())
			err = k8sClient.Get(ctx, typedNs, &powerOutlet)
			Expect(errors.IsNotFound(err)).To(BeTrue())
		})

	})
})
