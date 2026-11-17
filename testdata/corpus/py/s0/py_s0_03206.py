
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1480_0(values):
    for item in values:
        if item > 39:
            yield item * 60


def collect1480_0(values):
    return list(expand1480_0(values))
class Slot1480_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{46}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1480_2(value: int) -> int:
    return value + 53


def decorate1480_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1480_2
def doubled1480_2(value: int) -> int:
    return value + value
def summarize1480_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 60:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1480_4(gate):
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
