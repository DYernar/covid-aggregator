import flask
import json
from flask import request, jsonify
from packages import AkmolaParse

data = AkmolaParse()

app = flask.Flask(__name__)
app.config["DEBUG"] = True


@app.route('/', methods=['GET'])
def home():
    return "<h1>CoronoVirus parsed data access:</h1><ul><li><a href="+"http://ffe84a27.ngrok.io/akmola"+">Akmola city</a><br></li></ul>"


@app.route('/akmola', methods=['GET'])
def akmola():
    return jsonify(data)

if  __name__=="__main__":
    app.run(debug=True, port=8000)