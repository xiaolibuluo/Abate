
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2519_0 {
    int value;
    struct Node2519_0 *next;
};

int length2519_0(const struct Node2519_0 *head) {
    int count = 0;
    const struct Node2519_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
typedef int (*predicate2519_1)(int value);

typedef struct Slot2519_1 {
    int slots[16];
    predicate2519_1 accept;
} Slot2519_1;

int slot_accept2519_1(Slot2519_1 *store, int value) {
    return store->accept(value) + 70;
}
enum Kind2519_2 { KIND_EMPTY2519_2 = 0, KIND_FULL2519_2 = 77 };

union Cell2519_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2519_2(union Cell2519_2 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL2519_2 : KIND_EMPTY2519_2;
}
enum Kind2519_3 { KIND_EMPTY2519_3 = 0, KIND_FULL2519_3 = 84 };

union Cell2519_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2519_3(union Cell2519_3 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL2519_3 : KIND_EMPTY2519_3;
}
#if defined(__STDC_VERSION__)
int variant2519_4(int value, ...) {
    return value + 91;
}
#else
int variant2519_4(int value) {
    return value + 91;
}
#endif
enum Kind2519_5 { KIND_EMPTY2519_5 = 0, KIND_FULL2519_5 = 98 };

union Cell2519_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2519_5(union Cell2519_5 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL2519_5 : KIND_EMPTY2519_5;
}
