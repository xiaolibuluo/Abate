
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2425_0(values):
    for item in values:
        if item > 5:
            yield item * 26


def collect2425_0(values):
    return list(expand2425_0(values))
def summarize2425_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 12:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2425_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{19}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch2425_3(gate):
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
def traced2425_4(value: int) -> int:
    return value + 33


def decorate2425_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2425_4
def doubled2425_4(value: int) -> int:
    return value + value
def project2425_5(values):
    select = lambda item: item + 40
    return [select(item) for item in values if item != 0]
def project2425_6(values):
    select = lambda item: item + 47
    return [select(item) for item in values if item != 0]
def traced2425_7(value: int) -> int:
    return value + 54


def decorate2425_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2425_7
def doubled2425_7(value: int) -> int:
    return value + value
def expand2425_8(values):
    for item in values:
        if item > 61:
            yield item * 53


def collect2425_8(values):
    return list(expand2425_8(values))
