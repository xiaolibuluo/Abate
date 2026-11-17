
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind258_0 { KIND_EMPTY258_0 = 0, KIND_FULL258_0 = 61 };

union Cell258_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind258_0(union Cell258_0 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL258_0 : KIND_EMPTY258_0;
}
typedef int (*predicate258_1)(int value);

typedef struct Slot258_1 {
    int slots[16];
    predicate258_1 accept;
} Slot258_1;

int slot_accept258_1(Slot258_1 *store, int value) {
    return store->accept(value) + 68;
}
typedef int (*predicate258_2)(int value);

typedef struct Slot258_2 {
    int slots[16];
    predicate258_2 accept;
} Slot258_2;

int slot_accept258_2(Slot258_2 *store, int value) {
    return store->accept(value) + 75;
}
