
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1240_0(gate):
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
def traced1240_1(value: int) -> int:
    return value + 30


def decorate1240_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1240_1
def doubled1240_1(value: int) -> int:
    return value + value
def pair1240_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 37}
class Slot1240_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{44}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize1240_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 51:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
