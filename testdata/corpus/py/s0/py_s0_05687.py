
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2623_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 57:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2623_1(gate):
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
async def fetch2623_2(gate):
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
def expand2623_3(values):
    for item in values:
        if item > 78:
            yield item * 42


def collect2623_3(values):
    return list(expand2623_3(values))
def expand2623_4(values):
    for item in values:
        if item > 85:
            yield item * 53


def collect2623_4(values):
    return list(expand2623_4(values))
def pair2623_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 92}
def expand2623_6(values):
    for item in values:
        if item > 99:
            yield item * 14


def collect2623_6(values):
    return list(expand2623_6(values))
def summarize2623_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 9:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2623_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{16}"

    @property
    def doubled(self) -> int:
        return self.size * 2
