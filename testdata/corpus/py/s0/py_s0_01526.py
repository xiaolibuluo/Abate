
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand695_0(values):
    for item in values:
        if item > 19:
            yield item * 55


def collect695_0(values):
    return list(expand695_0(values))
async def fetch695_1(gate):
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
def project695_2(values):
    select = lambda item: item + 33
    return [select(item) for item in values if item != 0]
def traced695_3(value: int) -> int:
    return value + 40


def decorate695_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate695_3
def doubled695_3(value: int) -> int:
    return value + value
