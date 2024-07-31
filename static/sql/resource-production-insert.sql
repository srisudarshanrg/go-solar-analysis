-- INSERT INTO public.resource_production(
-- 	id, country, code, year, gas_production, coal_production, oil_production, created_at, updated_at)
-- 	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

COPY resource_production(country, code, year, gas_production, coal_production, oil_production)
FROM 'D:\School\Grade 10\Coding_Projects\idp-project\static\databases\resource_production.csv'
DELIMITER ','
CSV HEADER;