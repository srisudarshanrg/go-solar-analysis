-- INSERT INTO public.resource_consumption(
-- 	id, country, oil, electricity, coal, natural_gas, biofuel, created_at, updated_at)
-- 	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

COPY resource_consumption(country, oil, electricity, coal, natural_gas, biofuel)
FROM 'D:\School\Grade 10\Coding_Projects\idp-project\static\databases\resource_consumption.csv'
DELIMITER ','
CSV HEADER;