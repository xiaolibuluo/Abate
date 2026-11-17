
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced237_0(value: int) -> int:
    return value + 79


def decorate237_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate237_0
def doubled237_0(value: int) -> int:
    return value + value
class Slot237_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{86}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project237_2(values):
    select = lambda item: item + 93
    return [select(item) for item in values if item != 0]
async def fetch237_3(gate):
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
def expand237_4(values):
    for item in values:
        if item > 10:
            yield item * 65


def collect237_4(values):
    return list(expand237_4(values))
async def fetch237_5(gate):
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
def summarize237_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 24:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced237_7(value: int) -> int:
    return value + 31


def decorate237_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate237_7
def doubled237_7(value: int) -> int:
    return value + value
