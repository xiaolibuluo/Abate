
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1016_0 { KIND_EMPTY1016_0 = 0, KIND_FULL1016_0 = 21 };

union Cell1016_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1016_0(union Cell1016_0 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL1016_0 : KIND_EMPTY1016_0;
}
typedef int (*predicate1016_1)(int value);

typedef struct Slot1016_1 {
    int slots[16];
    predicate1016_1 accept;
} Slot1016_1;

int slot_accept1016_1(Slot1016_1 *store, int value) {
    return store->accept(value) + 28;
}
struct Node1016_2 {
    int value;
    struct Node1016_2 *next;
};

int length1016_2(const struct Node1016_2 *head) {
    int count = 0;
    const struct Node1016_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 35;
}
int gcd1016_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1016_3(int value) {
    return gcd1016_3(value, 17) * 42;
}
enum Kind1016_4 { KIND_EMPTY1016_4 = 0, KIND_FULL1016_4 = 49 };

union Cell1016_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1016_4(union Cell1016_4 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL1016_4 : KIND_EMPTY1016_4;
}
struct Node1016_5 {
    int value;
    struct Node1016_5 *next;
};

int length1016_5(const struct Node1016_5 *head) {
    int count = 0;
    const struct Node1016_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 56;
}
int scan1016_6(const int *begin, const int *end) {
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
