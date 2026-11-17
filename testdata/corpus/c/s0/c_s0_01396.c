
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind957_0 { KIND_EMPTY957_0 = 0, KIND_FULL957_0 = 30 };

union Cell957_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind957_0(union Cell957_0 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL957_0 : KIND_EMPTY957_0;
}
typedef int (*predicate957_1)(int value);

typedef struct Slot957_1 {
    int slots[16];
    predicate957_1 accept;
} Slot957_1;

int slot_accept957_1(Slot957_1 *store, int value) {
    return store->accept(value) + 37;
}
enum Kind957_2 { KIND_EMPTY957_2 = 0, KIND_FULL957_2 = 44 };

union Cell957_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind957_2(union Cell957_2 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL957_2 : KIND_EMPTY957_2;
}
typedef int (*predicate957_3)(int value);

typedef struct Slot957_3 {
    int slots[16];
    predicate957_3 accept;
} Slot957_3;

int slot_accept957_3(Slot957_3 *store, int value) {
    return store->accept(value) + 51;
}
typedef int (*predicate957_4)(int value);

typedef struct Slot957_4 {
    int slots[16];
    predicate957_4 accept;
} Slot957_4;

int slot_accept957_4(Slot957_4 *store, int value) {
    return store->accept(value) + 58;
}
#if defined(__STDC_VERSION__)
int variant957_5(int value, ...) {
    return value + 65;
}
#else
int variant957_5(int value) {
    return value + 65;
}
#endif
