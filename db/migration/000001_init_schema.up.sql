CREATE TABLE "carpools" (
  "id" bigserial PRIMARY KEY,
  "contact_info" varchar NOT NULL,
  "driver_id" bigserial NOT NULL,
  "max_passenger" int NOT NULL DEFAULT 1,
  "price" int NOT NULL,
  "able_pick_up" boolean NOT NULL DEFAULT false,
  "resort_id" bigserial NOT NULL,
  "departure_at" timestamp NOT NULL,
  "accept_payment_type" varchar NOT NULL DEFAULT 'Cash',
  "currency" varchar NOT NULL DEFAULT 'CAD',
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "stations" (
  "id" bigserial PRIMARY KEY,
  "carpool_id" bigserial NOT NULL,
  "station_name" VARCHAR NOT NULL,
  "arrival_minutes" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "carpool_users" (
  "id" bigserial PRIMARY KEY,
  "passenger_id" bigserial NOT NULL,
  "carpool_id" bigserial NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "carpool_requests" (
  "id" bigserial PRIMARY KEY,
  "carpool_id" bigserial NOT NULL,
  "passenger_id" bigserial NOT NULL,
  "boarding_station" bigserial NOT NULL,
  "payment_type" varchar NOT NULL,
  "currency" varchar NOT NULL DEFAULT 'CAD',
  "contact_info" varchar NOT NULL,
  "approved" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "carpools" ("resort_id");

CREATE INDEX ON "carpools" ("departure_at");

CREATE INDEX ON "carpools" ("accept_payment_type");

CREATE INDEX ON "carpools" ("able_pick_up");

CREATE INDEX ON "stations" ("carpool_id");

CREATE INDEX ON "carpool_users" ("passenger_id");

CREATE INDEX ON "carpool_users" ("carpool_id");

CREATE INDEX ON "carpool_requests" ("carpool_id");

COMMENT ON COLUMN "carpools"."accept_payment_type" IS 'provider write';

COMMENT ON COLUMN "stations"."arrival_minutes" IS 'How long it take from first station';

ALTER TABLE "stations" ADD FOREIGN KEY ("carpool_id") REFERENCES "carpools" ("id");

ALTER TABLE "carpool_users" ADD FOREIGN KEY ("carpool_id") REFERENCES "carpools" ("id");

ALTER TABLE "carpool_requests" ADD FOREIGN KEY ("carpool_id") REFERENCES "carpools" ("id");

ALTER TABLE "carpool_requests" ADD FOREIGN KEY ("boarding_station") REFERENCES "stations" ("id");
