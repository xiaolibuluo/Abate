
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1200_0(values):
    select = lambda item: item + 85
    return [select(item) for item in values if item != 0]
def guarded1200_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1200_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 99:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1200_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{9}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1200_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1200_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 23}
async def fetch1200_6(gate):
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
