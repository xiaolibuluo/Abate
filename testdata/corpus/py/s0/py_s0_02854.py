
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1307_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{21}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1307_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{28}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project1307_2(values):
    select = lambda item: item + 35
    return [select(item) for item in values if item != 0]
def expand1307_3(values):
    for item in values:
        if item > 42:
            yield item * 41


def collect1307_3(values):
    return list(expand1307_3(values))
def summarize1307_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 49:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1307_5(value: int) -> int:
    return value + 56


def decorate1307_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1307_5
def doubled1307_5(value: int) -> int:
    return value + value
class Slot1307_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{63}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1307_7(values):
    for item in values:
        if item > 70:
            yield item * 24


def collect1307_7(values):
    return list(expand1307_7(values))
async def fetch1307_8(gate):
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
