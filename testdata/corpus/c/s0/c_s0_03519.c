
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2374_0 {
    int value;
    struct Node2374_0 *next;
};

int length2374_0(const struct Node2374_0 *head) {
    int count = 0;
    const struct Node2374_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
int gcd2374_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2374_1(int value) {
    return gcd2374_1(value, 46) * 28;
}
enum Kind2374_2 { KIND_EMPTY2374_2 = 0, KIND_FULL2374_2 = 35 };

union Cell2374_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2374_2(union Cell2374_2 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL2374_2 : KIND_EMPTY2374_2;
}
typedef int (*predicate2374_3)(int value);

typedef struct Slot2374_3 {
    int slots[16];
    predicate2374_3 accept;
} Slot2374_3;

int slot_accept2374_3(Slot2374_3 *store, int value) {
    return store->accept(value) + 42;
}
typedef int (*predicate2374_4)(int value);

typedef struct Slot2374_4 {
    int slots[16];
    predicate2374_4 accept;
} Slot2374_4;

int slot_accept2374_4(Slot2374_4 *store, int value) {
    return store->accept(value) + 49;
}
int gcd2374_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2374_5(int value) {
    return gcd2374_5(value, 29) * 56;
}
int scan2374_6(const int *begin, const int *end) {
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
    return total + 63;
}
