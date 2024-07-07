CREATE TABLE reservations (
                              id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              user_id uuid,
                              restaurant_id UUID REFERENCES restaurants(id) ON DELETE CASCADE,
                              reservation_time TIMESTAMP NOT NULL,
                              status VARCHAR(20) CHECK (status IN ('pending', 'confirmed', 'cancelled')) NOT NULL,
                              created_at timestamp default current_timestamp,
                              updated_at timestamp default current_timestamp,
                              deleted_at timestamp
);
