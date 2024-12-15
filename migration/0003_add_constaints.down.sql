begin;

alter table "user"
    drop constraint unique_email;

commit;