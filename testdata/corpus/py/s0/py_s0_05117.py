
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2382_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2382_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2382_2(gate):
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
def pair2382_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 49}
def summarize2382_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 56:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2382_5(value: int) -> int:
    return value + 63


def decorate2382_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2382_5
def doubled2382_5(value: int) -> int:
    return value + value
def pair2382_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 70}
def expand2382_7(values):
    for item in values:
        if item > 77:
            yield item * 46


def collect2382_7(values):
    return list(expand2382_7(values))
def summarize2382_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 84:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2382_9:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{91}"

    @property
    def doubled(self) -> int:
        return self.size * 2
