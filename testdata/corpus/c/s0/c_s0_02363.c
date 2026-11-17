
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1594_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1594_0(int value) {
    return gcd1594_0(value, 65) * 66;
}
enum Kind1594_1 { KIND_EMPTY1594_1 = 0, KIND_FULL1594_1 = 73 };

union Cell1594_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1594_1(union Cell1594_1 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL1594_1 : KIND_EMPTY1594_1;
}
#if defined(__STDC_VERSION__)
int variant1594_2(int value, ...) {
    return value + 80;
}
#else
int variant1594_2(int value) {
    return value + 80;
}
#endif
struct Node1594_3 {
    int value;
    struct Node1594_3 *next;
};

int length1594_3(const struct Node1594_3 *head) {
    int count = 0;
    const struct Node1594_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
typedef int (*predicate1594_4)(int value);

typedef struct Slot1594_4 {
    int slots[16];
    predicate1594_4 accept;
} Slot1594_4;

int slot_accept1594_4(Slot1594_4 *store, int value) {
    return store->accept(value) + 94;
}
enum Kind1594_5 { KIND_EMPTY1594_5 = 0, KIND_FULL1594_5 = 101 };

union Cell1594_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1594_5(union Cell1594_5 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1594_5 : KIND_EMPTY1594_5;
}
struct Node1594_6 {
    int value;
    struct Node1594_6 *next;
};

int length1594_6(const struct Node1594_6 *head) {
    int count = 0;
    const struct Node1594_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
