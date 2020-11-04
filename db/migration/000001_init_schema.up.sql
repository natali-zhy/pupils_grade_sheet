CREATE TABLE "Schools" (
  "Id" bigserial PRIMARY KEY,
  "Name" varchar NOT NULL,
  "NumberOfClasses" int NOT NULL,
  "CallCenter" varchar NOT NULL,
  "Address" varchar NOT NULL,
  "CreatedAt" varchar NOT NULL,
  "UpdatedAt" varchar NOT NULL
);

CREATE TABLE "Classes" (
  "Id" bigserial PRIMARY KEY,
  "SchoolId" bigserial NOT NULL,
  "Name" varchar NOT NULL,
  "NumberOfPupils" varchar NOT NULL,
  "CreateAt" varchar NOT NULL,
  "UpdatedAt" varchar NOT NULL
);

CREATE TABLE "Pupils" (
  "Id" bigserial PRIMARY KEY,
  "ClassId" bigserial NOT NULL,
  "Name" varchar NOT NULL,
  "Surname" varchar NOT NULL,
  "Patronymic" varchar NOT NULL,
  "Gender" varchar NOT NULL,
  "Address" varchar NOT NULL,
  "CreateAt" varchar NOT NULL,
  "UpdatedAt" varchar NOT NULL
);

CREATE TABLE "Scores" (
  "Id" bigserial PRIMARY KEY,
  "SubjectId" bigserial NOT NULL,
  "PupilId" bigserial NOT NULL,
  "Score" varchar NOT NULL,
  "CreateAt" varchar NOT NULL,
  "UpdatedAt" varchar NOT NULL
);

CREATE TABLE "Subject" (
  "Id" bigserial PRIMARY KEY,
  "Name" varchar NOT NULL,
  "CreateAt" varchar NOT NULL,
  "UpdatedAt" varchar NOT NULL
);

ALTER TABLE "Classes" ADD FOREIGN KEY ("SchoolId") REFERENCES "Schools" ("Id");

ALTER TABLE "Pupils" ADD FOREIGN KEY ("ClassId") REFERENCES "Classes" ("Id");

ALTER TABLE "Scores" ADD FOREIGN KEY ("PupilId") REFERENCES "Pupils" ("Id");

ALTER TABLE "Scores" ADD FOREIGN KEY ("SubjectId") REFERENCES "Subject" ("Id");
