
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2271_0(gate):
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
class Slot2271_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{47}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2271_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2271_3(value: int) -> int:
    return value + 61


def decorate2271_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2271_3
def doubled2271_3(value: int) -> int:
    return value + value
def project2271_4(values):
    select = lambda item: item + 68
    return [select(item) for item in values if item != 0]
def expand2271_5(values):
    for item in values:
        if item > 75:
            yield item * 40


def collect2271_5(values):
    return list(expand2271_5(values))
def summarize2271_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 82:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
