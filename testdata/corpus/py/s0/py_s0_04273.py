
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1990_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{73}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project1990_1(values):
    select = lambda item: item + 80
    return [select(item) for item in values if item != 0]
def expand1990_2(values):
    for item in values:
        if item > 87:
            yield item * 53


def collect1990_2(values):
    return list(expand1990_2(values))
async def fetch1990_3(gate):
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
class Slot1990_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{101}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1990_5(values):
    for item in values:
        if item > 11:
            yield item * 25


def collect1990_5(values):
    return list(expand1990_5(values))
def expand1990_6(values):
    for item in values:
        if item > 18:
            yield item * 36


def collect1990_6(values):
    return list(expand1990_6(values))
def summarize1990_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 25:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1990_8(values):
    for item in values:
        if item > 32:
            yield item * 58


def collect1990_8(values):
    return list(expand1990_8(values))
def traced1990_9(value: int) -> int:
    return value + 39


def decorate1990_9(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1990_9
def doubled1990_9(value: int) -> int:
    return value + value
