
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded714_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced714_1(value: int) -> int:
    return value + 79


def decorate714_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate714_1
def doubled714_1(value: int) -> int:
    return value + value
def project714_2(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
def project714_3(values):
    select = lambda item: item + 93
    return [select(item) for item in values if item != 0]
async def fetch714_4(gate):
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
class Slot714_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{10}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize714_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 17:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
