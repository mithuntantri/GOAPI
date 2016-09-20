CREATE TABLE "profile" (
	"mobileno" varchar(10) NOT NULL,
	"email_id" varchar(127) NOT NULL,
	"client_id" int NOT NULL,
	"first_name" varchar(127) NOT NULL,
	"last_name" varchar(127) NOT NULL,
	"dob" DATE NOT NULL,
	"gender" varchar(1) NOT NULL,
	"address" varchar(127) NOT NULL,
	"street" varchar(127) NOT NULL,
	"pin_code" varchar(6) NOT NULL,
	"verified" BOOLEAN NOT NULL,
	"refferral_id" varchar(30) NOT NULL,
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
	"referral_credits" int NOT NULL,
	"profile_credits" int NOT NULL,
	"promo_credits" int NOT NULL,
	CONSTRAINT wallet_pk PRIMARY KEY ("wallet_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "credentials" (
	"mobileno" varchar(10) NOT NULL,
	"client_id" int NOT NULL,
	"password" varchar(255) NOT NULL,
	CONSTRAINT credentials_pk PRIMARY KEY ("mobileno")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "referral" (
	"refferral_id" varchar(30) NOT NULL,
	"referral_count" int NOT NULL,
	CONSTRAINT referral_pk PRIMARY KEY ("refferral_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "emailid_map" (
	"EmailId" varchar(127) NOT NULL,
	"mobileno" varchar(10) NOT NULL,
	CONSTRAINT emailid_map_pk PRIMARY KEY ("EmailId")
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



CREATE TABLE "fbid_map" (
	"fb_id" varchar(127) NOT NULL,
	"mobileno" varchar(10) NOT NULL,
	CONSTRAINT fbid_map_pk PRIMARY KEY ("fb_id")
) WITH (
  OIDS=FALSE
);



ALTER TABLE "profile" ADD CONSTRAINT "profile_fk0" FOREIGN KEY ("client_id") REFERENCES "clients"("client_id");
ALTER TABLE "profile" ADD CONSTRAINT "profile_fk1" FOREIGN KEY ("refferral_id") REFERENCES "referral"("refferral_id");
ALTER TABLE "profile" ADD CONSTRAINT "profile_fk2" FOREIGN KEY ("wallet_id") REFERENCES "wallet"("wallet_id");



ALTER TABLE "credentials" ADD CONSTRAINT "credentials_fk0" FOREIGN KEY ("client_id") REFERENCES "clients"("client_id");


ALTER TABLE "emailid_map" ADD CONSTRAINT "emailid_map_fk0" FOREIGN KEY ("mobileno") REFERENCES "credentials"("mobileno");

ALTER TABLE "username_map" ADD CONSTRAINT "username_map_fk0" FOREIGN KEY ("mobileno") REFERENCES "credentials"("mobileno");

ALTER TABLE "fbid_map" ADD CONSTRAINT "fbid_map_fk0" FOREIGN KEY ("mobileno") REFERENCES "credentials"("mobileno");
