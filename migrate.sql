CREATE TABLE "roles" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "descriptions" TEXT NOT NULL
);

CREATE TABLE "permissions" (
   "id" BIGSERIAL PRIMARY KEY,
   "name" TEXT NOT NULL,
   "description" TEXT NOT NULL
);


CREATE TABLE "priorities" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" TEXT NOT NULL,
  "descriptions" TEXT NOT NULL
);

CREATE TABLE "priority2permissions" (
    "priority_id" BIGINT references priorities,
    "permission_id" BIGINT references permissions,
);

CREATE TABLE "permissions2role" (
    "permission_id" BIGINT references permissions,
    "role_id" BIGINT references roles,
);

CREATE TABLE "employeers" (
      "id" BIGSERIAL PRIMARY KEY,
      "first_name" TEXT NOT NULL,
      "surname" TEXT NOT NULL,
      "father_name" TEXT NOT NULL,
      "date_of_birth" DATE NOT NULL,
      "gender" BOOLEAN NOT NULL,
      "citizenship" TEXT NOT NULL,
      "passport_series" TEXT NOT NULL,
      "passport_number" BIGINT NOT NULL,
      "date_of_issue" DATE NOT NULL,
      "place_of_issue" TEXT NOT NULL,
      "residential_address" TEXT NOT NULL,
      "role_id" BIGINT references roles,
);

CREATE TABLE "clients" (
   "id" BIGSERIAL PRIMARY KEY,
   "priority_id" BIGINT references priorities,
   "name" TEXT NOT NULL,
   "surname" TEXT NOT NULL,
   "father_name" TEXT NOT NULL,
   "date_of_birth" DATE NOT NULL,
   "gender" BOOLEAN NOT NULL,
   "citizenship" TEXT NOT NULL,
   "passport_series" TEXT NOT NULL,
   "passport_number" BIGINT NOT NULL,
   "date_of_issue" TEXT NOT NULL,
   "place_of_issue" TEXT NOT NULL,
   "residential_address" TEXT NOT NULL,
);

CREATE TABLE "rules" (
     "id" BIGSERIAL PRIMARY KEY,
     "commission" DOUBLE NOT NULL,
     "term" INT NOT NULL,
     "active" BOOLEAN NOT NULL default true
);

CREATE TABLE "types" (
     "id" BIGSERIAL PRIMARY KEY,
     "name" TEXT not null,
     "description" TEXT not null,
     "active" BOOLEAN NOT NULL default true
);

CREATE TABLE "services" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NULL,
    "description" TEXT NULL,
    "type_id" BIGINT references types,
    "rule_id" BIGINT references rules,
    "created_at" TIMESTAMP default current_timestamp,
    "created_by" bigint references employeers,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIME,
    "updated_by" BIGINT references employeers,
    "deleted_at" BIGINT,
    "deleted_by" BIGINT references employeers
);

CREATE TABLE "clients2services" (
    "client_id" BIGINT references clients,
    "service_id" BIGINT references services
);

