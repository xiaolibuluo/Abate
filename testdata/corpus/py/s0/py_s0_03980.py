
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1845_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 31:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1845_1(values):
    for item in values:
        if item > 38:
            yield item * 64


def collect1845_1(values):
    return list(expand1845_1(values))
def expand1845_2(values):
    for item in values:
        if item > 45:
            yield item * 14


def collect1845_2(values):
    return list(expand1845_2(values))
def expand1845_3(values):
    for item in values:
        if item > 52:
            yield item * 25


def collect1845_3(values):
    return list(expand1845_3(values))
def traced1845_4(value: int) -> int:
    return value + 59


def decorate1845_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1845_4
def doubled1845_4(value: int) -> int:
    return value + value
def guarded1845_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1845_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{73}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1845_7(gate):
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
class Slot1845_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{87}"

    @property
    def doubled(self) -> int:
        return self.size * 2
