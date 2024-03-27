from flask import Flask, request

from dss import Dss

import json

app = Flask(__name__)

database = {}

dss = Dss()


@app.route('/uss/v1/operational_intents/<operational_intent_id>', methods=['GET'])
def get_oir(operational_intent_id):
    print(database)
    if operational_intent_id not in database:
        return {"msg": 'No such OIR'}, 404

    return json.dumps(database[operational_intent_id], default=lambda o: o.__dict__)


@app.route('/injection/oir', methods=['PUT'])
def inject_oir():
    volume = request.get_json()
    try:
        dss.conflict_manager.check_restrictions(volume)
        dss.scd.check_strategic_conflicts(volume)
        oir = {}
        oir["operational_intent"] = {}
        oir["operational_intent"]["reference"] = dss.scd.put_operational_intent(volume)
        oir["operational_intent"]["details"] = {
            "volumes": [],
            "off_nominal_volumes": [],
            "priority": 0
        }
        oir["operational_intent"]["details"]["volumes"].append(volume)
        database[oir["operational_intent"]["reference"]['id']] = oir


    except Exception as ex:
        print(f"Erro na criação: {ex}")
        return {"msg": "Erro"}, 400

    return {"Success": True}, 201


if __name__ == '__main__':
    app.run(port=5050, host="0.0.0.0")
