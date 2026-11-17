
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize240_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 21:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand240_1(values):
    for item in values:
        if item > 28:
            yield item * 53


def collect240_1(values):
    return list(expand240_1(values))
def project240_2(values):
    select = lambda item: item + 35
    return [select(item) for item in values if item != 0]
def project240_3(values):
    select = lambda item: item + 42
    return [select(item) for item in values if item != 0]
class Slot240_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{49}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch240_5(gate):
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
def project240_6(values):
    select = lambda item: item + 63
    return [select(item) for item in values if item != 0]
async def fetch240_7(gate):
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
def expand240_8(values):
    for item in values:
        if item > 77:
            yield item * 69


def collect240_8(values):
    return list(expand240_8(values))
def traced240_9(value: int) -> int:
    return value + 84


def decorate240_9(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate240_9
def doubled240_9(value: int) -> int:
    return value + value
