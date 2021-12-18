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