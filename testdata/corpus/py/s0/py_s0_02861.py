
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1311_0(values):
    for item in values:
        if item > 73:
            yield item * 36


def collect1311_0(values):
    return list(expand1311_0(values))
class Slot1311_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{80}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1311_2(values):
    for item in values:
        if item > 87:
            yield item * 58


def collect1311_2(values):
    return list(expand1311_2(values))
class Slot1311_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{94}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1311_4(value: int) -> int:
    return value + 101


def decorate1311_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1311_4
def doubled1311_4(value: int) -> int:
    return value + value
async def fetch1311_5(gate):
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
def pair1311_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 18}
def traced1311_7(value: int) -> int:
    return value + 25


def decorate1311_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1311_7
def doubled1311_7(value: int) -> int:
    return value + value
def expand1311_8(values):
    for item in values:
        if item > 32:
            yield item * 63


def collect1311_8(values):
    return list(expand1311_8(values))
async def fetch1311_9(gate):
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
