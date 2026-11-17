
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1892_0(value: int) -> int:
    return value + 60


def decorate1892_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1892_0
def doubled1892_0(value: int) -> int:
    return value + value
def summarize1892_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 67:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1892_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{74}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1892_3(gate):
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
async def fetch1892_4(gate):
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
class Slot1892_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{95}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1892_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{5}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1892_7(value: int) -> int:
    return value + 12


def decorate1892_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1892_7
def doubled1892_7(value: int) -> int:
    return value + value
