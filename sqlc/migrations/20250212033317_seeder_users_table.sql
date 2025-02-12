-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, full_name, username,phone, email, password, role_id, department_id, position_id) VALUES
    ('b93ad1ab-3b15-4bbc-b4d0-de000c68d7e5',
     'Mizumaru',
     'mizumaru1108',
     '08612351623',
     'maru@gmail.com',
     '$2a$10$rqjMpYBlQZGP0LtJsyfyqeowrFnfSKA0o/.X03hnbO5VCGK3ZJxwC',
     'fc3d1018-5abf-4f70-92cd-409c199d021e',
     'b3eb5159-3a60-468a-9280-93b1065b203e',
     'e1c0cff9-67db-41cb-86fa-564be844e936'
     );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE username IN ('mizumaru1108');
-- +goose StatementEnd
