
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1151_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1151_1(values):
    for item in values:
        if item > 37:
            yield item * 25


def collect1151_1(values):
    return list(expand1151_1(values))
async def fetch1151_2(gate):
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
class Slot1151_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{51}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1151_4(value: int) -> int:
    return value + 58


def decorate1151_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1151_4
def doubled1151_4(value: int) -> int:
    return value + value
def summarize1151_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 65:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
