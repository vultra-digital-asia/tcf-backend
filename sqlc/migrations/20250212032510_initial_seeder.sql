-- +goose Up
-- +goose StatementBegin
INSERT INTO roles (id, name) VALUES
('5807117f-7bf0-4ed8-bdba-644a8b03d48e', 'admin'),
('d21fa4ec-8241-417f-8d8c-f091cdbe49c8', 'manager'),
('fc3d1018-5abf-4f70-92cd-409c199d021e', 'user');

INSERT INTO positions (id, name, hierarchy_level) VALUES
('e1c0cff9-67db-41cb-86fa-564be844e936', 'officer', 1);

INSERT INTO departments (id, name) VALUES
('b3eb5159-3a60-468a-9280-93b1065b203e', 'IT');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM positions WHERE name IN ('officer');
DELETE FROM departments WHERE name IN ('IT');
DELETE FROM roles WHERE name IN ('admin', 'manager', 'user');
-- +goose StatementEnd
