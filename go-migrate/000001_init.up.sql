CREATE TABLE public."DataValue"
(
    "Валюта1" character varying COLLATE pg_catalog."default",
    "Валюта2" character varying COLLATE pg_catalog."default",
    "курс" real,
    "последнее время обновления" date DEFAULT '2010-06-30'::date
)

TABLESPACE pg_default;

ALTER TABLE public."DataValue"
    OWNER to postgres;