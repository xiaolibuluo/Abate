
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand826_0(values):
    for item in values:
        if item > 73:
            yield item * 57


def collect826_0(values):
    return list(expand826_0(values))
def summarize826_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 80:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced826_2(value: int) -> int:
    return value + 87


def decorate826_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate826_2
def doubled826_2(value: int) -> int:
    return value + value
async def fetch826_3(gate):
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
def summarize826_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 101:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch826_5(gate):
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
async def fetch826_6(gate):
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
class Slot826_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{25}"

    @property
    def doubled(self) -> int:
        return self.size * 2
