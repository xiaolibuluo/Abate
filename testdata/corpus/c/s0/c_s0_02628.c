
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1782_0 { KIND_EMPTY1782_0 = 0, KIND_FULL1782_0 = 85 };

union Cell1782_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1782_0(union Cell1782_0 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL1782_0 : KIND_EMPTY1782_0;
}
typedef int (*predicate1782_1)(int value);

typedef struct Slot1782_1 {
    int slots[16];
    predicate1782_1 accept;
} Slot1782_1;

int slot_accept1782_1(Slot1782_1 *store, int value) {
    return store->accept(value) + 92;
}
enum Kind1782_2 { KIND_EMPTY1782_2 = 0, KIND_FULL1782_2 = 99 };

union Cell1782_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1782_2(union Cell1782_2 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL1782_2 : KIND_EMPTY1782_2;
}
enum Kind1782_3 { KIND_EMPTY1782_3 = 0, KIND_FULL1782_3 = 9 };

union Cell1782_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1782_3(union Cell1782_3 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL1782_3 : KIND_EMPTY1782_3;
}
typedef int (*predicate1782_4)(int value);

typedef struct Slot1782_4 {
    int slots[16];
    predicate1782_4 accept;
} Slot1782_4;

int slot_accept1782_4(Slot1782_4 *store, int value) {
    return store->accept(value) + 16;
}
