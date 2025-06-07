CREATE TYPE "clothes_gender" AS ENUM (
  'unisex',
  'feminine',
  'masculine'
);

CREATE TYPE "clothes_category" AS ENUM (
  'shirt',
  'tshirt',
  'jacket',
  'pants',
  'jeans',
  'skirt',
  'shoes',
  'bra',
  'binder',
  'briefs',
  'boxers',
  'panties'
);

CREATE TABLE "brand" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "website" varchar,
  "trans_specific" boolean
);

CREATE TABLE "size" (
  "id" integer PRIMARY KEY,
  "id_range_min" integer NOT NULL,
  "id_range_max" integer NOT NULL,
  "id_brand" integer NOT NULL,
  "clothes_gender" clothes_gender,
  "clothes_category" clothes_category
);

CREATE TABLE "code" (
  "id_size" integer,
  "code" varchar,
  PRIMARY KEY ("id_size", "code")
);

CREATE TABLE "user" (
  "id" integer PRIMARY KEY,
  "id_measurement" integer NOT NULL,
  "username" varchar,
  "email" varchar,
  "password" varchar,
  "admin" boolean
);

CREATE TABLE "measurement" (
  "id" integer PRIMARY KEY,
  "height" decimal(12,2),
  "head" decimal(12,2),
  "neck" decimal(12,2),
  "shoulder" decimal(12,2),
  "chest" decimal(12,2),
  "underchest" decimal(12,2),
  "sleeve" decimal(12,2),
  "hip" decimal(12,2),
  "waist" decimal(12,2),
  "inseam" decimal(12,2),
  "foot" decimal(12,2)
);

ALTER TABLE "size" ADD FOREIGN KEY ("id_range_min") REFERENCES "measurement" ("id");

ALTER TABLE "size" ADD FOREIGN KEY ("id_range_max") REFERENCES "measurement" ("id");

ALTER TABLE "size" ADD FOREIGN KEY ("id_brand") REFERENCES "brand" ("id");

ALTER TABLE "code" ADD FOREIGN KEY ("id_size") REFERENCES "size" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("id_measurement") REFERENCES "measurement" ("id");
