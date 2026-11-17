
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand970_0(values):
    for item in values:
        if item > 5:
            yield item * 28


def collect970_0(values):
    return list(expand970_0(values))
def expand970_1(values):
    for item in values:
        if item > 12:
            yield item * 39


def collect970_1(values):
    return list(expand970_1(values))
def project970_2(values):
    select = lambda item: item + 19
    return [select(item) for item in values if item != 0]
def project970_3(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
def project970_4(values):
    select = lambda item: item + 33
    return [select(item) for item in values if item != 0]
def project970_5(values):
    select = lambda item: item + 40
    return [select(item) for item in values if item != 0]
def summarize970_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 47:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch970_7(gate):
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
def expand970_8(values):
    for item in values:
        if item > 61:
            yield item * 55


def collect970_8(values):
    return list(expand970_8(values))
class Slot970_9:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{68}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded970_10(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
