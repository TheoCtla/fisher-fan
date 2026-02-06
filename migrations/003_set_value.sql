-- 1. Création d'un Utilisateur (Propriétaire)
INSERT INTO "users" (
    "id", "last_name", "first_name", "birth_date", "email", 
    "boat_license", "status"
) VALUES (
    '1', 
    'toto', 
    'tutu', 
    '1995-03-12', 
    'toto@test.com', 
    '87654321', 
    'professional'
);

-- 2. Création d'un Bateau lié à l'User 1
INSERT INTO "boats" (
    "id", "user_id", "name", "brand", 
    "manufacture_year", "max_capacity", "home_port"
) VALUES (
    '1', 
    '1', 
    'L''Espadon', 
    'Beneteau', 
    '2022-01-01', 
    8, 
    'Marseille'
);

-- 3. Ajout d'équipements pour le bateau 1
INSERT INTO "boatEquipment" ("boat_id", "name") VALUES 
('1', 'Sonar'),
('1', 'GPS');

-- 4. Création d'un Voyage (Trip) lié au Bateau 1 et User 1
INSERT INTO "trips" (
    "id", "user_id", "boat_id", "title", 
    "trip_type", "passenger_count", "price"
) VALUES (
    '1', 
    '1', 
    '1', 
    'Sortie Pêche au Thon', 
    'daily', 
    4, 
    150.0
);

-- 5. Ajout d'un horaire pour le voyage 1
INSERT INTO "tripSchedules" (
    "trip_id", "start_date", "departure_time"
) VALUES (
    '1', 
    '2026-07-10 00:00:00', 
    '07:00:00'
);

-- 1. Création d'un second utilisateur (le client)
INSERT INTO "users" (
    "id", "last_name", "first_name", "email", "status"
) VALUES (
    '2', 
    'Martin', 
    'Alice', 
    'alice@test.com', 
    'individual'
);

-- 2. Création d'une réservation
-- L'utilisateur 2 réserve 2 places pour le voyage 1
INSERT INTO "reservations" (
    "id", 
    "trip_id", 
    "user_id", 
    "date", 
    "reserved_seats", 
    "total_price"
) VALUES (
    '1', 
    '1', 
    '2', 
    '2026-07-10 07:00:00', 
    2, 
    300.0
);