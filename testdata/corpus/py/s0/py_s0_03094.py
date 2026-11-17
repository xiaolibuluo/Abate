
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1425_0(values):
    for item in values:
        if item > 100:
            yield item * 41


def collect1425_0(values):
    return list(expand1425_0(values))
def traced1425_1(value: int) -> int:
    return value + 10


def decorate1425_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1425_1
def doubled1425_1(value: int) -> int:
    return value + value
def project1425_2(values):
    select = lambda item: item + 17
    return [select(item) for item in values if item != 0]
async def fetch1425_3(gate):
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
