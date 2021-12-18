INSERT INTO users(firstname, surname, email, password, username, role) VALUES ('Joy','Halvorson','kayla.hilpert@gmail.com','7b4516dcd757b4cb40bc7cb23bec9534e32c9c39','keeling.else','user'),
                         ('Ona','Orn','ereilly@gmail.com','bf2039937cc6b1240a2e4e5e99bc46ee2e961ed0','parker.annie','user'),
                         ('Mabelle','Nienow','dawn47@hotmail.com','af46fcd85eaa2fcbea9762ffc0a7950565280fd9','forrest75','user'),
                         ('Leopoldo','Romaguera','nathaniel69@hotmail.com','3dfe955fdb4fddd5c1e187022d986413135219cd','hettie48','user'),
                         ('Mckenna','Kuhn','botsford.carlee@yahoo.com','3d549aa0d7d2dfd7f4d1e8ae00a5691e4c2c6828','jack.sauer','user'),
                         ('Admin','Admin','admin@shop.com','3d549aa0d7d2dfd7f4d1e8ae00a5691e4c2c6828','admin','admin');

INSERT INTO items(id, producer, description, price) VALUES (1,'Beier Ltd','Et sunt culpa unde distinctio quos.',244.30),
                           (2,'Parker, Hyatt and Kris','Quos vel ut esse incidunt minima minima quae.',302.10),
                           (3,'Kutch Ltd','Earum aliquid deleniti beatae quibusdam inventore itaque velit voluptas.',110.13),
                           (4,'Wisozk-Larson','Quae quis laborum odio provident.',13.20),
                           (5,'Dickinson, Collins and Cremin','Enim provident velit blanditiis ut exercitationem.',213.41);

INSERT INTO orders(userId, payment_method, total_price, status) VALUES (1,'card', 244.30,'paid'),
                                                                          (2,'card', 302.10,'paid'),
                                                                          (3,'card', 110.13,'paid'),
                                                                          (2,'card', 13.20,'paid'),
                                                                          (1,'card', 213.41,'paid');

INSERT  INTO orders_items(orderId, itemId, price) VALUES (1, 1, 244.30), (1, 2, 302.10), (1, 3, 110.13), (2, 4, 13.20), (2, 5, 213.41);