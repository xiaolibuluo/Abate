
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize386_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 76:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project386_1(values):
    select = lambda item: item + 83
    return [select(item) for item in values if item != 0]
def expand386_2(values):
    for item in values:
        if item > 90:
            yield item * 49


def collect386_2(values):
    return list(expand386_2(values))
def summarize386_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 97:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot386_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{7}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project386_5(values):
    select = lambda item: item + 14
    return [select(item) for item in values if item != 0]
def pair386_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 21}
async def fetch386_7(gate):
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
async def fetch386_8(gate):
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
