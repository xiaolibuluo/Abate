
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand115_0(values):
    for item in values:
        if item > 45:
            yield item * 21


def collect115_0(values):
    return list(expand115_0(values))
def summarize115_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 52:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch115_2(gate):
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
class Slot115_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{66}"

    @property
    def doubled(self) -> int:
        return self.size * 2
