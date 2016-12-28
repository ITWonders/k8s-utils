# k8s-utils
- very small and simple shell utils for kubernetes to extract output from kubectl for further usage (e.g. ssh into pod etc)
- tested on google cloud platform
- Note: nothing fancy or comprehensive stuff here, only so simple tool to save you some keystoke and repeated copy using mouse

## Example usage of copying pod name to clipboard
- to copy the pod name to clipboard, so that you can use it with other command
```
-> % ku
Retriving kind: po
0: cassandra-0
1: cassandrabk-bpqct
2: ecv-go-3332414171-dj9ht
3: ecv-storage
4: elasticsearch-rp2g7


Choose one number (will be copied): 2
```
- now string "ecv-go-3332414171-dj9ht" is copied to clipboard

## Example usage to ssh into pod
- to ssh into one of the pod
```
-> % kushell
Retriving kind: po
0: cassandra-0
1: cassandrabk-bpqct
2: ecv-go-3332414171-dj9ht
3: ecv-storage
4: elasticsearch-rp2g7


Choose one number to ssh: 2
kubectl exec -it ecv-go-3332414171-dj9ht -- /bin/sh
```

## Example usage to ssh into node
```
-> % kushell no
Retriving kind: no
0: gke-ecv-default-pool-387043f3-tlt6


Choose one number to ssh: 0
gcloud compute ssh gke-ecv-default-pool-387043f3-tlt6

Welcome to Kubernetes v1.5.1!
```
