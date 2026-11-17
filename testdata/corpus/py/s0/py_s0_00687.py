
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded318_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand318_1(values):
    for item in values:
        if item > 72:
            yield item * 50


def collect318_1(values):
    return list(expand318_1(values))
class Slot318_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{79}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair318_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 86}
async def fetch318_4(gate):
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
def traced318_5(value: int) -> int:
    return value + 100


def decorate318_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate318_5
def doubled318_5(value: int) -> int:
    return value + value
