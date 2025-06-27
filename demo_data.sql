-- TENANTS
INSERT INTO
    tenants (id, name)
VALUES
    (
        '4d223370-910f-49c8-be03-2d6bd46f86e5',
        'Tenant 1'
    );

-- USERS
INSERT INTO
    users (id, email, username, password, tenant_id)
VALUES
    (
        '489602d2-b96e-4dfa-920d-e1949c332f89',
        '1234@1234.1234',
        '1234',
        '$2a$12$JU8K/8GfUzb0rMZflUr9quLHgNsRA1VE7wyzM66cO6LYCrBkaaUr2',
        '4d223370-910f-49c8-be03-2d6bd46f86e5'
    );

-- GROUPS
INSERT INTO
    groups (id, name, tenant_id)
VALUES
    (
        '3818a0cb-0ae1-4010-a249-f5bdf6e949b1',
        'Group 1',
        '4d223370-910f-49c8-be03-2d6bd46f86e5'
    );

-- ROLES
INSERT INTO
    roles (id, name)
VALUES
    ('92b3abbf-9f71-425b-9ed6-3cf6fc4b922a', 'Role 1');

-- USER_GROUPS
INSERT INTO
    user_groups (user_id, group_id)
VALUES
    ('489602d2-b96e-4dfa-920d-e1949c332f89', '3818a0cb-0ae1-4010-a249-f5bdf6e949b1');

-- GROUP_ROLES
INSERT INTO
    group_roles (group_id, role_id)
VALUES
    ('3818a0cb-0ae1-4010-a249-f5bdf6e949b1', '92b3abbf-9f71-425b-9ed6-3cf6fc4b922a');

-- PERMISSIONS
INSERT INTO
    permissions (id, name, description, resource, scope)
VALUES
    (
        'a18e1987-4a51-465b-af22-c25e590e6564',
        'Permission 1',
        'Test permission',
        'resource1',
        'read'
    );

-- ROLE_PERMISSIONS
INSERT INTO
    role_permissions (role_id, permission_id)
VALUES
    ('92b3abbf-9f71-425b-9ed6-3cf6fc4b922a', 'a18e1987-4a51-465b-af22-c25e590e6564');