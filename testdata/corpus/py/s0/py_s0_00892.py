
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize407_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 58:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot407_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize407_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded407_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
