# K8s

Kubernetes ist eine Plattform für die Verwaltung von containerisierter Software, die sowohl die deklarative Konfiguration als auch die Automatisierung erleichtert.

## Wieso K8s?

### Service Discovery und Load Balancing
 Kubernetes kann einen Container sichtbar machen. Wenn der Datenverkehr zu einem Container hoch ist, kann Kubernetes den Netzwerkverkehr so verteilen, dass die Bereitstellung stabil ist und somit die Last ausgleichen.

### Orchestrierung von Speicher
Kubernetes ermöglicht es ein Speichersystem automatisch einzubinden, z. B. lokale Speicher, cloud-basierte Speicer, etc.

### Automatisierte Rollouts und Rollbacks
Mit Kubernetes lässt sich der gewünschte Zustand für bereitgestellte Container beschreiben, und Kubernetes kann den aktuellen Zustand in den gewünschten Zustand ändern. 

### Selbstheilung
Container werden automatisch nach Fehlern neu gestartet, ersetzt, nach einem nicht-reagieren auf die benutzerdefinierte Gesundheitsprüfung gelöscht, und Clients erst gezeigt, wenn die Container bereit sind die Clients zu bedienen.

### Horizontale Skalierung
Anwendungen können automatisch horizontal skaliert werden, d.h. ein Kubernetescluster kann auf weitere Maschinen ausgelagert werden und die Last somit verteilt werden.

### Gestaltet für Erweiterbarkeit
Es ist möglich ein Kubernetes-Cluster durch Funktionen zu erweitern, ohne mit dem Kubernetesquellcode interagieren zu müssen.

## Architektur

### Kubernetes Komponenten

Ein Kubernetes-Cluster besteht aus einem Steuerelement und einem oder mehreren Arbeitsknoten, so genannten worker nodes.

Das Steuerelement selbst besteht aus mehreren Komponenten. Darunter zählen:
- kube-apiserver, der server der die Kubernetes HTTP API zur Verfügung stellt
- etcd, ein Key-Value-Store für die Daten des API-Servers
- kube-scheduler, sorgt dafür, dass Pods, die noch keinem Knoten zugewiesen sind, einem passenden Knoten zugewiesen werden
- kube-controller-manager, betreibt Controller um das Verhalten der Kubernetes API zu implementieren. Controller dienen dazu den derzeitigen Zustand eines Clusters in den gewünschten Zustand zu bringen

Arbeitsknoten bestehen ebenfalls aus mehreren Komponenten und sorgen für das Aufrechterhalten von Pods und stellen die Kubernetes-Laufzeitumgebung bereit. Dazu gehört:

- kubelet, stellen sicher, dass Pods und deren Container laufen
- Container-Laufzeit, Software die für das betreiben von Containern verantwortlich ist

Ebenfalls existieren in Kubernetes Kubernetesobjekte. Das sind Entitäten im Kubernetessystem. Diese werden von Kubernetes genutzt um die Zustand von Clustern darzustellen. Diese beschreiben:
- Welche containerisierte Applikationen derzeit laufen und auf welchen Knoten
- Welche Ressourcen diesen Applikationen zur Verfügung stehen
- Die Richtlinien, an welche sich diese Applikationen halten. Dazu gehört z. B. das Verhalten in Beziehung zum Neustarten.

Kubernetesobjekte können mithilfe des CLI tools `kubectl` verwaltet werden.

### Cluster Architecture

Ein Kubernetes Cluster besteht aus einer Steuerelement und einer Reihe von Arbeitsmaschinen, so genannten Arbeitsknoten.

Kubernetes betreibt Applikationen indem es Container in Pods platziert, welche wiederum auf Knoten laufen. Ein Pod ist einfach eine Gruppe von Containern in einem Cluster. Arbeitsknoten können dabei sowohl virtuelle, als auch phyische Maschinen sein. Für gewöhnlich besteht ein Cluster aus mehreren Knoten.

## Quellenverweis
- https://kubernetes.io/docs/concepts/overview
- https://kubernetes.io/docs/concepts/overview/components
- https://kubernetes.io/docs/concepts/overview/working-with-objects
- https://kubernetes.io/docs/concepts/architecture
