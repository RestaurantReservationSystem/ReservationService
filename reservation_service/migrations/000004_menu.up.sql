CREATE TABLE menu (
                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                      restaurant_id UUID REFERENCES restaurants(id) ON DELETE CASCADE,
                      name VARCHAR(100) NOT NULL,
                      description TEXT,
                      price DECIMAL(10, 2) NOT NULL,
                      created_at timestamp default current_timestamp,
                      updated_at timestamp default current_timestamp,
                      deleted_at timestamp

);
