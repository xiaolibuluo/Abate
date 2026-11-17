
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize371_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 75:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch371_1(gate):
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
class Slot371_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{89}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch371_3(gate):
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
def expand371_4(values):
    for item in values:
        if item > 6:
            yield item * 27


def collect371_4(values):
    return list(expand371_4(values))
def pair371_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
