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