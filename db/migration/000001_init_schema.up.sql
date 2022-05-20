CREATE TABLE "users" (
  "email" varchar PRIMARY KEY,
  "password" varchar NOT NULL
);

CREATE TABLE "timetable" (
  "route" varchar PRIMARY KEY,
  "arrival" timestamptz NOT NULL,
  "departure" timestamptz NOT NULL,
  "distance" integer NOT NULL,
  "price" integer NOT NULL
);

CREATE TABLE "trains" (
  "train_number" bigserial PRIMARY KEY,
  "type" varchar NOT NULL,
  "class" varchar NOT NULL,
  "max_passenger_no" integer NOT NULL,
  "max_speed" integer NOT NULL,
  "route" varchar NOT NULL
);

CREATE TABLE "tickets" (
  "ticket_id" bigserial PRIMARY KEY,
  "route" varchar NOT NULL,
  "train_number" bigserial NOT NULL,
  "coach_number" integer NOT NULL,
  "seat_number" integer NOT NULL,
  "booking_date" date NOT NULL,
  "trip_date" date NOT NULL,
  "fare" integer NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "payments" (
  "ticket_id" bigserial NOT NULL,
  "amount" integer NOT NULL,
  "credit_card_number" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "trains" ("route");

CREATE INDEX ON "tickets" ("route");

CREATE INDEX ON "tickets" ("train_number");

CREATE INDEX ON "tickets" ("email");

CREATE INDEX ON "payments" ("ticket_id");

ALTER TABLE "trains" ADD FOREIGN KEY ("route") REFERENCES "timetable" ("route") ON DELETE CASCADE;

ALTER TABLE "tickets" ADD FOREIGN KEY ("route") REFERENCES "timetable" ("route") ON DELETE CASCADE;

ALTER TABLE "tickets" ADD FOREIGN KEY ("train_number") REFERENCES "trains" ("train_number") ON DELETE CASCADE;

ALTER TABLE "tickets" ADD FOREIGN KEY ("email") REFERENCES "users" ("email") ON DELETE CASCADE;

ALTER TABLE "payments" ADD FOREIGN KEY ("ticket_id") REFERENCES "tickets" ("ticket_id") ON DELETE CASCADE;
