ALTER TABLE "profile" DROP CONSTRAINT IF EXISTS "profile_fk0";

ALTER TABLE "profile" DROP CONSTRAINT IF EXISTS "profile_fk1";

ALTER TABLE "profile" DROP CONSTRAINT IF EXISTS "profile_fk2";

ALTER TABLE "referral" DROP CONSTRAINT IF EXISTS "referral_fk0";

ALTER TABLE "fbid_map" DROP CONSTRAINT IF EXISTS "fbid_map_fk0";

ALTER TABLE "measurements" DROP CONSTRAINT IF EXISTS "measurements_fk0";

ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_fk0";

ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_fk1";

ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_fk2";

ALTER TABLE "address" DROP CONSTRAINT IF EXISTS "address_fk0";

ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS "payments_fk0";

ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS "payments_fk1";

ALTER TABLE "delivery" DROP CONSTRAINT IF EXISTS "delivery_fk0";

ALTER TABLE "appointments" DROP CONSTRAINT IF EXISTS "appointments_fk0";

ALTER TABLE "appointments" DROP CONSTRAINT IF EXISTS "appointments_fk1";

ALTER TABLE "credentials" DROP CONSTRAINT IF EXISTS "credentials_fk0";

ALTER TABLE "username_map" DROP CONSTRAINT IF EXISTS "username_map_fk0";

DROP TABLE IF EXISTS "profile";

DROP TABLE IF EXISTS "clients";

DROP TABLE IF EXISTS "wallet";

DROP TABLE IF EXISTS "referral";

DROP TABLE IF EXISTS "fbid_map";

DROP TABLE IF EXISTS "measurements";

DROP TABLE IF EXISTS "options";

DROP TABLE IF EXISTS "orders";

DROP TABLE IF EXISTS "address";

DROP TABLE IF EXISTS "mode";

DROP TABLE IF EXISTS "coupons";

DROP TABLE IF EXISTS "payments";

DROP TABLE IF EXISTS "delivery";

DROP TABLE IF EXISTS "appointments";

DROP TABLE IF EXISTS "slots";

DROP TABLE IF EXISTS "credentials";

DROP TABLE IF EXISTS "username_map";

DROP TABLE IF EXISTS "fabrics";

DROP TABLE IF EXISTS "blouse";
