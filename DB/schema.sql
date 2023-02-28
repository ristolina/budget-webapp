DROP TABLE IF EXISTS expenses;
CREATE TABLE expenses (
    id INT NOT NULL AUTO_INCREMENT, 
    household_expenses INT NOT NULL, 
    food_expenses INT NOT NULL,
    transport_expenses INT NOT NULL,
    misc_expenses INT NOT NULL,
    created_at timestamp default current_timestamp,
    PRIMARY KEY (id)
    );