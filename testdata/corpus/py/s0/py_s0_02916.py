
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1337_0(gate):
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
def expand1337_1(values):
    for item in values:
        if item > 30:
            yield item * 46


def collect1337_1(values):
    return list(expand1337_1(values))
def expand1337_2(values):
    for item in values:
        if item > 37:
            yield item * 57


def collect1337_2(values):
    return list(expand1337_2(values))
def guarded1337_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1337_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{51}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project1337_5(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
