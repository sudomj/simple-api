# Simple API

<p>Project to test simple golang API running in K8S cluster.</p>


#### Run the project

```sh
kubectl apply -f ./.dev
``` 

#### Check if pods is running
```sh
kubectl get pods
```

#### Describe config map
```sh
kubectl describe configmap nginx-conf
```

#### Check services
```sh 
kubectl get svc
```

#### Port forward nginx to access API's
```sh
kubectl port-forward svc/nginx-service 5000:80
```

#### Curl endpoint
```sh
curl http://localhost:5000/health -vvv
```

