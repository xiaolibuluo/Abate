
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1116_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1116_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 67:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1116_2(gate):
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
async def fetch1116_3(gate):
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
def summarize1116_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 88:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1116_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 95}
class Slot1116_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{5}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1116_7(gate):
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
