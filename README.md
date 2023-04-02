# 1. Omborxonadagi (stock) productlarni boshqa bir magazinga otkizishlik kerak.
Masalan: Mega planeta magazinidan 10 ta product Texno mart magaziniga otkizdi


Hisobot 
### route http://localhost:4001/report/exchange ✅

# 2. Har bir hodim qancha mahsulot sotganligi boyicha malumot chiqishi kerak
Masalan:
Сотурник         | Категори     | Продукт    | Количество | Обший Цена   | Дата
---------------------------------------------------------------------------------
Eshmat Toshmatov  Oyoq-Kiyim	 Poyavzal 		 20 		300_000        2022-12-20
Ahmad Ahmadov	  Texnika		 Asus 5 		 43 		210_000        2022-12-30
Zafar Zafarov	  Telefon		 Iphone 		 65 		330_000        2022-12-20
Erkin Erkinboyev  Avto			 Kalit	 		 90 		3_000_000      2022-12-20

Hisoblash

### route http://localhost:4001/report/employee ✅

# 3. Promo code CRUD Method boladi. (Update shartmas)
Promo Code
	- name
	- discount = 47 500
	- discount_type => Фикс | Процент
	- order_limit_price => 95 000

### route http://localhost:4001/promo_code ✅

# 4. Order Total Sum Api boladi. Shu api order_id berilsa umumiy summa hisoblab
	berishi kerak. Agar promo code ham berilsa chegirmalar ham hisoblanishi kerak

Masalan:
	order_id : 1
	promo_code :"JUBAJUBA"

### route http://localhost:4001/total_order_price/:id  ✅

# 5. Order Item qoshilganda produclarni Stock (Склад) dan olishi kerak.
	Agar magazin sklad da product qolmagan bolsa "Товарь не найден" habari chiqishi kerak

report.go 

### route http://localhost:4001/order_item/ ✅

Bonus Vazifa -> Ифторлик
## 6. Stock (Склад) bor malumotlarni Excel formatida chiqishi kerak. xlsx yoki csv
Maslan shu formatida:

Намеклатура   |  Цена	         | Megan planeta  	|  Texnomart  | Ustore
--------------------------------------------------------------------------
Sport         |                  |      2           |    10       |   0
Krosovka      | 	100 000      |      2			|    10		  |   0
Texnika       |                  |      40          |    37       |   64
Iphone		  |		200 000      |      40			|    32		  |   55
HP computer   |     400 000      |      0			|  	 5		  |   9

Категори - серий ранг чикиши керак

### route http://localhost:4001/report/stock  ✅

Deadline: 2023-04-02 20:00
