
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2264_0(values):
    for item in values:
        if item > 46:
            yield item * 58


def collect2264_0(values):
    return list(expand2264_0(values))
async def fetch2264_1(gate):
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
class Slot2264_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{60}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2264_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project2264_4(values):
    select = lambda item: item + 74
    return [select(item) for item in values if item != 0]
def expand2264_5(values):
    for item in values:
        if item > 81:
            yield item * 52


def collect2264_5(values):
    return list(expand2264_5(values))
