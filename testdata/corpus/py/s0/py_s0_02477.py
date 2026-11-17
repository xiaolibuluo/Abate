
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1149_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1149_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 11}
def traced1149_2(value: int) -> int:
    return value + 18


def decorate1149_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1149_2
def doubled1149_2(value: int) -> int:
    return value + value
def expand1149_3(values):
    for item in values:
        if item > 25:
            yield item * 33


def collect1149_3(values):
    return list(expand1149_3(values))
def traced1149_4(value: int) -> int:
    return value + 32


def decorate1149_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1149_4
def doubled1149_4(value: int) -> int:
    return value + value
class Slot1149_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{39}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1149_6(gate):
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
