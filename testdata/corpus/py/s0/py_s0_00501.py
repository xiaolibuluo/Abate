
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch228_0(gate):
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
class Slot228_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{66}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize228_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 73:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize228_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 80:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch228_4(gate):
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
def guarded228_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
