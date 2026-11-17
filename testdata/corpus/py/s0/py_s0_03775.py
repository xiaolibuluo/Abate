
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1742_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1742_1(values):
    select = lambda item: item + 57
    return [select(item) for item in values if item != 0]
def expand1742_2(values):
    for item in values:
        if item > 64:
            yield item * 25


def collect1742_2(values):
    return list(expand1742_2(values))
def guarded1742_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1742_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 78:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1742_5(values):
    for item in values:
        if item > 85:
            yield item * 58


def collect1742_5(values):
    return list(expand1742_5(values))
class Slot1742_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{92}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1742_7(gate):
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
def project1742_8(values):
    select = lambda item: item + 9
    return [select(item) for item in values if item != 0]
