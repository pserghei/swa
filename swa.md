# K8s

Kubernetes ist eine Plattform für die Verwaltung von containerisierter Software, die sowohl die deklarative Konfiguration als auch die Automatisierung erleichtert.

## Why K8s?

### Service Discovery und Load Balancing
 Kubernetes kann einen Container sichtbar machen. Wenn der Datenverkehr zu einem Container hoch ist, kann Kubernetes den Netzwerkverkehr so verteilen, dass die Bereitstellung stabil ist und somit die Last ausgleichen.

### Orchestrierung von Speicher
Kubernetes ermöglicht es ein Speichersystem automatisch einzubinden, z. B. lokale Speicher, cloud-basierte Speicer, etc.

### Automatisierte Rollouts und Rollbacks
Mit Kubernetes lässt sich der gewünschte Zustand für bereitgestellte Container beschreiben, und Kubernetes kann den aktuellen Zustand in den gewünschten Zustand ändern. 

### Selbstheilung
Container werden automatisch nach Fehler neu gestartet, ersetzt, nach einem nicht-reagieren auf die benutzerdefinierte Gesundheitsprüfung gelöscht, und Clients erst gezeigt, wenn die Container bereit sind die Clients zu bedienen.

### Horizontale Skalierung
Anwendungen können automatisch horizontal skaliert werden, d.h. ein Kubernetescluster kann auf weitere Maschinen ausgelagert werden und die Last somit verteilt werden.

### Gestaltet für Erweiterbarkeit
Es ist möglich ein Kubernetes-Cluster durch Funktionen zu erweitern, ohne mit dem Kubernetesquellcode interagieren zu müssen.

## Architektur

### Kubernetes Components

- Control Plane
- Node

- Kubernetes Objects
- Object Management

### Cluster Architecture

- Nodes
- Communication between Nodes and Control Plane
- Controllers
- Leases
- Garbage Collection

## Quellenverweis
- https://kubernetes.io/docs/concepts/overview
- https://kubernetes.io/docs/concepts/overview/components
- https://kubernetes.io/docs/concepts/overview/working-with-objects
- https://kubernetes.io/docs/concepts/architecture
