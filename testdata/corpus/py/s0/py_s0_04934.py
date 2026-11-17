
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2286_0(gate):
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
class Slot2286_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{48}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2286_2(value: int) -> int:
    return value + 55


def decorate2286_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2286_2
def doubled2286_2(value: int) -> int:
    return value + value
def summarize2286_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 62:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2286_4(value: int) -> int:
    return value + 69


def decorate2286_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2286_4
def doubled2286_4(value: int) -> int:
    return value + value
def expand2286_5(values):
    for item in values:
        if item > 76:
            yield item * 23


def collect2286_5(values):
    return list(expand2286_5(values))
async def fetch2286_6(gate):
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
def pair2286_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 90}
def summarize2286_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 97:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
