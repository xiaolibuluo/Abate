
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1410_0(values):
    for item in values:
        if item > 99:
            yield item * 58


def collect1410_0(values):
    return list(expand1410_0(values))
class Slot1410_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{9}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1410_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 16:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1410_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 23:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1410_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1410_5(gate):
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
def traced1410_6(value: int) -> int:
    return value + 44


def decorate1410_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1410_6
def doubled1410_6(value: int) -> int:
    return value + value
def expand1410_7(values):
    for item in values:
        if item > 51:
            yield item * 13


def collect1410_7(values):
    return list(expand1410_7(values))
def expand1410_8(values):
    for item in values:
        if item > 58:
            yield item * 24


def collect1410_8(values):
    return list(expand1410_8(values))
