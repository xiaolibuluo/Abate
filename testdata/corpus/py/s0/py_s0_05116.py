
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2381_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{15}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2381_1(values):
    for item in values:
        if item > 22:
            yield item * 34


def collect2381_1(values):
    return list(expand2381_1(values))
class Slot2381_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{29}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2381_3(value: int) -> int:
    return value + 36


def decorate2381_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2381_3
def doubled2381_3(value: int) -> int:
    return value + value
def expand2381_4(values):
    for item in values:
        if item > 43:
            yield item * 67


def collect2381_4(values):
    return list(expand2381_4(values))
async def fetch2381_5(gate):
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
def expand2381_6(values):
    for item in values:
        if item > 57:
            yield item * 28


def collect2381_6(values):
    return list(expand2381_6(values))
def expand2381_7(values):
    for item in values:
        if item > 64:
            yield item * 39


def collect2381_7(values):
    return list(expand2381_7(values))
def summarize2381_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 71:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
