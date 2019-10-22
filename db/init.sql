create table client (
    client_id serial primary key,
    firstname varchar,
    lastname varchar,
    street varchar,
    city varchar,
    statec varchar,
    zip varchar,
    email varchar,
    username varchar unique,
    pass varchar,
    montlyincomes decimal,
    monthlyexpenses decimal,
    approve varchar,
    joint integer,
    requestjoint integer
    
);

create table checkingaccoclient (
    client_id integer REFERENCES client(client_id),
    checkingaccount serial primary key
);

create table accounts (
    idacc serial primary key,
    client_id integer REFERENCES client(client_id),
    checkingaccount integer REFERENCES checkingaccoclient,
    amount DECIMAL
);

create table joint(
    idjoint serial primary key,
    fisrtname1 varchar,
    lastname1 varchar,
    firstname2 varchar,
    lastname2 varchar
);




