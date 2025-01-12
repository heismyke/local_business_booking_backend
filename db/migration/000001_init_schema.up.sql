CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "businesses" (
  "id" bigserial PRIMARY KEY,
  "owner" bigint NOT NULL,
  "name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "lattitude" float NOT NULL,
  "longitude" float NOT NULL,
  "phone" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "category" varchar NOT NULL,
  "services" jsonb NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT(now())
);

CREATE TABLE "business_hours" (
  "id" bigserial PRIMARY KEY,
  "business_id" bigint NOT NULL,
  "day_of_week" varchar NOT NULL,
  "open_time" time NOT NULL,
  "close_time" time NOT NULL
);

CREATE TABLE "bookings" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "business_id" bigint NOT NULL,
  "service" varchar NOT NULL,
  "date" date NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "reviews" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "business_id" bigint NOT NULL,
  "rating" int NOT NULL,
  "comment" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "businesses" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "business_hours" ADD FOREIGN KEY ("business_id") REFERENCES "businesses" ("id");

ALTER TABLE "bookings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "bookings" ADD FOREIGN KEY ("business_id") REFERENCES "businesses" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("business_id") REFERENCES "businesses" ("id");