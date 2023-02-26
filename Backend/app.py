from flask import Flask, render_template, request
import sys, mariadb, json

app = Flask(__name__)
try:
    conn = mariadb.connect(
        user="budgetuser",
        password="budget-user-db-pass0!",
        host="127.0.0.1",
        port=3306,
        database="budget"

    )
except mariadb.Error as e:
    print(f"Error connecting to MariaDB Platform: {e}")
    sys.exit(1)

# Get Cursor
cur = conn.cursor()

@app.route('/expense', methods=["POST", "GET"])
def AddExpense():
#    error = None
	if request.method == 'POST':
		data = request.get_json()
		#print(data)
		print("Household: " + data['household'] + ", Food: " + data['food'] + ", Transport: " + data['transport'] + ", Misc: " + data['misc'])
		try:
			cur.execute("INSERT INTO expenses (household_expenses, food_expenses, transport_expenses, misc_expenses) VALUES (?, ?, ?, ?)", (data['household'], data['food'], data['transport'], data['misc']))
			conn.commit()
		except mariadb.Error as e:
			print(f"Error: {e}")
		return '', 200
	else:
		try:
			cur.execute("SELECT id, household_expenses, food_expenses, transport_expenses, misc_expenses, created_at FROM expenses ORDER BY id DESC LIMIT 1")
			res = cur.fetchone()
			total = res[1] + res[2] + res[3] + res[4]
			response = {"expenses": {"household": res[1], "food": res[2], "transport": res[3], "misc": res[4] }, "timestamp": res[5], "id": res[0], "total": total }
		except mariadb.Error as e:
			print(f"Error: {e}")
		return response
#budgetuser@localhost:budget-user-db-pass0!