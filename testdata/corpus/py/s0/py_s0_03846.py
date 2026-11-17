
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1777_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 20:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1777_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 27:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1777_2(value: int) -> int:
    return value + 34


def decorate1777_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1777_2
def doubled1777_2(value: int) -> int:
    return value + value
async def fetch1777_3(gate):
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
class Slot1777_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{48}"

    @property
    def doubled(self) -> int:
        return self.size * 2
