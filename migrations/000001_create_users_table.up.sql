
CREATE TABLE users_v2 
  ( 
     sha1               VARCHAR PRIMARY KEY NOT NULL, 
     id                 BIGSERIAL NOT NULL UNIQUE, 
     email              VARCHAR NOT NULL UNIQUE, 
     encrypted_password VARCHAR NOT NULL 
  ); 

COMMIT; 

CREATE TABLE bookings 
  ( 
     booking_key          VARCHAR PRIMARY KEY NOT NULL, --sha1 + property_internal_id 
     sha1                 VARCHAR NOT NULL, 
     email                VARCHAR NOT NULL, 
     phone                VARCHAR NOT NULL, 
     NAME                 VARCHAR NOT NULL, 
     surname              VARCHAR NOT NULL, 
     payment_id           VARCHAR NOT NULL, 
     status               VARCHAR NOT NULL, -- active / canceled 
     timestamp            VARCHAR NOT NULL, 
     property_internal_id VARCHAR NOT NULL,
     date_from            VARCHAR NOT NULL,
     date_to              VARCHAR NOT NULL,
     property_id          VARCHAR NOT NULL,
     hotel_internal_id    VARCHAR NOT NULL
  ); 

COMMIT; 

CREATE TABLE likes 
  ( 
     sha1              VARCHAR NOT NULL,
     hotel_internal_id VARCHAR NOT NULL 
  ); 

COMMIT; 

CREATE TABLE hotels_v2 
  ( 
     hotel_id          VARCHAR PRIMARY KEY, 
     hotel_name        VARCHAR NOT NULL, 
     hotel_internal_id VARCHAR NOT NULL UNIQUE, 
     provider          VARCHAR NOT NULL, 
     structure_type    VARCHAR NOT NULL, 
     min_nights        INT NOT NULL, 
     max_nights        INT NOT NULL, 
     stars             INT NOT NULL, 
     images            JSON NOT NULL, 
     start_date        VARCHAR NOT NULL, 
     active_day_period VARCHAR NOT NULL, 
     latitude          VARCHAR NOT NULL, 
     longitude         VARCHAR NOT NULL, 
     address           JSON NOT NULL, 
     description       VARCHAR NOT NULL, 
     active            BOOLEAN NOT NULL, 
     wifi              BOOLEAN NOT NULL, 
     breakfast         BOOLEAN NOT NULL, 
     parking           BOOLEAN NOT NULL, 
     pool              BOOLEAN NOT NULL, 
     playground        BOOLEAN NOT NULL, 
     garden            BOOLEAN NOT NULL, 
     check_in          VARCHAR NOT NULL, 
     check_out         VARCHAR NOT NULL 
  ); 

COMMIT; 

CREATE TABLE properties 
  ( 
     hotel_internal_id    VARCHAR NOT NULL, 
     property_id          VARCHAR NOT NULL, 
     property_internal_id VARCHAR NOT NULL UNIQUE, 
     property_name        VARCHAR NOT NULL, 
     price                FLOAT NOT NULL, 
     currency             VARCHAR NOT NULL, 
     nights               INT NOT NULL, 
     image                VARCHAR NOT NULL, 
     date_from            VARCHAR NOT NULL, 
     date_to              VARCHAR NOT NULL, 
     description          VARCHAR NOT NULL, 
     active               BOOLEAN NOT NULL, 
     rooms_number         INT NOT NULL, 
     beds_number          VARCHAR NOT NULL, 
     size_m               FLOAT NOT NULL, 
     lock                 VARCHAR NOT NULL, -- sha1 is written here when booking is in progress 
     timestamp            VARCHAR NOT NULL,
     capacity             INT NOT NULL 
  ); 

COMMIT; 

CREATE TABLE geoip2 
  ( 
     geoname_id       VARCHAR PRIMARY KEY, 
     country_iso_code VARCHAR NOT NULL, 
     country_name     VARCHAR NOT NULL, 
     city_name        VARCHAR NOT NULL 
  ); 

COMMIT; 



-- CREATE TABLE hotels (
--     providerPropertyID             VARCHAR NOT NULL PRIMARY KEY,                          
-- 	addresses                      JSON,                        
-- 	attributes                     JSON,                    
-- 	billingCurrencyCode            VARCHAR,                           
-- 	contacts                       JSON,                       
-- 	contents                       JSON,                           
-- 	currencyCode                   VARCHAR,                          
-- 	hideAddress                    BOOLEAN,                         
-- 	hideExactLocation              BOOLEAN,                            
-- 	inventorySettings              JSON,               
-- 	isVacationRental               BOOLEAN,                            
-- 	latitude                       VARCHAR NOT NULL,                           
-- 	longitude                      VARCHAR NOT NULL,                           
-- 	name                           VARCHAR,                           
-- 	policies                       JSON,                   
-- 	propertyCollectedMandatoryFees JSON,
-- 	provider                       VARCHAR NOT NULL,                           
-- 	providerPropertyURL            VARCHAR,                           
-- 	ratings                        JSON,                        
-- 	structureType                  VARCHAR,                           
-- 	taxes                          JSON,         
-- 	timeZone                       VARCHAR
-- )