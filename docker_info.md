**Docker : Images & Conteneurs - Résumé des Commandes**

---

### 📂 1. Trouver les Images Docker

Lister toutes les images stockées localement :  
```bash
docker images
```

Vérifier l'emplacement des images :  
- **Linux & Mac** : `/var/lib/docker`
- **Windows (WSL)** : `\\wsl$\docker-desktop-data\mnt\wsl\docker-desktop`

---

### 🚢 2. Trouver les Conteneurs Docker

Lister les conteneurs en cours d'exécution :  
```bash
docker ps
```

Lister **tous** les conteneurs (y compris ceux arrêtés) :  
```bash
docker ps -a
```

Voir les détails d'un conteneur spécifique :  
```bash
docker inspect <CONTAINER_ID>
```

Vérifier l'emplacement des fichiers d'un conteneur :  
- **Linux & Mac** : `/var/lib/docker/containers`
- **Windows (WSL)** : `\\wsl$\docker-desktop-data\mnt\wsl\docker-desktop`

---

### 🛢️ 3. Supprimer des Images et Conteneurs

Supprimer une image Docker :  
```bash
docker rmi <IMAGE_ID>
```

Supprimer un conteneur arrêté :  
```bash
docker rm <CONTAINER_ID>
```

Nettoyer toutes les images et conteneurs inutilisés :  
```bash
docker system prune -a
```

---

### ✨ Résumé des Commandes Clés :

| Commande | Fonction |
|----------|----------|
| `docker images` | Voir toutes les images Docker locales |
| `docker ps` | Voir les conteneurs actifs |
| `docker ps -a` | Voir tous les conteneurs (actifs et arrêtés) |
| `docker inspect <ID>` | Voir les détails d'un conteneur |
| `docker rmi <IMAGE_ID>` | Supprimer une image Docker |
| `docker rm <CONTAINER_ID>` | Supprimer un conteneur arrêté |
| `docker system prune -a` | Nettoyer les images et conteneurs inutilisés |

🚀 **Docker stocke ses fichiers dans `/var/lib/docker` sur Linux/Mac et dans `WSL` sur Windows.**
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
### Docker b Darija Marocaine

#### 1. Docker Chno Howa ?
Docker howa wa7ed l'outil li kaykhdem bach ndirou des applications f **containers**. Containers kaykhdmou b7al des mini-machines mais kayb9aw m3azlin (isolés) w kay5dmo dakhil système principal dyal l'ordinateur b wa7ed tariqa 7it "**lightweight**".

**Exemple** : Katkhayel dakchi b7al ila 3ndek PC w ghirti t5dam plusieurs applications f chi système d'exploitation akhor bla ma tzid virtual machine. Docker kaykhdm f chi haja b7alha.

---

#### 2. Chno Howa **Container** f Docker ?
Container howa wa7ed l'espace li kaydakhlo code dyalk w kaykhdem f wa7ed l'environement séparé b7al sandbox. 3ndna plusieurs containers f nafs PC, w kol container kayb9a indépendant b chwiya dial configuration.

**Exemple** : Ila 3ndek un projet li kaykhdm b React f frontend, Node.js f backend, w MySQL f database, Docker kayjma3hom 7ta wahd fi container b configuration spéciale.

---

#### 3. Docker Image Chno Hiya ?
Docker Image hiya b7al blueprint (plan) li kay3tik container mchi ndirih mn chi base. Kadirha marra wahda w men b3d katkhadam biha bach dir plusieurs containers menha.

**Exemple** : Kayna image dyal **alpine** hiya version light dyal Linux, kayna image dyal **node**, **mysql**... li 9darti tkhaddm b Docker.

Commande bach njibo image men Docker Hub :
```sh
  docker pull alpine:latest
```

Bach njibo plusieurs images f PC dyalna :
```sh
  docker images
```

---

#### 4. Chno Howa **Cache** w 3lach kaymhem f Docker ?

Cache howa **mémoire temporaire** li kay5zen fichier li kayt7taj bech tkhdam bchi haja. Bach tfhamha, khi lik chi exemple :

- Ila wslti chi site web, navigateur dyalk kaykhzen chi haja mnha bach mayloadach mra khra.
- F Docker, ila installina chi package, kayb9a l'install f cache, w hna l'image dyalna katzid f size bla faida.

**Exemple Docker bach ninstalliw Bash** :
```sh
  RUN apk add bash
```
Hna **apk** kay installi bash mais kaykhalli les fichiers f cache. L'image docker tatwli kbira.

**Solution ?** Kayna option **--no-cache** li kay3ti l'ordre bach maykhznch f cache :
```sh
  RUN apk add --no-cache bash
```
Hadchi kay5li l'image docker **sghira w rapide** ! 🚀

---

#### 5. Test Bach Tshouf Far9 bin `--no-cache` w bla biha

1️⃣ **Dockerfile sans --no-cache**
```dockerfile
FROM alpine:latest
RUN apk add bash
```
```sh
docker build -t test-cache .
docker images | grep test-cache
```

2️⃣ **Dockerfile avec --no-cache**
```dockerfile
FROM alpine:latest
RUN apk add --no-cache bash
```
```sh
docker build -t test-no-cache .
docker images | grep test-no-cache
```

📌 Hna, image **test-no-cache** ghadi tkoun **sghira** bzzaf 3la test-cache 7it makaykhznch les fichiers inutilement ! 🎯

---

### **Conclusion**
✔ **Docker** kay5dmek containers bach matb9ach tinstalli ghir les outils b system dyalk.
✔ **Container** howa environement séparé bach t5dam application f wa7ed l'espace m3azoul.
✔ **Docker Image** hiya b7al template li katbni 3liha container.
✔ **Cache** howa l'espace li kayb9a fih les fichiers, w **`--no-cache`** kaykhdm bach maykhznch chi haja zayda w ykhdm docker **bchkal optimal**.

🚀 Wach fhamti ? Oula briti chi haja akhor ndirou test b code ? 😃


