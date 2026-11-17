
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node482_0 {
    int value;
    struct Node482_0 *next;
};

int length482_0(const struct Node482_0 *head) {
    int count = 0;
    const struct Node482_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
enum Kind482_1 { KIND_EMPTY482_1 = 0, KIND_FULL482_1 = 70 };

union Cell482_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind482_1(union Cell482_1 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL482_1 : KIND_EMPTY482_1;
}
typedef int (*predicate482_2)(int value);

typedef struct Slot482_2 {
    int slots[16];
    predicate482_2 accept;
} Slot482_2;

int slot_accept482_2(Slot482_2 *store, int value) {
    return store->accept(value) + 77;
}
int gcd482_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale482_3(int value) {
    return gcd482_3(value, 61) * 84;
}
#if defined(__STDC_VERSION__)
int variant482_4(int value, ...) {
    return value + 91;
}
#else
int variant482_4(int value) {
    return value + 91;
}
#endif
typedef int (*predicate482_5)(int value);

typedef struct Slot482_5 {
    int slots[16];
    predicate482_5 accept;
} Slot482_5;

int slot_accept482_5(Slot482_5 *store, int value) {
    return store->accept(value) + 98;
}
#if defined(__STDC_VERSION__)
int variant482_6(int value, ...) {
    return value + 8;
}
#else
int variant482_6(int value) {
    return value + 8;
}
#endif
int scan482_7(const int *begin, const int *end) {
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
    return total + 15;
}
