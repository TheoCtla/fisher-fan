# 🐟 Fisher Fans API

Fisher Fans est une API REST développée en **Go** (Golang) conçue pour la gestion des passionnés de pêche. Elle permet de gérer les utilisateurs, leurs bateaux, leurs sorties de pêche (Trips), les réservations et un journal de bord (Log) pour consigner leurs plus belles captures.

## 🚀 Technologies utilisées

* **Langage :** Go (1.22+)
* **Framework Web :** Gin Gonic
* **ORM :** GORM
* **Base de données :** PostgreSQL 16
* **Conteneurisation :** Docker & Docker Compose
* **Documentation :** OpenAPI 3.1 (Swagger)

---

## 🛠️ Installation et Lancement (Docker)

Le projet est entièrement conteneurisé. Pour le lancer, assurez-vous d'avoir Docker installé, puis suivez ces étapes :

1.  **Cloner le dépôt :**
    ```bash
    git clone [https://github.com/ton-username/fisherfans-api.git](https://github.com/ton-username/fisherfans-api.git)
    cd fisherfans-api
    ```

2.  **Configurer les variables d'environnement :**
    Vérifiez le fichier `.env` à la racine. Pour Docker, l'hôte de la base de données doit être le nom du service :
    ```env
    SERVER_PORT=8080
    DB_HOST=postgres
    DB_USER=admin
    DB_PASSWORD=admin123
    DB_NAME=fisherfan
    ```

3.  **Lancer l'application :**
    ```bash
    docker-compose up --build
    ```

L'API sera accessible sur `http://localhost:8080/api/v1`.
Un outil de gestion de base de données (**Adminer**) est disponible sur `http://localhost:8081`.

---

## 📮 Tests avec Postman

Pour tester l'API, vous pouvez importer la collection fournie :

1.  Ouvrez **Postman**.
2.  Cliquez sur **Import** et sélectionnez le fichier `FisherFans.postman_collection.json`.
3.  Configurez une variable d'environnement `base_url` : `http://localhost:8080/api/v1`.
4.  Les dossiers Postman suivent l'ordre logique : User -> Boat -> Log -> Trip -> Reservation.

---

## 📂 Architecture du Projet

Le projet suit une architecture modulaire segmentée par version d'API :

* **`/cmd`** : Point d'entrée de l'application (initialisation du serveur).
* **`/internal/api/v1`** : Cœur de l'application (version 1) :
    * **`/models`** : Définition des schémas de données et structures GORM.
    * **`/handlers`** : Contrôleurs traitant les requêtes HTTP.
    * **`/services`** : Logique métier (traitement des données).
    * **`/repository`** : Couche d'accès aux données (requêtes SQL via GORM).
    * **`/routes`** : Définition des points d'accès (endpoints).
* **`/internal/database`** : Configuration et connexion à PostgreSQL.
* **`/internal/server`** : Configuration du serveur HTTP Gin.
* **`/internal/variables`** : Gestion des variables d'environnement.
* **`/migrations`** : Scripts SQL d'initialisation de la base de données.