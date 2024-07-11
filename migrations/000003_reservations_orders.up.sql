CREATE TABLE reservation_orders (
                                    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                    reservation_id UUID REFERENCES reservations(id) ON DELETE CASCADE,
                                    menu_item_id UUID REFERENCES menu(id) ON DELETE CASCADE,
                                    quantity INTEGER NOT NULL,
                                    created_at timestamp default current_timestamp,
                                    updated_at timestamp default current_timestamp,
                                    deleted_at timestamp
);
