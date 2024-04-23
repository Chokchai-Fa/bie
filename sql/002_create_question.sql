CREATE TABLE "Difficulty" (
  "difficulty_code" varchar PRIMARY KEY,
  "difficulty_description" varchar
);

CREATE TABLE "Question" (
  "question_id" serial PRIMARY KEY,
  "question_flag" varchar,
  "min_difficulty" varchar,
  "question" varchar,
  "question_result" varchar
);

CREATE TABLE "Solution" (
  "solution_id" serial PRIMARY KEY,
  "question_id" int,
  "solution_result" varchar,
  "solution_score" int,
  "difficulty_code" varchar,
  "solution" varchar
);


ALTER TABLE "Question" ADD FOREIGN KEY ("min_difficulty") REFERENCES "Difficulty"("difficulty_code");
ALTER TABLE "Solution" ADD FOREIGN KEY ("question_id") REFERENCES "Question"("question_id");
ALTER TABLE "Solution" ADD FOREIGN KEY ("difficulty_code") REFERENCES "Difficulty"("difficulty_code");
