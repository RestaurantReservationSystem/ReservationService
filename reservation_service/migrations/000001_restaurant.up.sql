CREATE TABLE restaurants (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             name VARCHAR(100) NOT NULL,
                             address TEXT NOT NULL,
                             phone_number VARCHAR(15),
                             description TEXT,
                             created_at timestamp default current_timestamp,
                             updated_at timestamp default current_timestamp,
                             deleted_at timestamp
);
