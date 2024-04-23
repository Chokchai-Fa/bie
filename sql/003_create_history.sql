CREATE TABLE "history" (
	"history_id" serial NOT NULL PRIMARY KEY,
	"permanent" boolean NOT NULL,
	"input_digits" varchar NOT NULL,
	"sorted_input_digits" varchar NOT NULL,
	"output_number" int2 NOT NULL,
	"difficulty" varchar NOT NULL,
	"start_time" timestamptz NOT NULL,
	"finish_time" timestamptz NULL,
	"email" varchar NOT NULL
);

alter table "history" add foreign key ("email") references "User" ("email");
