INSERT INTO users (user_id, username, email, nama_lengkap)
VALUES
  (1,'suci', 'suci@gmail.com', 'Suci Atsila'),
  (2, 'ibe', 'ibe@gmail.com', 'Iqbal Zharfan');
  
INSERT INTO orders (user_id, status)
VALUES
  (1, 'dalam pengiriman'),
  (2, 'selesai'),
  (1, 'dibatalkan');
  
  -- Pesanan 1
INSERT INTO order_items (order_id, product_name, quantity, harga_per_item)
VALUES
  (1, 'Buku', 4, 5000),
  (1, 'Pensil', 3, 2000);

-- Pesanan 2
INSERT INTO order_items (order_id, product_name, quantity, harga_per_item)
VALUES
  (2, 'Pena', 5, 2500),
  (2, 'PEnghapus', 2, 2000);

-- Pesanan 3
INSERT INTO order_items (order_id, product_name, quantity, harga_per_item)
VALUES
  (3, 'Penggaris', 1, 6000);
  
SELECT orders.order_id, users.username
FROM orders
JOIN users ON orders.user_id = users.user_id;

SELECT
  orders.order_id,
  users.username,
  SUM(order_items.quantity * order_items.harga_per_item) AS total_harga
FROM orders
JOIN users ON orders.user_id = users.user_id
JOIN order_items ON orders.order_id = order_items.order_id
GROUP BY orders.order_id, users.username;