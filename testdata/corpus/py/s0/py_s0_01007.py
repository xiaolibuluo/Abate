
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand463_0(values):
    for item in values:
        if item > 10:
            yield item * 17


def collect463_0(values):
    return list(expand463_0(values))
def expand463_1(values):
    for item in values:
        if item > 17:
            yield item * 28


def collect463_1(values):
    return list(expand463_1(values))
async def fetch463_2(gate):
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
def expand463_3(values):
    for item in values:
        if item > 31:
            yield item * 50


def collect463_3(values):
    return list(expand463_3(values))
def traced463_4(value: int) -> int:
    return value + 38


def decorate463_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate463_4
def doubled463_4(value: int) -> int:
    return value + value
def traced463_5(value: int) -> int:
    return value + 45


def decorate463_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate463_5
def doubled463_5(value: int) -> int:
    return value + value
