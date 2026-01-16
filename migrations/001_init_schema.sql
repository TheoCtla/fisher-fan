-- Migration initiale pour FisherFan
-- Cette migration crée toutes les tables nécessaires à l'application

-- Tables de référence (lookup tables)

CREATE TABLE IF NOT EXISTS status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS villes (
    id SERIAL PRIMARY KEY,
    nom VARCHAR(100) NOT NULL,
    code_postal VARCHAR(10) NOT NULL
);

CREATE TABLE IF NOT EXISTS types_bateau (
    id SERIAL PRIMARY KEY,
    nom VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS types_sortie (
    id SERIAL PRIMARY KEY,
    nom VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS types_tarif (
    id SERIAL PRIMARY KEY,
    nom VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS types_motorisation (
    id SERIAL PRIMARY KEY,
    nom VARCHAR(50) NOT NULL
);

-- Table utilisateurs
CREATE TABLE IF NOT EXISTS utilisateurs (
    id SERIAL PRIMARY KEY,
    nom VARCHAR(100),
    prenom VARCHAR(100),
    date_naissance DATE,
    email VARCHAR(255) UNIQUE NOT NULL,
    telephone VARCHAR(20),
    adresse VARCHAR(255),
    ville_id INT REFERENCES villes(id) ON DELETE SET NULL,
    langues_parlees VARCHAR(255),
    url_avatar VARCHAR(500),
    numero_permis_bateau VARCHAR(8),
    numero_assurance VARCHAR(12),
    id_status INT REFERENCES status(id) ON DELETE SET NULL,
    nom_societe VARCHAR(255),
    type_activite VARCHAR(50),
    numero_siret VARCHAR(14),
    numero_rc VARCHAR(50),
    date_creation TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    date_modification TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table bateaux
CREATE TABLE IF NOT EXISTS bateaux (
    id SERIAL PRIMARY KEY,
    utilisateur_id INT NOT NULL REFERENCES utilisateurs(id) ON DELETE CASCADE,
    nom VARCHAR(100) NOT NULL,
    description TEXT,
    marque VARCHAR(100),
    annee_fabrication INT,
    url_photo VARCHAR(500),
    type_permis_requis VARCHAR(50),
    type_bateau_id INT REFERENCES types_bateau(id) ON DELETE SET NULL,
    equipements TEXT,
    montant_caution DECIMAL(10,2),
    capacite_max INT,
    nombre_couchages INT,
    port_attache VARCHAR(100),
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),
    type_motorisation_id INT REFERENCES types_motorisation(id) ON DELETE SET NULL,
    puissance_moteur INT
);

-- Table sorties_peche
CREATE TABLE IF NOT EXISTS sorties_peche (
    id SERIAL PRIMARY KEY,
    utilisateur_id INT NOT NULL REFERENCES utilisateurs(id) ON DELETE CASCADE,
    bateau_id INT NOT NULL REFERENCES bateaux(id) ON DELETE CASCADE,
    titre VARCHAR(255) NOT NULL,
    informations_pratiques TEXT,
    type_sortie_id INT REFERENCES types_sortie(id) ON DELETE SET NULL,
    type_tarif_id INT REFERENCES types_tarif(id) ON DELETE SET NULL,
    date_debut TIMESTAMP,
    date_fin TIMESTAMP,
    heure_depart TIME,
    heure_fin TIME,
    nombre_passagers INT,
    prix DECIMAL(10,2)
);

-- Table reservations
CREATE TABLE IF NOT EXISTS reservations (
    id SERIAL PRIMARY KEY,
    sortie_id INT NOT NULL REFERENCES sorties_peche(id) ON DELETE CASCADE,
    utilisateur_id INT NOT NULL REFERENCES utilisateurs(id) ON DELETE CASCADE,
    date_reservation TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    nombre_places INT NOT NULL,
    prix_total DECIMAL(10,2) NOT NULL
);

-- Table carnets_peche
CREATE TABLE IF NOT EXISTS carnets_peche (
    id SERIAL PRIMARY KEY,
    utilisateur_id INT NOT NULL REFERENCES utilisateurs(id) ON DELETE CASCADE,
    titre VARCHAR(255),
    date_creation TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table pages_carnet
CREATE TABLE IF NOT EXISTS pages_carnet (
    id SERIAL PRIMARY KEY,
    carnet_id INT NOT NULL REFERENCES carnets_peche(id) ON DELETE CASCADE,
    nom_poisson VARCHAR(100),
    url_photo VARCHAR(500),
    commentaire TEXT,
    taille_cm DECIMAL(5,2),
    poids_kg DECIMAL(5,2),
    lieu_peche VARCHAR(255),
    date_peche DATE,
    poisson_relache BOOLEAN DEFAULT FALSE
);

-- Index pour améliorer les performances
CREATE INDEX IF NOT EXISTS idx_utilisateurs_email ON utilisateurs(email);
CREATE INDEX IF NOT EXISTS idx_utilisateurs_ville ON utilisateurs(ville_id);
CREATE INDEX IF NOT EXISTS idx_bateaux_utilisateur ON bateaux(utilisateur_id);
CREATE INDEX IF NOT EXISTS idx_bateaux_position ON bateaux(latitude, longitude);
CREATE INDEX IF NOT EXISTS idx_sorties_utilisateur ON sorties_peche(utilisateur_id);
CREATE INDEX IF NOT EXISTS idx_sorties_bateau ON sorties_peche(bateau_id);
CREATE INDEX IF NOT EXISTS idx_reservations_sortie ON reservations(sortie_id);
CREATE INDEX IF NOT EXISTS idx_reservations_utilisateur ON reservations(utilisateur_id);
CREATE INDEX IF NOT EXISTS idx_carnets_utilisateur ON carnets_peche(utilisateur_id);
CREATE INDEX IF NOT EXISTS idx_pages_carnet ON pages_carnet(carnet_id);

-- Insertion des données de référence

-- Status (particulier / professionnel)
INSERT INTO status (name) VALUES 
    ('particulier'),
    ('professionnel')
ON CONFLICT DO NOTHING;

-- Types de bateau
INSERT INTO types_bateau (nom) VALUES 
    ('open'),
    ('cabine'),
    ('catamaran'),
    ('voilier'),
    ('jetski'),
    ('canoë')
ON CONFLICT DO NOTHING;

-- Types de sortie
INSERT INTO types_sortie (nom) VALUES 
    ('journaliere'),
    ('recurrente')
ON CONFLICT DO NOTHING;

-- Types de tarif
INSERT INTO types_tarif (nom) VALUES 
    ('global'),
    ('par_personne')
ON CONFLICT DO NOTHING;

-- Types de motorisation
INSERT INTO types_motorisation (nom) VALUES 
    ('diesel'),
    ('essence'),
    ('aucun')
ON CONFLICT DO NOTHING;
