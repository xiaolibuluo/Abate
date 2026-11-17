
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1620_0(values):
    select = lambda item: item + 16
    return [select(item) for item in values if item != 0]
def traced1620_1(value: int) -> int:
    return value + 23


def decorate1620_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1620_1
def doubled1620_1(value: int) -> int:
    return value + value
def guarded1620_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1620_3(gate):
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
