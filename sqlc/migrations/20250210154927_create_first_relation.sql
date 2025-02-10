-- +goose Up
ALTER TABLE users ADD FOREIGN KEY (position_id) REFERENCES positions (id);
ALTER TABLE users ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE users ADD FOREIGN KEY (role_id) REFERENCES roles (id);
ALTER TABLE approval_flows ADD FOREIGN KEY (approval_id) REFERENCES users (id);
ALTER TABLE approval_flows ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE approval_flows ADD FOREIGN KEY (flows_name_id) REFERENCES approval_flow_names (id);
ALTER TABLE common_requests ADD FOREIGN KEY (approval_flows_id) REFERENCES approval_flows (id);
ALTER TABLE common_requests ADD FOREIGN KEY (user_request_id) REFERENCES users (id);
ALTER TABLE common_requests ADD FOREIGN KEY (current_approval_id) REFERENCES users (id);
ALTER TABLE common_requests ADD FOREIGN KEY (settle_by) REFERENCES users (id);
ALTER TABLE common_requests ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE permissions ADD FOREIGN KEY (role_id) REFERENCES roles (id);

-- +goose Down
ALTER TABLE users DROP CONSTRAINT users_position_id_fkey;
ALTER TABLE users DROP CONSTRAINT users_department_id_fkey;
ALTER TABLE users DROP CONSTRAINT users_role_id_fkey;
ALTER TABLE approval_flows DROP CONSTRAINT approval_flows_approval_id_fkey;
ALTER TABLE approval_flows DROP CONSTRAINT approval_flows_department_id_fkey;
ALTER TABLE approval_flows DROP CONSTRAINT approval_flows_flows_name_id_fkey;
ALTER TABLE common_requests DROP CONSTRAINT common_requests_approval_flows_id_fkey;
ALTER TABLE common_requests DROP CONSTRAINT common_requests_user_request_id_fkey;
ALTER TABLE common_requests DROP CONSTRAINT common_requests_current_approval_id_fkey;
ALTER TABLE common_requests DROP CONSTRAINT common_requests_settle_by_fkey;
ALTER TABLE common_requests DROP CONSTRAINT common_requests_department_id_fkey;
ALTER TABLE permissions DROP CONSTRAINT permissions_role_id_fkey;
