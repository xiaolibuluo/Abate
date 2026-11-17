
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1885_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1885_0(int value) {
    return gcd1885_0(value, 28) * 66;
}
typedef int (*predicate1885_1)(int value);

typedef struct Slot1885_1 {
    int slots[16];
    predicate1885_1 accept;
} Slot1885_1;

int slot_accept1885_1(Slot1885_1 *store, int value) {
    return store->accept(value) + 73;
}
int scan1885_2(const int *begin, const int *end) {
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
    return total + 80;
}
struct Node1885_3 {
    int value;
    struct Node1885_3 *next;
};

int length1885_3(const struct Node1885_3 *head) {
    int count = 0;
    const struct Node1885_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
typedef int (*predicate1885_4)(int value);

typedef struct Slot1885_4 {
    int slots[16];
    predicate1885_4 accept;
} Slot1885_4;

int slot_accept1885_4(Slot1885_4 *store, int value) {
    return store->accept(value) + 94;
}
enum Kind1885_5 { KIND_EMPTY1885_5 = 0, KIND_FULL1885_5 = 101 };

union Cell1885_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1885_5(union Cell1885_5 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1885_5 : KIND_EMPTY1885_5;
}
