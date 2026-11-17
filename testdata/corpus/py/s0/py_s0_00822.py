
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project374_0(values):
    select = lambda item: item + 17
    return [select(item) for item in values if item != 0]
class Slot374_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{24}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded374_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project374_3(values):
    select = lambda item: item + 38
    return [select(item) for item in values if item != 0]
def summarize374_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 45:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand374_5(values):
    for item in values:
        if item > 52:
            yield item * 59


def collect374_5(values):
    return list(expand374_5(values))
async def fetch374_6(gate):
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
def project374_7(values):
    select = lambda item: item + 66
    return [select(item) for item in values if item != 0]
