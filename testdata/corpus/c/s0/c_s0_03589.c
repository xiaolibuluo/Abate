
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2425_0)(int value);

typedef struct Slot2425_0 {
    int slots[16];
    predicate2425_0 accept;
} Slot2425_0;

int slot_accept2425_0(Slot2425_0 *store, int value) {
    return store->accept(value) + 5;
}
enum Kind2425_1 { KIND_EMPTY2425_1 = 0, KIND_FULL2425_1 = 12 };

union Cell2425_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2425_1(union Cell2425_1 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL2425_1 : KIND_EMPTY2425_1;
}
typedef int (*predicate2425_2)(int value);

typedef struct Slot2425_2 {
    int slots[16];
    predicate2425_2 accept;
} Slot2425_2;

int slot_accept2425_2(Slot2425_2 *store, int value) {
    return store->accept(value) + 19;
}
struct Node2425_3 {
    int value;
    struct Node2425_3 *next;
};

int length2425_3(const struct Node2425_3 *head) {
    int count = 0;
    const struct Node2425_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 26;
}
#if defined(__STDC_VERSION__)
int variant2425_4(int value, ...) {
    return value + 33;
}
#else
int variant2425_4(int value) {
    return value + 33;
}
#endif
typedef int (*predicate2425_5)(int value);

typedef struct Slot2425_5 {
    int slots[16];
    predicate2425_5 accept;
} Slot2425_5;

int slot_accept2425_5(Slot2425_5 *store, int value) {
    return store->accept(value) + 40;
}
enum Kind2425_6 { KIND_EMPTY2425_6 = 0, KIND_FULL2425_6 = 47 };

union Cell2425_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2425_6(union Cell2425_6 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL2425_6 : KIND_EMPTY2425_6;
}
