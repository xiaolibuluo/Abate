
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2569_0 {
    int value;
    struct Node2569_0 *next;
};

int length2569_0(const struct Node2569_0 *head) {
    int count = 0;
    const struct Node2569_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
#if defined(__STDC_VERSION__)
int variant2569_1(int value, ...) {
    return value + 41;
}
#else
int variant2569_1(int value) {
    return value + 41;
}
#endif
enum Kind2569_2 { KIND_EMPTY2569_2 = 0, KIND_FULL2569_2 = 48 };

union Cell2569_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2569_2(union Cell2569_2 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL2569_2 : KIND_EMPTY2569_2;
}
typedef int (*predicate2569_3)(int value);

typedef struct Slot2569_3 {
    int slots[16];
    predicate2569_3 accept;
} Slot2569_3;

int slot_accept2569_3(Slot2569_3 *store, int value) {
    return store->accept(value) + 55;
}
enum Kind2569_4 { KIND_EMPTY2569_4 = 0, KIND_FULL2569_4 = 62 };

union Cell2569_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2569_4(union Cell2569_4 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL2569_4 : KIND_EMPTY2569_4;
}
enum Kind2569_5 { KIND_EMPTY2569_5 = 0, KIND_FULL2569_5 = 69 };

union Cell2569_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2569_5(union Cell2569_5 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL2569_5 : KIND_EMPTY2569_5;
}
typedef int (*predicate2569_6)(int value);

typedef struct Slot2569_6 {
    int slots[16];
    predicate2569_6 accept;
} Slot2569_6;

int slot_accept2569_6(Slot2569_6 *store, int value) {
    return store->accept(value) + 76;
}
