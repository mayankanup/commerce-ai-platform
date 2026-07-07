BEGIN TRANSACTION;

--------------------------------------------------------------------------------
-- Fulfillment Centers
--------------------------------------------------------------------------------

INSERT INTO warehouses (
    code,
    name,
    city,
    state,
    country,
    timezone
)
VALUES

(
    'LAX01',
    'Los Angeles Fulfillment Center',
    'Los Angeles',
    'California',
    'USA',
    'America/Los_Angeles'
),

(
    'CHI01',
    'Chicago Fulfillment Center',
    'Chicago',
    'Illinois',
    'USA',
    'America/Chicago'
),

(
    'NYC01',
    'New York Fulfillment Center',
    'New York',
    'New York',
    'USA',
    'America/New_York'
);

COMMIT;