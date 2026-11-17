
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1060_0 { KIND_EMPTY1060_0 = 0, KIND_FULL1060_0 = 11 };

union Cell1060_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1060_0(union Cell1060_0 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL1060_0 : KIND_EMPTY1060_0;
}
struct Node1060_1 {
    int value;
    struct Node1060_1 *next;
};

int length1060_1(const struct Node1060_1 *head) {
    int count = 0;
    const struct Node1060_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
int scan1060_2(const int *begin, const int *end) {
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
    return total + 25;
}
typedef int (*predicate1060_3)(int value);

typedef struct Slot1060_3 {
    int slots[16];
    predicate1060_3 accept;
} Slot1060_3;

int slot_accept1060_3(Slot1060_3 *store, int value) {
    return store->accept(value) + 32;
}
typedef int (*predicate1060_4)(int value);

typedef struct Slot1060_4 {
    int slots[16];
    predicate1060_4 accept;
} Slot1060_4;

int slot_accept1060_4(Slot1060_4 *store, int value) {
    return store->accept(value) + 39;
}
int gcd1060_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1060_5(int value) {
    return gcd1060_5(value, 42) * 46;
}
