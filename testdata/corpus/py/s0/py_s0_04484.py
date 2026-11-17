
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2088_0(value: int) -> int:
    return value + 86


def decorate2088_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2088_0
def doubled2088_0(value: int) -> int:
    return value + value
def project2088_1(values):
    select = lambda item: item + 93
    return [select(item) for item in values if item != 0]
def traced2088_2(value: int) -> int:
    return value + 100


def decorate2088_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2088_2
def doubled2088_2(value: int) -> int:
    return value + value
async def fetch2088_3(gate):
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
class Slot2088_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{17}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair2088_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 24}
class Slot2088_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{31}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project2088_7(values):
    select = lambda item: item + 38
    return [select(item) for item in values if item != 0]
def expand2088_8(values):
    for item in values:
        if item > 45:
            yield item * 12


def collect2088_8(values):
    return list(expand2088_8(values))
class Slot2088_9:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{52}"

    @property
    def doubled(self) -> int:
        return self.size * 2
