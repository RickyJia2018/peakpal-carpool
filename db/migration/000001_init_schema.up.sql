CREATE TABLE "trips" (
  "id" bigserial PRIMARY KEY,
  "contact_info" varchar NOT NULL,
  "driver_id" bigserial NOT NULL,
  "max_passenger" int NOT NULL DEFAULT 1,
  "price" int NOT NULL,
  "able_pick_up" boolean NOT NULL DEFAULT false,
  "resort_id" bigserial NOT NULL,
  "departure_at" timestamp NOT NULL,
  "return_at" timestamp NOT NULL,
  "round_trip" boolean NOT NULL DEFAULT true,
  "accept_payment_type" varchar NOT NULL DEFAULT 'Cash',
  "currency" varchar NOT NULL DEFAULT 'CAD',
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "stations" (
  "id" bigserial PRIMARY KEY,
  "trip_id" bigserial NOT NULL,
  "station_name" VARCHAR NOT NULL,
  "arrival_minutes" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "passengers" (
  "id" bigserial PRIMARY KEY,
  "passenger_id" bigserial NOT NULL,
  "trip_id" bigserial NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "trip_requests" (
  "id" bigserial PRIMARY KEY,
  "trip_id" bigserial NOT NULL,
  "passenger_id" bigserial NOT NULL,
  "boarding_station" bigserial NOT NULL,
  "payment_type" varchar NOT NULL,
  "currency" varchar NOT NULL DEFAULT 'CAD',
  "contact_info" varchar NOT NULL,
  "approved" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "trips" ("resort_id");

CREATE INDEX ON "trips" ("departure_at");

CREATE INDEX ON "trips" ("accept_payment_type");

CREATE INDEX ON "trips" ("able_pick_up");

CREATE INDEX ON "stations" ("trip_id");

CREATE INDEX ON "passengers" ("passenger_id");

CREATE INDEX ON "passengers" ("trip_id");

CREATE INDEX ON "trip_requests" ("trip_id");

COMMENT ON COLUMN "trips"."accept_payment_type" IS 'provider write';

COMMENT ON COLUMN "stations"."arrival_minutes" IS 'How long it take from first station';

ALTER TABLE "stations" ADD FOREIGN KEY ("trip_id") REFERENCES "trips" ("id");

ALTER TABLE "passengers" ADD FOREIGN KEY ("trip_id") REFERENCES "trips" ("id");

ALTER TABLE "trip_requests" ADD FOREIGN KEY ("trip_id") REFERENCES "trips" ("id");

ALTER TABLE "trip_requests" ADD FOREIGN KEY ("boarding_station") REFERENCES "stations" ("id");
