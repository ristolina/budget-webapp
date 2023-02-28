from flask import Flask, render_template, redirect, url_for, request
import requests, json, os, sys
from datetime import date

app = Flask(__name__)
app.config['SECRET_KEY'] = os.getenv('FLASK_SECRET')
@app.route("/")
def start():
	try:
		requestUrl = "http://backend:5001/expense"
		response = requests.get(requestUrl)
		data = json.loads(response.text)
		expenses = data['expenses']
		total = data['total']
	except requests.exceptions.ConnectionError as e:
		print(e)
		expenses = {}
		total = None
	return render_template("base.html", expense = expenses, total = total)

@app.route('/add-expense/', methods=["POST"])
def AddExpense():
	if request.method == 'POST':
		Date = date.today()
		Data = {"household": int(request.form["household"]), "food": int(request.form["food"]), "transport": int(request.form["household"]), "misc": int(request.form["misc"])}
		Headers = {"Content-Type":"application/json"}
		#	Data['yearmonth'] = str(Date.year)[2:] + str(Date.month).zfill(2)
		print(json.dumps(Data), file=sys.stderr)
		requestUrl = "http://backend:5001/expense"
		try:
			response = requests.post(requestUrl, headers=Headers, data=json.dumps(Data))
			print(response)
		except requests.exceptions.ConnectionError as e:
			print(e)
		print(response, file=sys.stderr)
		return redirect(url_for('start'))