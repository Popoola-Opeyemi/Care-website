CREATE TABLE "pages" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "route" varchar(150) NOT NULL,
  "uuid" varchar(20) NOT NULL,
  "attributes" jsonb NOT NULL,
  "created_at" date NOT NULL DEFAULT LOCALTIMESTAMP
);

CREATE TABLE "users" (
  "id" serial not null PRIMARY KEY,
  "role_id" bigint NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" date NOT NULL DEFAULT LOCALTIMESTAMP
);

CREATE TABLE "services" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(250),
  "uuid" varchar(20) not null,
  "url" varchar(200) not NULL,
  "content" text,
  "image" jsonb NOT NULL,
  "created_at" date NOT NULL DEFAULT LOCALTIMESTAMP
);

CREATE TABLE "faq" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(250),
  "content" text,
  "created_at" date NOT NULL DEFAULT LOCALTIMESTAMP
);

CREATE TABLE "testimonials" (
  "id" bigserial PRIMARY KEY,
  "author" varchar(250),
  "content" text,
  "created_at" date NOT NULL DEFAULT LOCALTIMESTAMP
);

CREATE TABLE "contacts"(
  "id" bigserial PRIMARY KEY,
  "name" varchar(250) not null,
  "email" varchar(100) not null,
  "mobile_number" varchar(50) not null,
  "message" text not null,
  "created_at" date NOT NULL DEFAULT LOCALTIMESTAMP
);

CREATE TABLE "settings"(
  "id" serial not null PRIMARY key,
  "name" varchar(250),
  "url" varchar(200) not NULL,
  "uuid" varchar(20) NOT NULL,
  "attributes" jsonb NOT NULL,
  "created_at" date NOT NULL DEFAULT LOCALTIMESTAMP
);

CREATE TABLE "logger"(
  "id" serial not null PRIMARY key,
  "admin_id" int not null references users(id),
  "action" varchar not null,
  "data" json,
  "user_agent" text,
  "ip_address" character varying(200),
  "created_at" timestamp not null default current_timestamp
);



INSERT INTO "users" 
	VALUES (1, 0, 'o.davies@pbgdigital.co.uk','$2y$12$qEPrsVZe19ck8Knf/FRrR.rvWUSl7h1dfFsTapDVNl79X4edGDhoy');
  
insert into "users" values (2, 0, 'ericka.murphy@pbgdigital.co.uk', '$2y$12$weMCmrdVACWk4um7yXO93uXDkBF/OldFmA9KjrHeF25VEt9/86C0G');
