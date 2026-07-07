BEGIN TRANSACTION;

INSERT INTO returns (
    order_id,
    status,
    reason,
    requested_at,
    approved_at
)
VALUES

(2,'COMPLETED','Size was too small','2026-01-15','2026-01-16'),

(5,'APPROVED','Changed my mind','2026-01-18','2026-01-19'),

(7,'REQUESTED','Received wrong color','2026-01-21',NULL),

(10,'REJECTED','Outside return policy','2026-01-25','2026-01-26'),

(14,'COMPLETED','Item arrived damaged','2026-02-01','2026-02-02');

COMMIT;