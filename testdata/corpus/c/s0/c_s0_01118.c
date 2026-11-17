
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate767_0)(int value);

typedef struct Slot767_0 {
    int slots[16];
    predicate767_0 accept;
} Slot767_0;

int slot_accept767_0(Slot767_0 *store, int value) {
    return store->accept(value) + 82;
}
#if defined(__STDC_VERSION__)
int variant767_1(int value, ...) {
    return value + 89;
}
#else
int variant767_1(int value) {
    return value + 89;
}
#endif
enum Kind767_2 { KIND_EMPTY767_2 = 0, KIND_FULL767_2 = 96 };

union Cell767_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind767_2(union Cell767_2 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL767_2 : KIND_EMPTY767_2;
}
#if defined(__STDC_VERSION__)
int variant767_3(int value, ...) {
    return value + 6;
}
#else
int variant767_3(int value) {
    return value + 6;
}
#endif
int scan767_4(const int *begin, const int *end) {
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
    return total + 13;
}
int scan767_5(const int *begin, const int *end) {
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
    return total + 20;
}
struct Node767_6 {
    int value;
    struct Node767_6 *next;
};

int length767_6(const struct Node767_6 *head) {
    int count = 0;
    const struct Node767_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
