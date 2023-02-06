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

var _ = Describe("Power strip controller", func() {
	Context("Power strip controller resource test", func() {

		testName := "test-powerstrip-controller-test" + rand.String(4)

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

		It("should create power strip resource in namespace", func() {

			By("Non existent resource 'Powerstrip' is expected")
			powerStrip := v1alpha1.Powerstrip{}
			err := k8sClient.Get(ctx, typedNs, &powerStrip)
			Expect(errors.IsNotFound(err)).To(BeTrue())

			powerStripList := &v1alpha1.PowerstripList{}
			//err = k8sClient.List(ctx, powerStripList, &client.ListOptions{Namespace: testName})
			Expect(k8sClient.List(ctx, powerStripList, &client.ListOptions{Namespace: testName})).Should(Succeed())

			By("Create a power strip resource")
			powerStrip.Namespace = testName
			// TODO talk:
			// first: leave powerstrip name empty
			// second: set powerstrip to wrong name powerStrip.Name = "Test-Power-Strip"
			powerStrip.Name = "powerstrip"
			err = k8sClient.Create(ctx, &powerStrip)
			Expect(err).ToNot(HaveOccurred())

			By("Power strip object should be found.")
			Expect(k8sClient.List(ctx, powerStripList, &client.ListOptions{Namespace: testName})).Should(Succeed())
			Expect(len(powerStripList.Items)).To(BeIdenticalTo(1))

			// TODO Questions to explain: How is the test local kubeapi reached, how is the test local etcd inspected?

			objKey := client.ObjectKeyFromObject(&powerStrip)
			err = k8sClient.Get(ctx, objKey, &powerStrip)
			Expect(err).ShouldNot(HaveOccurred())

			By("Delete the created power strip resource")
			err = k8sClient.Delete(ctx, &powerStrip)
			Expect(err).ToNot(HaveOccurred())
			err = k8sClient.Get(ctx, typedNs, &powerStrip)
			Expect(errors.IsNotFound(err)).To(BeTrue())
		})

	})
})
