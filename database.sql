CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    phone VARCHAR(20),
    role VARCHAR(20) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE equipment_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE equipment (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES equipment_categories(id),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    daily_rate DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) DEFAULT 'available',
    purchase_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rentals (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    equipment_id INTEGER REFERENCES equipment(id),
    rental_date DATE NOT NULL,
    return_date DATE NOT NULL,
    total_cost DECIMAL(10, 2),
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

    CREATE TABLE IF NOT EXISTS payments (
      id SERIAL PRIMARY KEY,
      rental_id INTEGER REFERENCES rentals(id),
      amount DECIMAL(10, 2) NOT NULL,
      method VARCHAR(50) NOT NULL,
      status VARCHAR(20) DEFAULT 'pending',
      paid_at TIMESTAMP,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS maintenance (
      id SERIAL PRIMARY KEY,
      equipment_id INTEGER REFERENCES equipment(id),
      description TEXT,
      status VARCHAR(20) DEFAULT 'scheduled',
      start_date DATE,
      end_date DATE,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS notifications (
      id SERIAL PRIMARY KEY,
      user_id INTEGER REFERENCES users(id),
      channel VARCHAR(20) NOT NULL,
      recipient VARCHAR(100) NOT NULL,
      subject VARCHAR(255),
      message TEXT NOT NULL,
      status VARCHAR(20) DEFAULT 'sent',
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );



