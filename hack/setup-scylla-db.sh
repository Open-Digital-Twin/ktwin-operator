# Install operator helm dependencies
helm repo add scylla https://scylla-operator-charts.storage.googleapis.com/stable
helm repo update

# Configure cert-manager
kubectl apply -f https://raw.githubusercontent.com/scylladb/scylla-operator/v1.9/examples/common/cert-manager.yaml
kubectl wait -n cert-manager --for=condition=ready pod -l app=cert-manager --timeout=60s

# Configure operator with helm values
helm install scylla-operator scylla/scylla-operator --values hack/scylla-operator/helm/values.operator.yaml --create-namespace --namespace scylla-operator
kubectl wait -n scylla-operator --for=condition=ready pod -l app.kubernetes.io/name=scylla-operator --timeout=60s

# Configure manager with helm values
helm install scylla-manager scylla/scylla-manager --values hack/scylla-operator/helm/values.manager.yaml --create-namespace --namespace scylla-manager
kubectl wait -n scylla-manager --for=condition=ready pod -l app.kubernetes.io/name=scylla-manager --timeout=60s
kubectl wait -n scylla-manager --for=condition=ready pod -l app.kubernetes.io/name=scylla-manager-controller --timeout=60s

# Configure scylla with helm values
helm install scylla scylla/scylla --values hack/scylla-operator/helm/values.cluster.yaml --create-namespace --namespace scylla
kubectl wait -n scylla --for=condition=ready pod -l app.kubernetes.io/name=scylla --timeout=60s

# Uninstall
# helm uninstall scylla -n scylla
# helm uninstall scylla-manager -n scylla-manager
# helm uninstall scylla-operator -n scylla-operator

# Expose scylla
# kubectl port-forward --address 0.0.0.0 svc/scylla-client 9042:9042 -n scylla