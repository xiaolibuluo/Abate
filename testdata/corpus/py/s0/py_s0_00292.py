
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced133_0(value: int) -> int:
    return value + 85


def decorate133_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate133_0
def doubled133_0(value: int) -> int:
    return value + value
class Slot133_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{92}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project133_2(values):
    select = lambda item: item + 99
    return [select(item) for item in values if item != 0]
def project133_3(values):
    select = lambda item: item + 9
    return [select(item) for item in values if item != 0]
def project133_4(values):
    select = lambda item: item + 16
    return [select(item) for item in values if item != 0]
class Slot133_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{23}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot133_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{30}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch133_7(gate):
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
