
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1097_0)(int value);

typedef struct Slot1097_0 {
    int slots[16];
    predicate1097_0 accept;
} Slot1097_0;

int slot_accept1097_0(Slot1097_0 *store, int value) {
    return store->accept(value) + 7;
}
typedef int (*predicate1097_1)(int value);

typedef struct Slot1097_1 {
    int slots[16];
    predicate1097_1 accept;
} Slot1097_1;

int slot_accept1097_1(Slot1097_1 *store, int value) {
    return store->accept(value) + 14;
}
enum Kind1097_2 { KIND_EMPTY1097_2 = 0, KIND_FULL1097_2 = 21 };

union Cell1097_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1097_2(union Cell1097_2 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL1097_2 : KIND_EMPTY1097_2;
}
enum Kind1097_3 { KIND_EMPTY1097_3 = 0, KIND_FULL1097_3 = 28 };

union Cell1097_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1097_3(union Cell1097_3 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL1097_3 : KIND_EMPTY1097_3;
}
enum Kind1097_4 { KIND_EMPTY1097_4 = 0, KIND_FULL1097_4 = 35 };

union Cell1097_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1097_4(union Cell1097_4 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL1097_4 : KIND_EMPTY1097_4;
}
