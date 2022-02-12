CREATE TABLE IF NOT EXISTS stubook_list (
    "stubook_list_id" bigserial primary key,
    "status" varchar(100),
    "given_date" DATE NOT NULL DEFAULT CURRENT_DATE,
    "deadline" DATE,
    "bk_name" varchar(150),
    "bk_authorname" varchar(150),
    "bk_id" varchar(50),
    "bk_numsbook" int,
    "student_id" varchar(100),
    constraint fk_student foreign key("student_id") references student ("student_id")
);