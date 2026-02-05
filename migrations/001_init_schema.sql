CREATE TABLE "users" (
  "id" varchar PRIMARY KEY NOT NULL,
  "last_name" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "birth_bate" date,
  "email" varchar UNIQUE NOT NULL,
  "boat_license" varchar(8),
  "status" varchar NOT NULL,
  "company_name" varchar,
  "activity_type" varchar,
  "siret_number" varchar(14),
  "rc_number" varchar
);

CREATE TABLE "boats" (
  "id" varchar PRIMARY KEY NOT NULL,
  "user_id" varchar NOT NULL,
  "name" varchar NOT NULL,
  "description" text,
  "brand" varchar,
  "manufacture_year" date,
  "photo_url" varchar,
  "license_type" varchar,
  "boat_type" varchar,
  "deposit_amount" float,
  "max_capacity" int,
  "number_of_beds" int,
  "home_port" varchar,
  "latitude" float,
  "longitude" float,
  "engine_type" varchar,
  "engine_power" int
);

CREATE TABLE "boatEquipment" (
  "boat_id" varchar NOT NULL,
  "name" varchar NOT NULL,
  PRIMARY KEY ("boat_id", "name")
);

CREATE TABLE "trips" (
  "id" varchar PRIMARY KEY NOT NULL,
  "user_id" varchar NOT NULL,
  "boat_id" varchar NOT NULL,
  "title" varchar NOT NULL,
  "practical_info" text,
  "trip_type" varchar,
  "rate_type" varchar,
  "passenger_count" int,
  "price" float
);

CREATE TABLE "tripSchedules" (
  "trip_id" varchar NOT NULL,
  "start_date" timestamp NOT NULL,
  "end_date" timestamp,
  "departure_time" time,
  "end_time" time,
  PRIMARY KEY ("trip_id", "start_date", "departure_time")
);

CREATE TABLE "reservations" (
  "id" varchar PRIMARY KEY NOT NULL,
  "trip_id" varchar NOT NULL,
  "user_id" varchar NOT NULL,
  "date" timestamp NOT NULL,
  "reserved_seats" int,
  "total_price" float
);

CREATE TABLE "logs" (
  "id" varchar PRIMARY KEY NOT NULL,
  "user_id" varchar UNIQUE NOT NULL
);

CREATE TABLE "pages" (
  "id" varchar PRIMARY KEY NOT NULL,
  "log_id" varchar NOT NULL,
  "user_id" varchar NOT NULL,
  "fish_name" varchar,
  "fish_photo_url" varchar,
  "comment" text,
  "length" float,
  "weight" float,
  "fishing_spot" varchar,
  "fishing_date" date,
  "release" boolean
);

ALTER TABLE "boats" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "boatEquipment" ADD FOREIGN KEY ("boat_id") REFERENCES "boats" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("boat_id") REFERENCES "boats" ("id");

ALTER TABLE "tripSchedules" ADD FOREIGN KEY ("trip_id") REFERENCES "trips" ("id");

ALTER TABLE "reservations" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reservations" ADD FOREIGN KEY ("trip_id") REFERENCES "trips" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "pages" ADD FOREIGN KEY ("log_id") REFERENCES "logs" ("id");

ALTER TABLE "pages" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
