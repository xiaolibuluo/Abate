
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project903_0(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
async def fetch903_1(gate):
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
class Slot903_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{21}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced903_3(value: int) -> int:
    return value + 28


def decorate903_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate903_3
def doubled903_3(value: int) -> int:
    return value + value
def summarize903_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 35:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced903_5(value: int) -> int:
    return value + 42


def decorate903_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate903_5
def doubled903_5(value: int) -> int:
    return value + value
