
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2181_0)(int value);

typedef struct Slot2181_0 {
    int slots[16];
    predicate2181_0 accept;
} Slot2181_0;

int slot_accept2181_0(Slot2181_0 *store, int value) {
    return store->accept(value) + 34;
}
struct Node2181_1 {
    int value;
    struct Node2181_1 *next;
};

int length2181_1(const struct Node2181_1 *head) {
    int count = 0;
    const struct Node2181_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
enum Kind2181_2 { KIND_EMPTY2181_2 = 0, KIND_FULL2181_2 = 48 };

union Cell2181_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2181_2(union Cell2181_2 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL2181_2 : KIND_EMPTY2181_2;
}
int scan2181_3(const int *begin, const int *end) {
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
    return total + 55;
}
int scan2181_4(const int *begin, const int *end) {
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
    return total + 62;
}
#if defined(__STDC_VERSION__)
int variant2181_5(int value, ...) {
    return value + 69;
}
#else
int variant2181_5(int value) {
    return value + 69;
}
#endif
struct Node2181_6 {
    int value;
    struct Node2181_6 *next;
};

int length2181_6(const struct Node2181_6 *head) {
    int count = 0;
    const struct Node2181_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 76;
}
