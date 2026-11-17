
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1016_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1016_1(value: int) -> int:
    return value + 28


def decorate1016_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1016_1
def doubled1016_1(value: int) -> int:
    return value + value
class Slot1016_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{35}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1016_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{42}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1016_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 49:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1016_5(values):
    select = lambda item: item + 56
    return [select(item) for item in values if item != 0]
def expand1016_6(values):
    for item in values:
        if item > 63:
            yield item * 50


def collect1016_6(values):
    return list(expand1016_6(values))
class Slot1016_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{70}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1016_8(gate):
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
