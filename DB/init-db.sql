USE budget;
CREATE TABLE expenses (
    id INT NOT NULL AUTO_INCREMENT, 
    household_expenses INT NOT NULL, 
    food_expenses INT NOT NULL,
    transport_expenses INT NOT NULL,
    misc_expenses INT NOT NULL,
    created_at timestamp default current_timestamp,
    PRIMARY KEY (id)
    );

INSERT INTO expenses (household_expenses, food_expenses, transport_expenses, misc_expenses) 
    VALUES 
        (16000, 8000, 5000, 2000), 
        (12000, 10000, 4000, 1500);