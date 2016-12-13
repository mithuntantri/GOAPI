CREATE TABLE "profile" (
	"mobileno" varchar(10) NOT NULL,
	"email_id" varchar(127) NOT NULL,
	"client_id" int NOT NULL,
	"first_name" varchar(127) NOT NULL,
	"last_name" varchar(127) NOT NULL,
	"gender" varchar(1) NOT NULL,
	"referral_id" varchar(30) NOT NULL,
	"referred_id" varchar(20) NOT NULL,
	"wallet_id" varchar(20) NOT NULL,
	CONSTRAINT profile_pk PRIMARY KEY ("mobileno")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "clients" (
	"client_id" int NOT NULL,
	"client_name" varchar(30) NOT NULL,
	CONSTRAINT clients_pk PRIMARY KEY ("client_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "wallet" (
	"wallet_id" varchar(20) NOT NULL,
	"referral_credits" float4 NOT NULL,
	"profile_credits" float4 NOT NULL,
	"promo_credits" float4 NOT NULL,
	"total_credits" float4 NOT NULL,
	CONSTRAINT wallet_pk PRIMARY KEY ("wallet_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "referral" (
	"referral_id" varchar(30) NOT NULL,
	"referral_count" int NOT NULL,
	"wallet_id" varchar(20) NOT NULL,
	CONSTRAINT referral_pk PRIMARY KEY ("referral_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "fbid_map" (
	"fb_id" varchar(127) NOT NULL,
	"mobileno" varchar(10) NOT NULL,
	CONSTRAINT fbid_map_pk PRIMARY KEY ("fb_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "measurements" (
	"measurement_id" varchar(20) NOT NULL,
	"mobileno" varchar(10) NOT NULL,
	"name" varchar(20) NOT NULL,
	"units" varchar(2) NOT NULL,
	"neck" int NOT NULL,
	"chest" int NOT NULL,
	"waist" int NOT NULL,
	"hip" int NOT NULL,
	"length" int NOT NULL,
	"shoulder" int NOT NULL,
	"sleeve" int NOT NULL,
	"is_default" BOOLEAN NOT NULL,
	CONSTRAINT measurements_pk PRIMARY KEY ("measurement_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "options" (
	"option_key" varchar(5) NOT NULL,
	"option_name" varchar(30) NOT NULL,
	"option_code" varchar(30) NOT NULL,
	"price" varchar(10) NOT NULL,
	CONSTRAINT options_pk PRIMARY KEY ("option_key")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "blouse" (
	"option_key" varchar(5) NOT NULL,
	"option_name" varchar(30) NOT NULL,
	"option_category" varchar(30) NOT NULL,
	"option_type" varchar(30) NOT NULL,
	"img" varchar(127) NOT NULL,
	"price" float4 NOT NULL,
	"disable_list" varchar(256) NOT NULL,
	CONSTRAINT blouse_pk PRIMARY KEY ("option_key")
) WITH (
	OIDS=FALSE
);

CREATE TABLE "orders" (
	"order_id" varchar(30) NOT NULL,
	"order_time" bigint NOT NULL,
	"mobileno" varchar(10) NOT NULL,
	"product_id" varchar(30) NOT NULL,
	"fabric_id" varchar(20) NOT NULL,
	"design_hash" varchar(30) NOT NULL,
	"measurement_id" varchar(8) NOT NULL,
	"fabric_pickup" bool NOT NULL,
	"measurement_pickup" bool NOT NULL,
	"appointment_id" varchar(30) NOT NULL,
	"address_id" varchar(8) NOT NULL,
	"coupon_id" varchar(10) NOT NULL,
	"credits" float4 NOT NULL,
	"discount" float4 NOT NULL,
	"final_price" float4 NOT NULL,
	CONSTRAINT orders_pk PRIMARY KEY ("order_id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "fabrics" (
	"fabric_id" varchar(20) NOT NULL,
	"brand" varchar(30) NOT NULL,
	"gender" varchar(1) NOT NULL,
	"category" varchar(30) NOT NULL,
	"material" varchar(30) NOT NULL,
	"quality" varchar(30) NOT NULL,
	"img1"  varchar(127) NOT NULL,
	"img2" varchar(127) NOT NULL,
	"quantity" float4 NOT NULL,
	"rate" float4 NOT NULL,
	"disc_rate" float4 NOT NULL,
	"title" varchar(127) NOT NULL,
	"description" varchar(256) NOT NULL,
	CONSTRAINT fabrics_pk PRIMARY KEY ("fabric_id")
) WITH (
	OIDS=FALSE
);


CREATE TABLE "address" (
	"address_id" varchar(30) NOT NULL,
	"mobileno" varchar(10) NOT NULL,
	"street" varchar(127) NOT NULL,
	"address" varchar(127) NOT NULL,
	"pin_code" varchar(6) NOT NULL,
	"is_default" BOOLEAN NOT NULL,
	CONSTRAINT address_pk PRIMARY KEY ("address_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "mode" (
	"payment_mode" varchar(10) NOT NULL,
	"payment_name" varchar(30) NOT NULL,
	CONSTRAINT mode_pk PRIMARY KEY ("payment_mode")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "coupons" (
	"coupon_id" varchar(10) NOT NULL,
	"description" varchar(127) NOT NULL,
	"expiry" bigint NOT NULL,
	"only_new" BOOLEAN NOT NULL,
	"only_first" BOOLEAN NOT NULL,
	"only_app" BOOLEAN NOT NULL,
	"min_amount" float4 NOT NULL,
	"discount" float4 NOT NULL,
	CONSTRAINT coupons_pk PRIMARY KEY ("coupon_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "payments" (
	"payment_id" varchar(30) NOT NULL,
	"payment_time" bigint NOT NULL,
	"order_id" varchar(30) NOT NULL,
	"payment_mode" varchar(10) NOT NULL,
	CONSTRAINT payments_pk PRIMARY KEY ("payment_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "delivery" (
	"delivery_id" varchar(30) NOT NULL,
	"order_id" varchar(30) NOT NULL,
	"deliver_time" bigint NOT NULL,
	"delivery_status" varchar(20) NOT NULL,
	CONSTRAINT delivery_pk PRIMARY KEY ("delivery_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "appointments" (
	"appointment_id" varchar(30) NOT NULL,
	"appointment_day" bigint NOT NULL,
	"slot_id" varchar(10) NOT NULL,
	"username" varchar(30) NOT NULL,
	CONSTRAINT appointments_pk PRIMARY KEY ("appointment_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "slots" (
	"slot_id" varchar(10) NOT NULL,
	"slot_name" varchar(30) NOT NULL,
	CONSTRAINT slots_pk PRIMARY KEY ("slot_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "credentials" (
	"mobileno" varchar(10) NOT NULL,
	"email_id" varchar(127) NOT NULL,
	"verified_mobile" bool NOT NULL,
	"verified_email" bool NOT NULL,
	"client_id" int NOT NULL,
	"password" varchar(255) NOT NULL,
	CONSTRAINT credentials_pk PRIMARY KEY ("mobileno")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "username_map" (
	"username" varchar(30) NOT NULL,
	"mobileno" varchar(10) NOT NULL,
	CONSTRAINT username_map_pk PRIMARY KEY ("username")
) WITH (
  OIDS=FALSE
);


ALTER TABLE "profile" ADD CONSTRAINT "profile_fk0" FOREIGN KEY ("client_id") REFERENCES "clients"("client_id");
ALTER TABLE "profile" ADD CONSTRAINT "profile_fk1" FOREIGN KEY ("referral_id") REFERENCES "referral"("referral_id");
ALTER TABLE "profile" ADD CONSTRAINT "profile_fk2" FOREIGN KEY ("wallet_id") REFERENCES "wallet"("wallet_id");



ALTER TABLE "referral" ADD CONSTRAINT "referral_fk0" FOREIGN KEY ("wallet_id") REFERENCES "wallet"("wallet_id");

ALTER TABLE "fbid_map" ADD CONSTRAINT "fbid_map_fk0" FOREIGN KEY ("mobileno") REFERENCES "credentials"("mobileno");

ALTER TABLE "measurements" ADD CONSTRAINT "measurements_fk0" FOREIGN KEY ("mobileno") REFERENCES "profile"("mobileno");


ALTER TABLE "orders" ADD CONSTRAINT "orders_fk0" FOREIGN KEY ("mobileno") REFERENCES "profile"("mobileno");
ALTER TABLE "orders" ADD CONSTRAINT "orders_fk1" FOREIGN KEY ("appointment_id") REFERENCES "appointments"("appointment_id");
ALTER TABLE "orders" ADD CONSTRAINT "orders_fk2" FOREIGN KEY ("coupon_id") REFERENCES "coupons"("coupon_id");

ALTER TABLE "address" ADD CONSTRAINT "address_fk0" FOREIGN KEY ("mobileno") REFERENCES "profile"("mobileno");



ALTER TABLE "payments" ADD CONSTRAINT "payments_fk0" FOREIGN KEY ("order_id") REFERENCES "orders"("order_id");
ALTER TABLE "payments" ADD CONSTRAINT "payments_fk1" FOREIGN KEY ("payment_mode") REFERENCES "mode"("payment_mode");

ALTER TABLE "delivery" ADD CONSTRAINT "delivery_fk0" FOREIGN KEY ("order_id") REFERENCES "orders"("order_id");

ALTER TABLE "appointments" ADD CONSTRAINT "appointments_fk0" FOREIGN KEY ("slot_id") REFERENCES "slots"("slot_id");
ALTER TABLE "appointments" ADD CONSTRAINT "appointments_fk1" FOREIGN KEY ("username") REFERENCES "username_map"("username");


ALTER TABLE "credentials" ADD CONSTRAINT "credentials_fk0" FOREIGN KEY ("client_id") REFERENCES "clients"("client_id");

ALTER TABLE "username_map" ADD CONSTRAINT "username_map_fk0" FOREIGN KEY ("mobileno") REFERENCES "credentials"("mobileno");
