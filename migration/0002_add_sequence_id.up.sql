alter table "user"
    alter column id set default gen_random_uuid();