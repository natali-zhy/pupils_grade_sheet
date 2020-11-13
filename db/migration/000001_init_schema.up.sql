CREATE TABLE "schools" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "numberofclasses" int NOT NULL,
  "callcenter" varchar NOT NULL,
  "address" varchar NOT NULL,
  "createdat" timestamp NOT NULL,
  "updatedat" timestamp NOT NULL
);

CREATE TABLE "classes" (
  "id" bigserial PRIMARY KEY,
  "schoolid" int NOT NULL,
  "name" varchar NOT NULL,
  "numberofpupils" integer NOT NULL,
  "createdat" timestamp NOT NULL,
  "updatedat" timestamp NOT NULL
);

CREATE TABLE "pupils" (
  "id" bigserial PRIMARY KEY,
  "classid" int NOT NULL,
  "name" varchar NOT NULL,
  "surname" varchar NOT NULL,
  "patronymic" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "address" varchar NOT NULL,
  "createdat" timestamp NOT NULL,
  "updatedat" timestamp NOT NULL
);

CREATE TABLE "scores" (
  "id" bigserial PRIMARY KEY,
  "subjectid" int NOT NULL,
  "pupilid" int NOT NULL,
  "score" integer NOT NULL,
  "createdat" timestamp NOT NULL,
  "updatedat" timestamp NOT NULL
);

CREATE TABLE "subject" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "createdat" timestamp NOT NULL,
  "updatedat" timestamp NOT NULL
);

ALTER TABLE "classes" ADD FOREIGN KEY ("schoolid") REFERENCES "schools" ("id");

ALTER TABLE "pupils" ADD FOREIGN KEY ("classid") REFERENCES "classes" ("id");

ALTER TABLE "scores" ADD FOREIGN KEY ("pupilid") REFERENCES "pupils" ("id");

ALTER TABLE "scores" ADD FOREIGN KEY ("subjectid") REFERENCES "subject" ("id");
