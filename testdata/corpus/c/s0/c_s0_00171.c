
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant117_0(int value, ...) {
    return value + 71;
}
#else
int variant117_0(int value) {
    return value + 71;
}
#endif
enum Kind117_1 { KIND_EMPTY117_1 = 0, KIND_FULL117_1 = 78 };

union Cell117_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind117_1(union Cell117_1 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL117_1 : KIND_EMPTY117_1;
}
typedef int (*predicate117_2)(int value);

typedef struct Slot117_2 {
    int slots[16];
    predicate117_2 accept;
} Slot117_2;

int slot_accept117_2(Slot117_2 *store, int value) {
    return store->accept(value) + 85;
}
enum Kind117_3 { KIND_EMPTY117_3 = 0, KIND_FULL117_3 = 92 };

union Cell117_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind117_3(union Cell117_3 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL117_3 : KIND_EMPTY117_3;
}
