
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1005_0(values):
    select = lambda item: item + 72
    return [select(item) for item in values if item != 0]
class Slot1005_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{79}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1005_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 86:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1005_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{93}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1005_4(value: int) -> int:
    return value + 100


def decorate1005_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1005_4
def doubled1005_4(value: int) -> int:
    return value + value
def expand1005_5(values):
    for item in values:
        if item > 10:
            yield item * 23


def collect1005_5(values):
    return list(expand1005_5(values))
async def fetch1005_6(gate):
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
