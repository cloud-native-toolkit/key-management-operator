/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	keymanagementv1 "github.com/ibmgaragecloud/key-management-operator/api/v1"
)

// SecretTemplateReconciler reconciles a SecretTemplate object
type SecretTemplateReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=keymanagement.ibm,resources=secrettemplates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=keymanagement.ibm,resources=secrettemplates/status,verbs=get;update;patch

func (r *SecretTemplateReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("secrettemplate", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *SecretTemplateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&keymanagementv1.SecretTemplate{}).
		Complete(r)
}
