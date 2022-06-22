CREATE OR REPLACE FUNCTION set_update_time()
RETURNS trigger AS $$
    BEGIN
        NEW.updated_at = NOW();
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

-- Users
DROP TRIGGER IF EXISTS "users_updated_at_trigger" ON "users";
CREATE TRIGGER "users_updated_at_trigger"
BEFORE UPDATE ON "users"
FOR EACH ROW
EXECUTE PROCEDURE set_update_time();

-- Events
DROP TRIGGER IF EXISTS "events_updated_at_trigger" ON "events";
CREATE TRIGGER "events_updated_at_trigger"
BEFORE UPDATE ON "events"
FOR EACH ROW
EXECUTE PROCEDURE set_update_time();
