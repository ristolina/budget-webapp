from flask import Flask, render_template, redirect, url_for, request
import requests, json
from datetime import date

app = Flask(__name__)
app.config['SECRET_KEY'] = '997ca1b14c99c48c3cc32dd70d23a6ef45cbaa71aacefce3'
@app.route("/")
def start():
	try:
		requestUrl = "http://127.0.0.1:5001/expense"
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
		Data = dict(request.form)
		Headers = {"Content-Type":"application/json"}
		#	Data['yearmonth'] = str(Date.year)[2:] + str(Date.month).zfill(2)
		print(json.dumps(Data))
		requestUrl = "http://127.0.0.1:5001/expense"
		try:
			response = requests.post(requestUrl, headers=Headers, data=json.dumps(Data))
		except requests.exceptions.ConnectionError as e:
			print(e)
		print(response)
		return redirect(url_for('start'))