INSERT INTO clients (client_id, client_name) VALUES ('1', 'SuperAdmin');
INSERT INTO clients (client_id, client_name) VALUES ('2', 'Admin');
INSERT INTO clients (client_id, client_name) VALUES ('3', 'Pilot');
INSERT INTO clients (client_id, client_name) VALUES ('4', 'Tailor');
INSERT INTO clients (client_id, client_name) VALUES ('5', 'User');

INSERT INTO coupons (coupon_id,description,expiry,only_new,only_first,only_app, min_amount, discount) VALUES ('FIRST100','Rs. 100 OFF on First Order',1477898999,false,true,false, 300, 100);
INSERT INTO coupons (coupon_id,description,expiry,only_new,only_first,only_app, min_amount, discount) VALUES ('SUPERZIG','FLAT 10% OFF on orders more than RS. 300',1477898999,true,false,false, 300, 10);
INSERT INTO coupons (coupon_id,description,expiry,only_new,only_first,only_app, min_amount, discount) VALUES ('APP10','FLAT 10% OFF on all orders using app',1477898999,true,false,true, 300, 10);

INSERT INTO slots (slot_id, slot_name) VALUES ('slot1', '9:00 AM to 11:00 AM');
INSERT INTO slots (slot_id, slot_name) VALUES ('slot2', '11:00 AM to 1:00 PM');
INSERT INTO slots (slot_id, slot_name) VALUES ('slot3', '1:00 PM to 3:00 PM');
INSERT INTO slots (slot_id, slot_name) VALUES ('slot4', '3:00 PM to 5:00 PM');
INSERT INTO slots (slot_id, slot_name) VALUES ('slot5', '5:00 PM to 7:00 PM');
INSERT INTO slots (slot_id, slot_name) VALUES ('slot6', '7:00 PM to 9:00 PM');

INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f01', 'Raymonds', 'M', 'Shirts', '1', '/img/fabrics/1.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f02', 'Raymonds', 'M', 'Suits', '2', '/img/fabrics/2.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f03', 'Raymonds', 'F', 'Shirts', '3', '/img/fabrics/6.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f04', 'Raymonds', 'F', 'Suits', '1', '/img/fabrics/7.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f05', 'Calvin Klein', 'M', 'Shirts', '1', '/img/fabrics/8.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f06', 'Calvin Klein', 'M', 'Suits', '1', '/img/fabrics/20.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f07', 'Calvin Klein', 'F', 'Shirts', '3', '/img/fabrics/3.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f08', 'Calvin Klein', 'F', 'Suits', '3', '/img/fabrics/3.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f09', 'Arvind', 'M', 'Shirts', '2', '/img/fabrics/5.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f10', 'Arvind', 'M', 'Suits', '1', '/img/fabrics/24.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f11', 'Arvind', 'F', 'Shirts', '1', '/img/fabrics/19.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
INSERT INTO fabrics (fabric_id, brand, gender, category, quality, img, quantity, rate, disc_rate, description) VALUES ('f12', 'Arvind', 'F', 'Suits', '1', '/img/fabrics/11.jpg', 100, 100, 89, 'A textured woven, usually printed cotton fabric that was popular in the 30s-40s and 50s as an interiors fabric. The prints were often large vines, leaves and florals.');
-- INSERT INTO credentials (mobileno, client_id, password) VALUES ('1111111111', '3', 'mithun');
-- INSERT INTO credentials (mobileno, client_id, password) VALUES ('1111111112', '3', 'mithun');
--
-- INSERT INTO username_map (username, mobileno) VALUES ('pilot1', '1111111111');
-- INSERT INTO username_map (username, mobileno) VALUES ('pilot2', '1111111112');
--
-- INSERT INTO appointments (appointment_id, appointment_day, slot_id, username) VALUES ('app1', 1477333800, 'slot1', 'pilot1');
-- INSERT INTO appointments (appointment_id, appointment_day, slot_id, username) VALUES ('app2', 1477420200, 'slot1', 'pilot2');
-- INSERT INTO appointments (appointment_id, appointment_day, slot_id, username) VALUES ('app3', 1477420200, 'slot1', 'pilot1');
-- INSERT INTO appointments (appointment_id, appointment_day, slot_id, username) VALUES ('app4', 1477506600, 'slot1', 'pilot1');

INSERT INTO options (option_key, option_name, option_code, price) VALUES ('101', 'Slim Fit', 'slim_fit' , 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('102', 'Loose Fit', 'loose_fit', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('103', 'Normal Fit', 'normal_fit', 0);

INSERT INTO options (option_key, option_name, option_code, price) VALUES ('201', 'Long Sleeve', 'long_sleeve', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('202', 'Roll-Up Sleeve', 'roll_up_sleeve', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('203', 'Short Sleeve', 'short_sleeve', 0);

INSERT INTO options (option_key, option_name, option_code, price) VALUES ('301', 'Business Classic', 'business_classic', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('302', 'Business Superior', 'business_superior', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('303', 'Button-Down Classic', 'button_down_classic', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('304', 'Button-Down Modern', 'button_down_modern', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('305', 'Club', 'club', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('306', 'Club Modern', 'club_modern', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('307', 'Cut-Away Classic', 'cut_away_classic', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('308', 'Cut-Away Casual', 'cut_away_casual', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('309', 'Cut-Away Extreme', 'cut_away_extreme', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('310', 'Cut-Away Modern', 'cut_away_modern', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('311', 'Cut-Away Superior', 'cut_away_superior', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('312', 'Cut-Away Two-Button', 'cut_away_twobutton', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('313', 'Turndown Superior', 'turndown_superior', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('314', 'Tab', 'tab', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('315', 'Wing Collar', 'wing_collar', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('316', 'Mao', 'mao', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('317', 'Pin', 'pin', 0);

INSERT INTO options (option_key, option_name, option_code, price) VALUES ('401', 'Single Button Rounded', 'single_button_rounded', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('402', 'Single Button Beveled', 'single_button_beveled', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('403', 'Single Button Straight', 'single_button_straight', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('404', 'Convertible Rounded', 'convertible_rounded', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('405', 'Double Button Rounded', 'double_button_rounded', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('406', 'Double Button Beveled', 'double_button_beveled', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('407', 'French Cuff', 'french_cuff', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('408', 'Link Cuff', 'link_cuff', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('409', 'Narrow', 'narrow', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('410', 'Casual Single Button', 'casual_single_button', 0);

INSERT INTO options (option_key, option_name, option_code, price) VALUES ('501', 'With Placket', 'with_placket', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('502', 'Without Placket', 'without_placket', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('503', 'Hidden Buttons', 'hidden_buttons', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('504', 'Narrow Placket', 'narrrow_placket', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('505', 'Tuxedo Placket', 'tuxedo_placket', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('506', 'White Tie Placket for Studs', 'white_tie_placket_for_studs', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('507', 'White Tie Placket', 'white_tie_placket', 0);

INSERT INTO options (option_key, option_name, option_code, price) VALUES ('601', 'No Pocket', 'no_pocket', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('602', 'Left', 'left', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('603', 'Both', 'both', 0);

INSERT INTO options (option_key, option_name, option_code, price) VALUES ('701', 'Straight', 'straight', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('702', 'Skewed', 'skewed', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('703', 'V-Shaped', 'v_shaped', 0);

INSERT INTO options (option_key, option_name, option_code, price) VALUES ('801', 'No', 'no', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('802', 'Yes', 'yes', 0);

INSERT INTO options (option_key, option_name, option_code, price) VALUES ('901', 'No Back Details', 'no_back_details', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('902', 'Center Folds', 'center_folds', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('903', 'Side Folds', 'side_folds', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('904', 'Back Darts', 'back_darts', 0);


INSERT INTO options (option_key, option_name, option_code, price) VALUES ('1001', 'Classic', 'classic', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('1002', 'Modern', 'modern', 0);
INSERT INTO options (option_key, option_name, option_code, price) VALUES ('1003', 'Straight', 'straight', 0);
