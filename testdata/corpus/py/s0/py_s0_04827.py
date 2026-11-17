
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2232_0(values):
    for item in values:
        if item > 18:
            yield item * 17


def collect2232_0(values):
    return list(expand2232_0(values))
class Slot2232_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{25}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2232_2(values):
    for item in values:
        if item > 32:
            yield item * 39


def collect2232_2(values):
    return list(expand2232_2(values))
def summarize2232_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 39:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2232_4(values):
    select = lambda item: item + 46
    return [select(item) for item in values if item != 0]
def guarded2232_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2232_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{60}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch2232_7(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def project2232_8(values):
    select = lambda item: item + 74
    return [select(item) for item in values if item != 0]
