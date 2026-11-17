
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2412_0)(int value);

typedef struct Slot2412_0 {
    int slots[16];
    predicate2412_0 accept;
} Slot2412_0;

int slot_accept2412_0(Slot2412_0 *store, int value) {
    return store->accept(value) + 30;
}
#if defined(__STDC_VERSION__)
int variant2412_1(int value, ...) {
    return value + 37;
}
#else
int variant2412_1(int value) {
    return value + 37;
}
#endif
struct Node2412_2 {
    int value;
    struct Node2412_2 *next;
};

int length2412_2(const struct Node2412_2 *head) {
    int count = 0;
    const struct Node2412_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
int scan2412_3(const int *begin, const int *end) {
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
    return total + 51;
}
enum Kind2412_4 { KIND_EMPTY2412_4 = 0, KIND_FULL2412_4 = 58 };

union Cell2412_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2412_4(union Cell2412_4 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL2412_4 : KIND_EMPTY2412_4;
}
