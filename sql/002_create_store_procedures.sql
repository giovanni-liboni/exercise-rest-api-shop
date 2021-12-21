# Stored Procedures for the item table
DROP PROCEDURE IF EXISTS `sp_GetItems`;
CREATE PROCEDURE sp_GetItems()
BEGIN
    select * from items;
END;

DROP PROCEDURE IF EXISTS sp_GetItem;
CREATE PROCEDURE sp_GetItem(IN idItem bigint)
BEGIN
    select * from items where id=idItem;
END;

DROP PROCEDURE IF EXISTS sp_CreateItem;
CREATE PROCEDURE sp_CreateItem(IN name varchar(255), IN producer varchar(255), IN description varchar(255), IN price float, IN category bigint)
BEGIN
    insert into items (name, producer, description, price, category) values (name, producer, description, price, category);
END;

DROP PROCEDURE IF EXISTS sp_UpdateItem;
CREATE PROCEDURE sp_UpdateItem(IN idItem bigint, IN name varchar(255), IN producer varchar(255), IN description varchar(255), IN price float, IN category bigint)
BEGIN
    update items set name=name, producer=producer, description=description, price=price, category=category where id=idItem;
END;

DROP PROCEDURE IF EXISTS sp_DeleteItem;
CREATE PROCEDURE sp_DeleteItem(IN idItem bigint)
BEGIN
    delete from items where id=idItem;
END;

# Store Procedures for the users table

DROP PROCEDURE IF EXISTS sp_GetUsers;
CREATE PROCEDURE sp_GetUsers()
BEGIN
    select * from users;
END;

DROP PROCEDURE IF EXISTS sp_GetUserById;
CREATE PROCEDURE sp_GetUserById(IN idUser bigint)
BEGIN
    select * from users where id=idUser;
END;

DROP PROCEDURE IF EXISTS sp_GetUserByUsername;
CREATE PROCEDURE sp_GetUserByUsername(IN _username varchar(255))
BEGIN
    select * from users where username=_username;
END;

DROP PROCEDURE IF EXISTS sp_CreateUser;
CREATE PROCEDURE sp_CreateUser(IN firstName varchar(255), IN lastName varchar(255), IN email varchar(255), IN username varchar(255), IN password varchar(255))
BEGIN
    insert into users (firstname, lastname, email, password, username, role) VALUES (firstName, lastName, email, password, username, 'user');
END;

DROP PROCEDURE IF EXISTS sp_UpdateUser;
CREATE PROCEDURE sp_UpdateUser(IN idUser bigint, IN firstName varchar(255), IN lastName varchar(255), IN email varchar(255), IN password varchar(255), IN username varchar(255), IN role varchar(255))
BEGIN
    update users set firstname=firstName, lastname=lastName, email=email, password=password, username=username, role=role where id=idUser;
END;

DROP PROCEDURE IF EXISTS sp_DeleteUser;
CREATE PROCEDURE sp_DeleteUser(IN idUser bigint)
BEGIN
    delete from users where id=idUser;
END;

# Procedures for the ordes table

DROP PROCEDURE IF EXISTS sp_GetOrders;
CREATE PROCEDURE sp_GetOrders()
BEGIN
    select * from orders;
END;

DROP PROCEDURE IF EXISTS sp_GetOrdersByUserID;
CREATE PROCEDURE sp_GetOrdersByUserID(IN idUser bigint)
BEGIN
    select * from orders where user_id=idUser;
END;

DROP PROCEDURE IF EXISTS sp_GetOrder;
CREATE PROCEDURE sp_GetOrder(IN idOrder bigint)
BEGIN
    select * from orders where id=idOrder;
END;

DROP PROCEDURE IF EXISTS sp_CreateOrder;
CREATE PROCEDURE sp_CreateOrder(IN userID bigint, IN paymentMethod varchar(255), paymentMethodID varchar(255), IN total float, IN status varchar(255) )
BEGIN
    insert into orders (user_id, payment_method, payment_id, total_price, status) VALUES  (userID, paymentMethod, paymentMethodID, total, status);
END;

DROP PROCEDURE IF EXISTS sp_UpdateOrder;
CREATE PROCEDURE sp_UpdateOrder(IN idOrder bigint, IN userID bigint, IN paymentMethod varchar(255), IN total float, IN status varchar(255), paymentMethodID varchar(255))
BEGIN
    update orders set user_id=userID, payment_method=paymentMethod, payment_id=paymentMethodID, total_price=total, status=status where id=idOrder;
END;

DROP PROCEDURE IF EXISTS sp_GetItemsByOrderID;
CREATE PROCEDURE sp_GetItemsByOrderID(IN idOrder bigint)
BEGIN
    select * from items where id in (select item_id from orders_items where order_id=idOrder);
END;

DROP PROCEDURE IF EXISTS sp_CreateOrderItem;
CREATE PROCEDURE sp_CreateOrderItem(IN idOrder bigint, IN idItem bigint, IN price float)
BEGIN
    insert into orders_items (order_id, item_id, price) VALUES (idOrder, idItem, price);
END;


DROP PROCEDURE IF EXISTS sp_GetTotalNumberOfItems;
CREATE PROCEDURE sp_GetTotalNumberOfItems()
BEGIN
    select count(*) from items;
END;

DROP PROCEDURE IF EXISTS sp_GetTotalNumberOfUsers;
CREATE PROCEDURE sp_GetTotalNumberOfUsers()
BEGIN
    select count(*) from users;
END;

DROP PROCEDURE IF EXISTS sp_GetTotalNumberOfOrders;
CREATE PROCEDURE sp_GetTotalNumberOfOrders()
BEGIN
    select count(*) from orders;
END;

# Procedures to retrieve some statistics about the orders, items and users

DROP PROCEDURE IF EXISTS sp_GetStatLastMonth;
CREATE PROCEDURE sp_GetStatLastMonth()
BEGIN
    select (select sum(total_price) from orders where created_at > DATE_SUB(NOW(), INTERVAL 1 MONTH)) as total_amount,
        (select count(*) from orders where created_at > DATE_SUB(NOW(), INTERVAL 1 MONTH)) as total_orders,
        (select count(*) from users where created_at > DATE_SUB(NOW(), INTERVAL 1 MONTH)) as total_users;
END;

DROP PROCEDURE IF EXISTS sp_GetStatLastWeek;
CREATE PROCEDURE sp_GetStatLastWeek()
BEGIN
    select (select sum(total_price) from orders where created_at > DATE_SUB(NOW(), INTERVAL 1 WEEK)) as total_amount,
        (select count(*) from orders where created_at > DATE_SUB(NOW(), INTERVAL 1 WEEK)) as total_orders,
        (select count(*) from users where created_at > DATE_SUB(NOW(), INTERVAL 1 WEEK)) as total_users;
END;

DROP PROCEDURE IF EXISTS sp_GetStatLastDay;
CREATE PROCEDURE sp_GetStatLastDay()
BEGIN
    select (select sum(total_price) from orders where created_at > DATE_SUB(NOW(), INTERVAL 1 DAY)) as total_amount,
        (select count(*) from orders where created_at > DATE_SUB(NOW(), INTERVAL 1 DAY)) as total_orders,
        (select count(*) from users where created_at > DATE_SUB(NOW(), INTERVAL 1 DAY)) as total_users;
END;

# Get users who have spent more
DROP PROCEDURE IF EXISTS sp_GetUsersWhoSpentMore;
CREATE PROCEDURE sp_GetUsersWhoSpentMore()
BEGIN
    select u.id, u.username, u.email, u.created_at, sum(oi.price) as total_spent from users u, orders o, orders_items oi where u.id=o.user_id and o.id=oi.order_id group by u.id order by total_spent desc limit 10;
END;

# Get Most Ordered Items
DROP PROCEDURE IF EXISTS sp_GetMostOrderedItems;
CREATE PROCEDURE sp_GetMostOrderedItems()
BEGIN
    select i.id, i.name, i.description, i.price, count(*) as total_orders from items i, orders_items oi where i.id=oi.item_id group by i.id order by total_orders desc limit 10;
END;

# Get least ordered items
DROP PROCEDURE IF EXISTS sp_GetLeastOrderedItems;
CREATE PROCEDURE sp_GetLeastOrderedItems()
BEGIN
    select i.id, i.name, i.description, i.price, count(*) as total_orders from items i, orders_items oi where i.id=oi.item_id group by i.id order by total_orders
    limit 10;
END;

# Get items that have not been ordered
DROP PROCEDURE IF EXISTS sp_GetItemsNotOrdered;
CREATE PROCEDURE sp_GetItemsNotOrdered()
BEGIN
    select i.id, i.name, i.description, i.price from items i where i.id not in (select oi.item_id from orders_items oi);
END;
