
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1522_0(values):
    for item in values:
        if item > 100:
            yield item * 49


def collect1522_0(values):
    return list(expand1522_0(values))
async def fetch1522_1(gate):
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
class Slot1522_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{17}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1522_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 24:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
