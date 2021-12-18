INSERT INTO users(firstname, lastname, email, password, username, role) VALUES ('Joy','Halvorson','kayla.hilpert@gmail.com','7b4516dcd757b4cb40bc7cb23bec9534e32c9c39','keeling.else','user'),
                         ('Ona','Orn','ereilly@gmail.com','bf2039937cc6b1240a2e4e5e99bc46ee2e961ed0','parker.annie','user'),
                         ('Mabelle','Nienow','dawn47@hotmail.com','af46fcd85eaa2fcbea9762ffc0a7950565280fd9','forrest75','user'),
                         ('Leopoldo','Romaguera','nathaniel69@hotmail.com','3dfe955fdb4fddd5c1e187022d986413135219cd','hettie48','user'),
                         ('Mckenna','Kuhn','botsford.carlee@yahoo.com','3d549aa0d7d2dfd7f4d1e8ae00a5691e4c2c6828','jack.sauer','user'),
                         ('Admin','Admin','admin@shop.com','3d549aa0d7d2dfd7f4d1e8ae00a5691e4c2c6828','admin','admin');

INSERT INTO items(id, name, producer, description, price, category) VALUES (1, 'The Misty Cup', 'Beier Ltd','Et sunt culpa unde distinctio quos.',244.30, 'garden'),
                           (2, 'The Begging Jug', 'Parker, Hyatt and Kris','Quos vel ut esse incidunt minima minima quae.',302.10, 'home'),
                           (3, 'The Expensive Flower','Kutch Ltd','Earum aliquid deleniti beatae quibusdam inventore itaque velit voluptas.',110.13, 'electronic'),
                           (4, 'The Challenging Stove Salon' ,'Wisozk-Larson','Quae quis laborum odio provident.',13.20, 'garden'),
                           (5, 'The Performing Window Boutique' ,'Dickinson, Collins and Cremin','Enim provident velit blanditiis ut exercitationem.',213.41, 'home');

INSERT INTO orders(user_id, payment_method, total_price, status, payment_id) VALUES (1,'card', 244.30,'paid', '7548759847598437'),
                                                                          (2,'stripe', 302.10,'paid', '4324234234'),
                                                                          (3,'card', 110.13,'paid', '432423423423'),
                                                                          (2,'paypal', 13.20,'paid', '532525454545'),
                                                                          (1,'card', 213.41,'paid', '232342342324234');

INSERT  INTO orders_items(order_id, item_id, price) VALUES (1, 1, 244.30), (1, 2, 302.10), (1, 3, 110.13), (2, 4, 13.20), (2, 5, 213.41);