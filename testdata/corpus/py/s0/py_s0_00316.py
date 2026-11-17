
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded144_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced144_1(value: int) -> int:
    return value + 41


def decorate144_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate144_1
def doubled144_1(value: int) -> int:
    return value + value
def summarize144_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 48:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize144_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 55:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch144_4(gate):
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
def expand144_5(values):
    for item in values:
        if item > 69:
            yield item * 35


def collect144_5(values):
    return list(expand144_5(values))
class Slot144_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{76}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project144_7(values):
    select = lambda item: item + 83
    return [select(item) for item in values if item != 0]
def guarded144_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
