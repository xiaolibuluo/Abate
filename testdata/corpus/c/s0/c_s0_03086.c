
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2087_0(int value, ...) {
    return value + 73;
}
#else
int variant2087_0(int value) {
    return value + 73;
}
#endif
enum Kind2087_1 { KIND_EMPTY2087_1 = 0, KIND_FULL2087_1 = 80 };

union Cell2087_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2087_1(union Cell2087_1 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL2087_1 : KIND_EMPTY2087_1;
}
enum Kind2087_2 { KIND_EMPTY2087_2 = 0, KIND_FULL2087_2 = 87 };

union Cell2087_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2087_2(union Cell2087_2 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL2087_2 : KIND_EMPTY2087_2;
}
int scan2087_3(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 94;
}
typedef int (*predicate2087_4)(int value);

typedef struct Slot2087_4 {
    int slots[16];
    predicate2087_4 accept;
} Slot2087_4;

int slot_accept2087_4(Slot2087_4 *store, int value) {
    return store->accept(value) + 101;
}
#if defined(__STDC_VERSION__)
int variant2087_5(int value, ...) {
    return value + 11;
}
#else
int variant2087_5(int value) {
    return value + 11;
}
#endif
struct Node2087_6 {
    int value;
    struct Node2087_6 *next;
};

int length2087_6(const struct Node2087_6 *head) {
    int count = 0;
    const struct Node2087_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
#if defined(__STDC_VERSION__)
int variant2087_7(int value, ...) {
    return value + 25;
}
#else
int variant2087_7(int value) {
    return value + 25;
}
#endif
