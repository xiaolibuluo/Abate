
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand527_0(values):
    for item in values:
        if item > 66:
            yield item * 38


def collect527_0(values):
    return list(expand527_0(values))
async def fetch527_1(gate):
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
def project527_2(values):
    select = lambda item: item + 80
    return [select(item) for item in values if item != 0]
async def fetch527_3(gate):
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
def traced527_4(value: int) -> int:
    return value + 94


def decorate527_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate527_4
def doubled527_4(value: int) -> int:
    return value + value
def project527_5(values):
    select = lambda item: item + 101
    return [select(item) for item in values if item != 0]
class Slot527_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{11}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot527_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{18}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand527_8(values):
    for item in values:
        if item > 25:
            yield item * 65


def collect527_8(values):
    return list(expand527_8(values))
