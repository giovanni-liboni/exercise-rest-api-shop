
# Stored Procedures for the item table
DROP PROCEDURE IF EXISTS `sp_GetItems`;
CREATE PROCEDURE sp_GetItems()
BEGIN
    select * from items;
END;

DROP PROCEDURE IF EXISTS sp_GetItem;
CREATE PROCEDURE sp_GetItem(IN idItem int)
BEGIN
    select * from items where id=idItem;
END;

DROP PROCEDURE IF EXISTS sp_CreateItem;
CREATE PROCEDURE sp_CreateItem(IN name varchar(255), IN producer varchar(255), IN description varchar(255), IN price float, IN category int)
BEGIN
    insert into items (name, producer, description, price, category) values (name, producer, description, price, category);
END;

DROP PROCEDURE IF EXISTS sp_UpdateItem;
CREATE PROCEDURE sp_UpdateItem(IN idItem int, IN name varchar(255), IN producer varchar(255), IN description varchar(255), IN price float, IN category int)
BEGIN
    update items set name=name, producer=producer, description=description, price=price, category=category where id=idItem;
END;

DROP PROCEDURE IF EXISTS sp_DeleteItem;
CREATE PROCEDURE sp_DeleteItem(IN idItem int)
BEGIN
    delete from items where id=idItem;
END;

# Store Procedures for the users table

DROP PROCEDURE IF EXISTS sp_GetUsers;
CREATE PROCEDURE sp_GetUsers()
BEGIN
    select * from users;
END;

DROP PROCEDURE IF EXISTS sp_GetUser;
CREATE PROCEDURE sp_GetUser(IN idUser int)
BEGIN
    select * from users where id=idUser;
END;

DROP PROCEDURE IF EXISTS sp_CreateUser;
CREATE PROCEDURE sp_CreateUser(IN firstName varchar(255), IN lastName varchar(255), IN email varchar(255), IN password varchar(255))
BEGIN
    insert into users (firstname, lastname, email, password, username, role) VALUES (firstName, lastName, email, password, email, 'user');
END;

DROP PROCEDURE IF EXISTS sp_UpdateUser;
CREATE PROCEDURE sp_UpdateUser(IN idUser int, IN firstName varchar(255), IN lastName varchar(255), IN email varchar(255), IN password varchar(255), IN username varchar(255), IN role varchar(255))
BEGIN
    update users set firstname=firstName, lastname=lastName, email=email, password=password, username=username, role=role where id=idUser;
END;

DROP PROCEDURE IF EXISTS sp_DeleteUser;
CREATE PROCEDURE sp_DeleteUser(IN idUser int)
BEGIN
    delete from users where id=idUser;
END;

# Procedures for the ordes table

DROP PROCEDURE IF EXISTS sp_GetOrders;
CREATE PROCEDURE sp_GetOrders()
BEGIN
    select * from orders;
END;

DROP PROCEDURE IF EXISTS sp_GetOrder;
CREATE PROCEDURE sp_GetOrder(IN idOrder int)
BEGIN
    select * from orders where id=idOrder;
END;

DROP PROCEDURE IF EXISTS sp_CreateOrder;
CREATE PROCEDURE sp_CreateOrder(IN userID int, IN paymentMethod varchar(255), IN total float, IN status varchar(255), paymentMethodID varchar(255))
BEGIN
    insert into orders (user_id, payment_method, payment_id, total_price, status) VALUES  (userID, paymentMethod, paymentMethodID, total, status);
END;

DROP PROCEDURE IF EXISTS sp_UpdateOrder;
CREATE PROCEDURE sp_UpdateOrder(IN idOrder int, IN userID int, IN paymentMethod varchar(255), IN total float, IN status varchar(255), paymentMethodID varchar(255))
BEGIN
    update orders set user_id=userID, payment_method=paymentMethod, payment_id=paymentMethodID, total_price=total, status=status where id=idOrder;
END;