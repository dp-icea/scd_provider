import requests

from env import AUTH_URL, DSS_HOST


class ConflictManager:

    dss_host = DSS_HOST

    def __init__(self) -> None:
        self.auth()

    def auth(self):
        url = AUTH_URL + "/token?grant_type=client_credentials&intended_audience=localhost&issuer=localhost&scope={0}"
        self.constraint_management = requests.get(url.format("utm.constraint_management")).json()["access_token"]

    def check_restrictions(self, volume):
        self.auth()
        url = self.dss_host + "/dss/v1/constraint_references/query"
        body = {
            "area_of_interest": volume
        }
        header = {"authorization": f"Bearer {self.constraint_management}"}
        response = requests.post(url, headers=header, json=body).json()
        print(response)

        if len(response['constraint_references']) > 0:
            raise Exception("Interseção com Constraint")
        else:
            print("Sem restrições para o volume")

