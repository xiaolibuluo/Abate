
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced373_0(value: int) -> int:
    return value + 101


def decorate373_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate373_0
def doubled373_0(value: int) -> int:
    return value + value
async def fetch373_1(gate):
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
def expand373_2(values):
    for item in values:
        if item > 18:
            yield item * 19


def collect373_2(values):
    return list(expand373_2(values))
class Slot373_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{25}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot373_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{32}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced373_5(value: int) -> int:
    return value + 39


def decorate373_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate373_5
def doubled373_5(value: int) -> int:
    return value + value
def pair373_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 46}
