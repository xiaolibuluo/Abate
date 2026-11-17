
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1727_0(values):
    for item in values:
        if item > 49:
            yield item * 20


def collect1727_0(values):
    return list(expand1727_0(values))
def expand1727_1(values):
    for item in values:
        if item > 56:
            yield item * 31


def collect1727_1(values):
    return list(expand1727_1(values))
def guarded1727_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1727_3(values):
    for item in values:
        if item > 70:
            yield item * 53


def collect1727_3(values):
    return list(expand1727_3(values))
async def fetch1727_4(gate):
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
class Slot1727_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{84}"

    @property
    def doubled(self) -> int:
        return self.size * 2
