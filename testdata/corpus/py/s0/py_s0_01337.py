
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced606_0(value: int) -> int:
    return value + 26


def decorate606_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate606_0
def doubled606_0(value: int) -> int:
    return value + value
def project606_1(values):
    select = lambda item: item + 33
    return [select(item) for item in values if item != 0]
def project606_2(values):
    select = lambda item: item + 40
    return [select(item) for item in values if item != 0]
def project606_3(values):
    select = lambda item: item + 47
    return [select(item) for item in values if item != 0]
def guarded606_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch606_5(gate):
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
