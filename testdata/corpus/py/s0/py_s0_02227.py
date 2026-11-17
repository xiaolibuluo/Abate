
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1015_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1015_1(values):
    select = lambda item: item + 15
    return [select(item) for item in values if item != 0]
class Slot1015_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{22}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1015_3(value: int) -> int:
    return value + 29


def decorate1015_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1015_3
def doubled1015_3(value: int) -> int:
    return value + value
async def fetch1015_4(gate):
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
class Slot1015_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{43}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1015_6(values):
    for item in values:
        if item > 50:
            yield item * 43


def collect1015_6(values):
    return list(expand1015_6(values))
def summarize1015_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 57:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1015_8(value: int) -> int:
    return value + 64


def decorate1015_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1015_8
def doubled1015_8(value: int) -> int:
    return value + value
