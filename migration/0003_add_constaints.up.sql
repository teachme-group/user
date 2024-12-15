begin;

alter table "user"
    add constraint unique_email unique (email);

commit;