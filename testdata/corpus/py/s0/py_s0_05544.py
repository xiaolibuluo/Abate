
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2564_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{66}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2564_1(value: int) -> int:
    return value + 73


def decorate2564_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2564_1
def doubled2564_1(value: int) -> int:
    return value + value
def summarize2564_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 80:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2564_3(value: int) -> int:
    return value + 87


def decorate2564_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2564_3
def doubled2564_3(value: int) -> int:
    return value + value
async def fetch2564_4(gate):
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
def summarize2564_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 101:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
