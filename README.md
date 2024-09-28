# gowebgin
demo


数据库创建
gomarket
表格
CREATE TABLE product (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Description TEXT,
    Price DECIMAL(10, 2) NOT NULL,
    Stock INT NOT NULL,
    Category VARCHAR(100) NOT NULL
);


CREATE TABLE user (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Username VARCHAR(100) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL
);



数据插入
INSERT INTO Product (name, description, price, stock, category) VALUES
('商品1', '这是商品11的描述', 10.99, 100, '类别A'),
('商品2', '这是商品12的描述', 20.50, 200, '类别B'),
('商品3', '这是商品13的描述', 15.75, 150, '类别A'),
('商品4', '这是商品14的描述', 30.00, 80, '类别C'),
('商品5', '这是商品15的描述', 5.25, 300, '类别B'),
...







