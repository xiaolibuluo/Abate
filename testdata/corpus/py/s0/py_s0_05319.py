
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2461_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{85}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair2461_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 92}
def expand2461_2(values):
    for item in values:
        if item > 99:
            yield item * 56


def collect2461_2(values):
    return list(expand2461_2(values))
async def fetch2461_3(gate):
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
def guarded2461_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2461_5(value: int) -> int:
    return value + 23


def decorate2461_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2461_5
def doubled2461_5(value: int) -> int:
    return value + value
