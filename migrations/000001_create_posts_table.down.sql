CREATE TABLE "adverts" {
    "id" bigserial PRIMARY KEY,
    "title" varchar NOT NULL,
    "photos" varchar[] NOT NULL,
    "price"bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
}