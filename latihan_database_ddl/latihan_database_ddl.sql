--membuat table users
CREATE TABLE users (
user_id INT PRIMARY KEY,
username VARCHAR UNIQUE NOT NULL,
email VARCHAR UNIQUE NOT NULL,
nama_lengkap VARCHAR NOT NULL
);

-- Membuat tabel orders
CREATE TABLE orders (
order_id INT PRIMARY KEY,
user_id INT NOT NULL,
tanggal_pemesanan TIMESTAMP DEFAULT NOW(),
status VARCHAR NOT NULL,
CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (user_id)
);

-- Membuat tabel order_items
CREATE TABLE order_items (
item_id INT PRIMARY KEY,
order_id INT NOT NULL,
product_name VARCHAR NOT NULL,
quantity INT NOT NULL,
harga_per_item NUMERIC NOT NULL,
CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (order_id)
);