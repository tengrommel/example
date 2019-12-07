# Kubernetes Node


## file content

    apiVersion: v1
    kine: ReplicationController
    metadata:
        name: nginx
    spec:
        replicas: 3
        selector:
            app: nginx
        template:
            metadata:
                name: nginx
                labels:
                    app: nginx
            spec:
                containers:
                    name: nginx
                    image: nginx
                    ports:
                        containerPort: 80