-- task5
CREATE OR REPLACE FUNCTION get_product_from_store() RETURNS TRIGGER LANGUAGE PLPGSQL
    AS
$$
    DECLARE 
        storeId integer;
    BEGIN
        
        SELECT orders.store_id FROM orders INTO storeId WHERE orders.order_id = new.order_id;

        UPDATE stocks SET quantity = quantity - new.quantity WHERE store_id = storeId AND product_id =  new.product_id;

        return new;
    END;
$$;

CREATE TRIGGER order_item_product_tg
BEFORE INSERT  ON order_items
FOR EACH ROW EXECUTE PROCEDURE get_product_from_store();


CREATE OR REPLACE FUNCTION add_product_to_store() RETURNS TRIGGER LANGUAGE PLPGSQL
    AS
$$
    DECLARE 
        storeId integer;
    BEGIN
        
        SELECT orders.store_id FROM orders INTO storeId WHERE orders.order_id = old.order_id;

        UPDATE stocks SET quantity = quantity + old.quantity WHERE store_id = storeId AND product_id =  old.product_id;

        return new;
    END;
$$;

CREATE TRIGGER add_product_to_store_tg
BEFORE DELETE ON order_items
FOR EACH ROW EXECUTE PROCEDURE add_product_to_store();