
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand620_0(values):
    for item in values:
        if item > 14:
            yield item * 18


def collect620_0(values):
    return list(expand620_0(values))
def traced620_1(value: int) -> int:
    return value + 21


def decorate620_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate620_1
def doubled620_1(value: int) -> int:
    return value + value
async def fetch620_2(gate):
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
def traced620_3(value: int) -> int:
    return value + 35


def decorate620_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate620_3
def doubled620_3(value: int) -> int:
    return value + value
async def fetch620_4(gate):
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
def pair620_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 49}
class Slot620_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{56}"

    @property
    def doubled(self) -> int:
        return self.size * 2
