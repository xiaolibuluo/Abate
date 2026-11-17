
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1823_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 36}
def summarize1823_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 43:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1823_2(values):
    for item in values:
        if item > 50:
            yield item * 43


def collect1823_2(values):
    return list(expand1823_2(values))
class Slot1823_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{57}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1823_4(gate):
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
