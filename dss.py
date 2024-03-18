from conflict_manager import ConflictManager
from scd import Scd


class Dss:
    def __init__(self) -> None:
        self.conflict_manager = ConflictManager()
        self.scd = Scd()


        