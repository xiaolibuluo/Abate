
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1531_0 { KIND_EMPTY1531_0 = 0, KIND_FULL1531_0 = 23 };

union Cell1531_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1531_0(union Cell1531_0 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1531_0 : KIND_EMPTY1531_0;
}
typedef int (*predicate1531_1)(int value);

typedef struct Slot1531_1 {
    int slots[16];
    predicate1531_1 accept;
} Slot1531_1;

int slot_accept1531_1(Slot1531_1 *store, int value) {
    return store->accept(value) + 30;
}
enum Kind1531_2 { KIND_EMPTY1531_2 = 0, KIND_FULL1531_2 = 37 };

union Cell1531_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1531_2(union Cell1531_2 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL1531_2 : KIND_EMPTY1531_2;
}
#if defined(__STDC_VERSION__)
int variant1531_3(int value, ...) {
    return value + 44;
}
#else
int variant1531_3(int value) {
    return value + 44;
}
#endif
typedef int (*predicate1531_4)(int value);

typedef struct Slot1531_4 {
    int slots[16];
    predicate1531_4 accept;
} Slot1531_4;

int slot_accept1531_4(Slot1531_4 *store, int value) {
    return store->accept(value) + 51;
}
