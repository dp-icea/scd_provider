import requests
import uuid

from env import AUTH_URL, USS_BASE_URL, DSS_HOST


class Scd:

    def __init__(self) -> None:
        self.auth()

    def auth(self):
        url = "{0}/token?grant_type=client_credentials&intended_audience=localhost&issuer=localhost&scope={1}"
        self.strategic_coordination = requests.get(
            url.format(AUTH_URL, "utm.strategic_coordination")
        ).json()["access_token"]

    def check_strategic_conflicts(self, volume):
        url = DSS_HOST + "/dss/v1/operational_intent_references/query"
        body = {"area_of_interest": volume}
        header = {"authorization": f"Bearer {self.strategic_coordination}"}
        response = requests.post(url, headers=header, json=body).json()

        if len(response["operational_intent_references"]) > 0:
            raise Exception(
                f"Interseção com outra Intenção {response['operational_intent_references'][0]['id']}"
            )
        else:
            print("Sem intenções para o volume")

    def put_operational_intent(self, volume):
        id = str(uuid.uuid4())
        url = DSS_HOST + f"/dss/v1/operational_intent_references/{id}"

        body = {
            "flight_type": "VLOS",
            "extents": [volume],
            "key": [],
            "state": "Accepted",
            "uss_base_url": USS_BASE_URL,
            "new_subscription": {
                "uss_base_url": USS_BASE_URL,
                "notify_for_constraint": False,
            },
        }

        print(body)

        header = {"authorization": f"Bearer {self.strategic_coordination}"}

        response = requests.put(url, headers=header, json=body).json()

        print(f"OIR criada com id: {id}")
        print(response)

        return response["operational_intent_reference"]
