
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1259_0(gate):
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
class Slot1259_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{83}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1259_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{90}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1259_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{97}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair1259_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 7}
def guarded1259_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1259_6(value: int) -> int:
    return value + 21


def decorate1259_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1259_6
def doubled1259_6(value: int) -> int:
    return value + value
def summarize1259_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 28:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
