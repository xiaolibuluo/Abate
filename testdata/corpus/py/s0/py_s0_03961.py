
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1836_0(value: int) -> int:
    return value + 11


def decorate1836_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1836_0
def doubled1836_0(value: int) -> int:
    return value + value
def expand1836_1(values):
    for item in values:
        if item > 18:
            yield item * 62


def collect1836_1(values):
    return list(expand1836_1(values))
class Slot1836_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{25}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1836_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{32}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair1836_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 39}
class Slot1836_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{46}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair1836_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 53}
async def fetch1836_7(gate):
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
