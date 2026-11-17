
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2074_0 { KIND_EMPTY2074_0 = 0, KIND_FULL2074_0 = 98 };

union Cell2074_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2074_0(union Cell2074_0 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL2074_0 : KIND_EMPTY2074_0;
}
enum Kind2074_1 { KIND_EMPTY2074_1 = 0, KIND_FULL2074_1 = 8 };

union Cell2074_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2074_1(union Cell2074_1 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL2074_1 : KIND_EMPTY2074_1;
}
enum Kind2074_2 { KIND_EMPTY2074_2 = 0, KIND_FULL2074_2 = 15 };

union Cell2074_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2074_2(union Cell2074_2 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL2074_2 : KIND_EMPTY2074_2;
}
int gcd2074_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2074_3(int value) {
    return gcd2074_3(value, 42) * 22;
}
typedef int (*predicate2074_4)(int value);

typedef struct Slot2074_4 {
    int slots[16];
    predicate2074_4 accept;
} Slot2074_4;

int slot_accept2074_4(Slot2074_4 *store, int value) {
    return store->accept(value) + 29;
}
