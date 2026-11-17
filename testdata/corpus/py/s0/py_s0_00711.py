
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot326_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{72}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project326_1(values):
    select = lambda item: item + 79
    return [select(item) for item in values if item != 0]
async def fetch326_2(gate):
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
def guarded326_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project326_4(values):
    select = lambda item: item + 100
    return [select(item) for item in values if item != 0]
def summarize326_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 10:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand326_6(values):
    for item in values:
        if item > 17:
            yield item * 39


def collect326_6(values):
    return list(expand326_6(values))
def traced326_7(value: int) -> int:
    return value + 24


def decorate326_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate326_7
def doubled326_7(value: int) -> int:
    return value + value
