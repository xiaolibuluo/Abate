
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1264_0(values):
    for item in values:
        if item > 44:
            yield item * 12


def collect1264_0(values):
    return list(expand1264_0(values))
def expand1264_1(values):
    for item in values:
        if item > 51:
            yield item * 23


def collect1264_1(values):
    return list(expand1264_1(values))
def summarize1264_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 58:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1264_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1264_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1264_5(value: int) -> int:
    return value + 79


def decorate1264_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1264_5
def doubled1264_5(value: int) -> int:
    return value + value
def project1264_6(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
def traced1264_7(value: int) -> int:
    return value + 93


def decorate1264_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1264_7
def doubled1264_7(value: int) -> int:
    return value + value
def expand1264_8(values):
    for item in values:
        if item > 100:
            yield item * 39


def collect1264_8(values):
    return list(expand1264_8(values))
async def fetch1264_9(gate):
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
