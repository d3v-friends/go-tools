~~~mysql
CREATE VIEW coupons_view AS
(
SELECT C.id,
       C.account_id,
       C.point,
       C.currency,
       U.point             as left_point,
       (U.point * C.price) AS left_balance,
       C.created_at
FROM coupons C,
     coupon_balances U
WHERE C.coupon_balance_id = U.id
ORDER BY C.created_at DESC
    );

~~~
