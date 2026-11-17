
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1589_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{98}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1589_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{8}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project1589_2(values):
    select = lambda item: item + 15
    return [select(item) for item in values if item != 0]
def traced1589_3(value: int) -> int:
    return value + 22


def decorate1589_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1589_3
def doubled1589_3(value: int) -> int:
    return value + value
async def fetch1589_4(gate):
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
def summarize1589_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 36:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1589_6(gate):
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
