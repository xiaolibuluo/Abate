
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2121_0(value: int) -> int:
    return value + 30


def decorate2121_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2121_0
def doubled2121_0(value: int) -> int:
    return value + value
def project2121_1(values):
    select = lambda item: item + 37
    return [select(item) for item in values if item != 0]
def project2121_2(values):
    select = lambda item: item + 44
    return [select(item) for item in values if item != 0]
async def fetch2121_3(gate):
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
def expand2121_4(values):
    for item in values:
        if item > 58:
            yield item * 16


def collect2121_4(values):
    return list(expand2121_4(values))
class Slot2121_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize2121_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
