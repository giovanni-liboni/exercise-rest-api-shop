INSERT INTO users(firstname, lastname, email, password, username, role) VALUES ('Joy','Halvorson','kayla.hilpert@gmail.com','$2a$10$ABLI.6bEhm5FyVSBlsuTWuzJ.Mfcsh2E9cAddatiH9G1A7wasTjn6','keeling.else','user'),
                         ('Ona','Orn','ereilly@gmail.com','$2a$10$ABLI.6bEhm5FyVSBlsuTWuzJ.Mfcsh2E9cAddatiH9G1A7wasTjn6','parker.annie','user'),
                         ('Mabelle','Nienow','dawn47@hotmail.com','$2a$10$ABLI.6bEhm5FyVSBlsuTWuzJ.Mfcsh2E9cAddatiH9G1A7wasTjn6','forrest75','user'),
                         ('Leopoldo','Romaguera','nathaniel69@hotmail.com','$2a$10$ABLI.6bEhm5FyVSBlsuTWuzJ.Mfcsh2E9cAddatiH9G1A7wasTjn6','hettie48','user'),
                         ('Mckenna','Kuhn','botsford.carlee@yahoo.com','$2a$10$ABLI.6bEhm5FyVSBlsuTWuzJ.Mfcsh2E9cAddatiH9G1A7wasTjn6','jack.sauer','user'),
                          ('Test','Kuhn','botsford.carlee@yahoo.com','$2a$10$ABLI.6bEhm5FyVSBlsuTWuzJ.Mfcsh2E9cAddatiH9G1A7wasTjn6','test','user'),
                         ('Admin','Admin','admin@shop.com','$2a$10$zniXjF0e3Q3mc/AESAwUZ.0nPHUUI2dVjbrx9JW7muKAgjqSn1dDe','admin','admin');

INSERT INTO items(id, name, producer, description, price, category) VALUES (1, 'The Misty Cup', 'Beier Ltd','Et sunt culpa unde distinctio quos.',244.30, 'garden'),
                           (2, 'The Begging Jug', 'Parker, Hyatt and Kris','Quos vel ut esse incidunt minima minima quae.',302.10, 'home'),
                           (3, 'The Expensive Flower','Kutch Ltd','Earum aliquid deleniti beatae quibusdam inventore itaque velit voluptas.',110.13, 'electronic'),
                           (4, 'The Challenging Stove Salon' ,'Wisozk-Larson','Quae quis laborum odio provident.',13.20, 'garden'),
                           (5, 'The Performing Window Boutique' ,'Dickinson, Collins and Cremin','Enim provident velit blanditiis ut exercitationem.',213.41, 'home');

INSERT INTO orders(user_id, payment_method, total_price, status, payment_id) VALUES (1,'card', 244.30,'created', '7548759847598437'),
                                                                          (2,'stripe', 302.10,'paid', '4324234234'),
                                                                          (3,'card', 110.13,'created', '432423423423'),
                                                                          (2,'paypal', 13.20,'paid', '532525454545'),
                                                                          (1,'card', 213.41,'created', '232342342324234'),
                                                                        (6,'card', 110.13,'created', '432423423423'),
                                                                        (6,'paypal', 13.20,'paid', '532525454545'),
                                                                        (6,'card', 213.41,'created', '232342342324234');

INSERT  INTO orders_items(order_id, item_id, price) VALUES (1, 1, 244.30), (1, 2, 302.10), (1, 3, 110.13), (2, 4, 13.20), (2, 5, 213.41), (6, 1, 244.30), (3, 2, 302.10), (4, 3, 110.13), (5, 4, 13.20), (5, 5, 213.41);