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

var _ = Describe("Location integration", func() {
	Context("Location integration resource test", func() {

		testName := "test-location-test" + rand.String(4)

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

		It("should create location in namespace", func() {
			By("Non existent resource 'Location' is expected.")
			locationOne := v1alpha1.Location{}
			err := k8sClient.Get(ctx, typedNs, &locationOne)
			Expect(errors.IsNotFound(err)).To(BeTrue())

			locationList := v1alpha1.LocationList{}
			Expect(k8sClient.List(ctx, &locationList, &client.ListOptions{Namespace: testName})).Should(Succeed())

			//By("Create poweroutlets and powerstrip to add to location.")
			//
			//powerOutletOne := v1alpha1.Poweroutlet{
			//	Spec: v1alpha1.PoweroutletSpec{
			//		Switch: "off",
			//	},
			//}
			//powerOutletOne.Name = "light-one"
			//powerOutletOne.Namespace = testName
			//err = k8sClient.Create(ctx, &powerOutletOne)
			//Expect(err).ToNot(HaveOccurred())
			//
			//powerOutletTwo := v1alpha1.Poweroutlet{
			//	Spec: v1alpha1.PoweroutletSpec{
			//		Switch: "off",
			//	},
			//}
			//powerOutletTwo.Name = "light-two"
			//powerOutletTwo.Namespace = testName
			//err = k8sClient.Create(ctx, &powerOutletTwo)
			//Expect(err).ToNot(HaveOccurred())
			//
			//powerOutletThree := v1alpha1.Poweroutlet{
			//	Spec: v1alpha1.PoweroutletSpec{
			//		Switch: "on",
			//	},
			//}
			//powerOutletThree.Name = "light-three"
			//powerOutletThree.Namespace = testName
			//err = k8sClient.Create(ctx, &powerOutletThree)
			//Expect(err).ToNot(HaveOccurred())
			//
			//powerOutletList := v1alpha1.PoweroutletList{}
			//Expect(k8sClient.List(ctx, &powerOutletList, &client.ListOptions{Namespace: testName})).Should(Succeed())
			//Expect(len(powerOutletList.Items)).To(BeIdenticalTo(3))
			//
			//powerStrip := v1alpha1.Powerstrip{
			//	Spec: v1alpha1.PowerstripSpec{
			//		Outlets: []v1alpha1.Poweroutlet{
			//			powerOutletOne,
			//			powerOutletTwo,
			//			powerOutletThree,
			//		},
			//	},
			//}
			//powerStrip.Name = "light-strip"
			//powerStrip.Namespace = testName
			//err = k8sClient.Create(ctx, &powerStrip)
			//Expect(err).ToNot(HaveOccurred())
			//
			//// No more power outlets should occur.
			//Expect(k8sClient.List(ctx, &powerOutletList, &client.ListOptions{Namespace: testName})).Should(Succeed())
			//Expect(len(powerOutletList.Items)).To(BeIdenticalTo(3))
			//
			//locationName := "room-one"
			//locationOne.Name = locationName
			//locationOne.Namespace = testName
			//locationOne.Spec = v1alpha1.LocationSpec{
			//	Powerstrips: []v1alpha1.Powerstrip{
			//		powerStrip,
			//	},
			//}
			//err = k8sClient.Create(ctx, &locationOne)
			//Expect(err).ToNot(HaveOccurred())
			//
			//By("Reread location to check if nested objects are loaded.")
			//
			//locationList = v1alpha1.LocationList{}
			//Expect(k8sClient.List(ctx, &locationList, &client.ListOptions{Namespace: testName})).Should(Succeed())
			//
			//locationReread := v1alpha1.Location{}
			//locationReread.Name = locationName
			//locationReread.Namespace = testName
			//objKey := client.ObjectKeyFromObject(&locationReread)
			//err = k8sClient.Get(ctx, objKey, &locationReread)
			//Expect(err).To(BeNil())
			//Expect(locationReread).ToNot(BeNil())
			//
			//Expect(len(locationReread.Spec.Powerstrips)).To(BeIdenticalTo(1))
			//rereadPowerStrip := locationReread.Spec.Powerstrips[0]
			//Expect(len(rereadPowerStrip.Spec.Outlets)).To(BeIdenticalTo(3))

		})
	})
})
