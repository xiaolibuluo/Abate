
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1797_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 86:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1797_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 93:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1797_2(gate):
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
def traced1797_3(value: int) -> int:
    return value + 10


def decorate1797_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1797_3
def doubled1797_3(value: int) -> int:
    return value + value
def expand1797_4(values):
    for item in values:
        if item > 17:
            yield item * 66


def collect1797_4(values):
    return list(expand1797_4(values))
def pair1797_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 24}
class Slot1797_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{31}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1797_7(values):
    for item in values:
        if item > 38:
            yield item * 38


def collect1797_7(values):
    return list(expand1797_7(values))
def summarize1797_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 45:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
