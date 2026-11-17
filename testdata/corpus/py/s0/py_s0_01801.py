
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch825_0(gate):
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
class Slot825_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{67}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand825_2(values):
    for item in values:
        if item > 74:
            yield item * 11


def collect825_2(values):
    return list(expand825_2(values))
def traced825_3(value: int) -> int:
    return value + 81


def decorate825_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate825_3
def doubled825_3(value: int) -> int:
    return value + value
class Slot825_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{88}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced825_5(value: int) -> int:
    return value + 95


def decorate825_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate825_5
def doubled825_5(value: int) -> int:
    return value + value
def summarize825_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 5:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch825_7(gate):
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
def summarize825_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 19:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
