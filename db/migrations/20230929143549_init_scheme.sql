-- +goose Up
-- +goose StatementBegin
CREATE TABLE airline (
        code VARCHAR(2) PRIMARY KEY,
        name TEXT
);

CREATE TABLE provider (
        id VARCHAR(2) PRIMARY KEY,
        name TEXT
);

CREATE TABLE schema (
        id SERIAL PRIMARY KEY,
        name TEXT
);

CREATE TABLE account (
        id SERIAL PRIMARY KEY,
        schema_id INT REFERENCES schema(id)
);

CREATE TABLE provider_airline (
        id SERIAL PRIMARY KEY,
        provider_id VARCHAR(2) REFERENCES provider(id),
        airline_code VARCHAR(2) REFERENCES airline(code),

        UNIQUE(provider_id, airline_code)
);

CREATE TABLE schema_provider (
        id SERIAL PRIMARY KEY,
        schema_id INT REFERENCES schema(id),
        provider_id VARCHAR(2) REFERENCES provider(id)
);

CREATE TABLE account_airline (
        id SERIAL PRIMARY KEY,
        account_id INT REFERENCES account(id),
        airline_code VARCHAR(2) REFERENCES airline(code)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE account_airline;

DROP TABLE schema_provider;

DROP TABLE provider_airline;

DROP TABLE account;

DROP TABLE schema;

DROP TABLE provider;

DROP TABLE airline;
-- +goose StatementEnd
