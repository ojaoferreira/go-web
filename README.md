```
docker run \
        --name postgres \
        --publish 5432:5432 \
        --env POSTGRES_PASSWORD=mysecretpassword \
        --env POSTGRES_DB=alura_loja \
        -d postgres
```

```
create table produtos (
	id serial not null primary key,
	nome varchar not null,
	descricao varchar,
	preco decimal not null,
	quantidade integer not null
);

insert into produtos (nome, descricao, preco, quantidade) values ('Camisa', 'Amarelo', 19.99, 5);

insert into produtos (nome, descricao, preco, quantidade) values ('Short', 'Preto', 8.82, 3);

select * from produtos;
```