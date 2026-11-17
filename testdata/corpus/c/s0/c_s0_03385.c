
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2284_0 {
    int value;
    struct Node2284_0 *next;
};

int length2284_0(const struct Node2284_0 *head) {
    int count = 0;
    const struct Node2284_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
typedef int (*predicate2284_1)(int value);

typedef struct Slot2284_1 {
    int slots[16];
    predicate2284_1 accept;
} Slot2284_1;

int slot_accept2284_1(Slot2284_1 *store, int value) {
    return store->accept(value) + 22;
}
#if defined(__STDC_VERSION__)
int variant2284_2(int value, ...) {
    return value + 29;
}
#else
int variant2284_2(int value) {
    return value + 29;
}
#endif
struct Node2284_3 {
    int value;
    struct Node2284_3 *next;
};

int length2284_3(const struct Node2284_3 *head) {
    int count = 0;
    const struct Node2284_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 36;
}
enum Kind2284_4 { KIND_EMPTY2284_4 = 0, KIND_FULL2284_4 = 43 };

union Cell2284_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2284_4(union Cell2284_4 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL2284_4 : KIND_EMPTY2284_4;
}
int scan2284_5(const int *begin, const int *end) {
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
    return total + 50;
}
int scan2284_6(const int *begin, const int *end) {
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
    return total + 57;
}
