
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1055_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{43}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair1055_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 50}
async def fetch1055_2(gate):
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
def guarded1055_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1055_4(value: int) -> int:
    return value + 71


def decorate1055_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1055_4
def doubled1055_4(value: int) -> int:
    return value + value
def expand1055_5(values):
    for item in values:
        if item > 78:
            yield item * 68


def collect1055_5(values):
    return list(expand1055_5(values))
def expand1055_6(values):
    for item in values:
        if item > 85:
            yield item * 18


def collect1055_6(values):
    return list(expand1055_6(values))
