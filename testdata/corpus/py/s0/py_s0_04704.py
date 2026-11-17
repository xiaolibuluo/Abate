
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2179_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 8:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2179_1(value: int) -> int:
    return value + 15


def decorate2179_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2179_1
def doubled2179_1(value: int) -> int:
    return value + value
def traced2179_2(value: int) -> int:
    return value + 22


def decorate2179_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2179_2
def doubled2179_2(value: int) -> int:
    return value + value
async def fetch2179_3(gate):
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
def project2179_4(values):
    select = lambda item: item + 36
    return [select(item) for item in values if item != 0]
async def fetch2179_5(gate):
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
class Slot2179_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{50}"

    @property
    def doubled(self) -> int:
        return self.size * 2
