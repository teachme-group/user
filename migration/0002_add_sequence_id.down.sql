begin;

alter table "user"
    alter column id drop default;

commit;