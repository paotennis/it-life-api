-- Users
DROP   TABLE    IF     EXISTS "public"."users" CASCADE;
CREATE TABLE    IF NOT EXISTS "public"."users" (
    "id"                    BIGSERIAL                NOT NULL,
    "created_at"            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at"            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "firebase_uid"          TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("firebase_uid") > 0),
    "name"                  TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("name") > 0),
    "email"                 TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("email") > 0),
    "age"                   INTEGER                  NOT NULL CHECK ("age" > 0),
    "role"                  TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("role") > 0),
    "company"               TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("company") > 0),
    PRIMARY KEY ("id"),
    UNIQUE ("firebase_uid")
);

-- Events
DROP   TABLE    IF     EXISTS "public"."events" CASCADE;
CREATE TABLE    IF NOT EXISTS "public"."events" (
    "id"                    BIGSERIAL                NOT NULL,
    "created_at"            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at"            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "user_id"               BIGINT                   NOT NULL,
    "name"                  TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("name") > 0),
    "description"           TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("description") > 0),
    "starts_at"             TIMESTAMP WITH TIME ZONE NOT NULL,
    "ends_at"               TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON UPDATE CASCADE ON DELETE CASCADE
);
