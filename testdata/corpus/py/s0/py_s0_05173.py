
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2410_0(values):
    select = lambda item: item + 101
    return [select(item) for item in values if item != 0]
async def fetch2410_1(gate):
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
class Slot2410_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{18}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2410_3(value: int) -> int:
    return value + 25


def decorate2410_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2410_3
def doubled2410_3(value: int) -> int:
    return value + value
async def fetch2410_4(gate):
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
class Slot2410_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{39}"

    @property
    def doubled(self) -> int:
        return self.size * 2
