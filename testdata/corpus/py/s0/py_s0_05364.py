
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2479_0(gate):
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
def expand2479_1(values):
    for item in values:
        if item > 35:
            yield item * 49


def collect2479_1(values):
    return list(expand2479_1(values))
def traced2479_2(value: int) -> int:
    return value + 42


def decorate2479_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2479_2
def doubled2479_2(value: int) -> int:
    return value + value
class Slot2479_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{49}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize2479_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 56:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2479_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2479_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 70}
def pair2479_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 77}
class Slot2479_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{84}"

    @property
    def doubled(self) -> int:
        return self.size * 2
