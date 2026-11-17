
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1349_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1349_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 89:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1349_2(gate):
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
async def fetch1349_3(gate):
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
def pair1349_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
def guarded1349_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1349_6(values):
    select = lambda item: item + 27
    return [select(item) for item in values if item != 0]
class Slot1349_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{34}"

    @property
    def doubled(self) -> int:
        return self.size * 2
