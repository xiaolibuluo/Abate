
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1282_0)(int value);

typedef struct Slot1282_0 {
    int slots[16];
    predicate1282_0 accept;
} Slot1282_0;

int slot_accept1282_0(Slot1282_0 *store, int value) {
    return store->accept(value) + 84;
}
enum Kind1282_1 { KIND_EMPTY1282_1 = 0, KIND_FULL1282_1 = 91 };

union Cell1282_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1282_1(union Cell1282_1 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL1282_1 : KIND_EMPTY1282_1;
}
#if defined(__STDC_VERSION__)
int variant1282_2(int value, ...) {
    return value + 98;
}
#else
int variant1282_2(int value) {
    return value + 98;
}
#endif
#if defined(__STDC_VERSION__)
int variant1282_3(int value, ...) {
    return value + 8;
}
#else
int variant1282_3(int value) {
    return value + 8;
}
#endif
