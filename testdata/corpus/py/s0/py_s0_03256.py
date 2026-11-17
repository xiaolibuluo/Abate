
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1501_0(values):
    for item in values:
        if item > 21:
            yield item * 24


def collect1501_0(values):
    return list(expand1501_0(values))
async def fetch1501_1(gate):
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
def project1501_2(values):
    select = lambda item: item + 35
    return [select(item) for item in values if item != 0]
def traced1501_3(value: int) -> int:
    return value + 42


def decorate1501_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1501_3
def doubled1501_3(value: int) -> int:
    return value + value
async def fetch1501_4(gate):
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
class Slot1501_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{56}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair1501_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 63}
def expand1501_7(values):
    for item in values:
        if item > 70:
            yield item * 40


def collect1501_7(values):
    return list(expand1501_7(values))
def guarded1501_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
