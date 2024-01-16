# balance 쉽게 보는 view

~~~mysql
CREATE VIEW balances AS
(
SELECT C.id,
       C.account_id,
       C.total_point,
       C.currency,
       U.point             as left_point,
       (U.point * C.price) AS left_balance,
       C.created_at
FROM coupons C,
     usages U
WHERE C.usage_id = U.id
ORDER BY C.created_at DESC
    );

DROP VIEW balances;
~~~
