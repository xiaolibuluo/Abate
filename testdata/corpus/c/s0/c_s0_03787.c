
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2563_0(int value, ...) {
    return value + 53;
}
#else
int variant2563_0(int value) {
    return value + 53;
}
#endif
typedef int (*predicate2563_1)(int value);

typedef struct Slot2563_1 {
    int slots[16];
    predicate2563_1 accept;
} Slot2563_1;

int slot_accept2563_1(Slot2563_1 *store, int value) {
    return store->accept(value) + 60;
}
enum Kind2563_2 { KIND_EMPTY2563_2 = 0, KIND_FULL2563_2 = 67 };

union Cell2563_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2563_2(union Cell2563_2 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL2563_2 : KIND_EMPTY2563_2;
}
enum Kind2563_3 { KIND_EMPTY2563_3 = 0, KIND_FULL2563_3 = 74 };

union Cell2563_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2563_3(union Cell2563_3 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL2563_3 : KIND_EMPTY2563_3;
}
#if defined(__STDC_VERSION__)
int variant2563_4(int value, ...) {
    return value + 81;
}
#else
int variant2563_4(int value) {
    return value + 81;
}
#endif
struct Node2563_5 {
    int value;
    struct Node2563_5 *next;
};

int length2563_5(const struct Node2563_5 *head) {
    int count = 0;
    const struct Node2563_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
