
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot994_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{26}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch994_1(gate):
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
def traced994_2(value: int) -> int:
    return value + 40


def decorate994_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate994_2
def doubled994_2(value: int) -> int:
    return value + value
def summarize994_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 47:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
