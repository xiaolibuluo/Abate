
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1633_0 { KIND_EMPTY1633_0 = 0, KIND_FULL1633_0 = 88 };

union Cell1633_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1633_0(union Cell1633_0 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL1633_0 : KIND_EMPTY1633_0;
}
typedef int (*predicate1633_1)(int value);

typedef struct Slot1633_1 {
    int slots[16];
    predicate1633_1 accept;
} Slot1633_1;

int slot_accept1633_1(Slot1633_1 *store, int value) {
    return store->accept(value) + 95;
}
struct Node1633_2 {
    int value;
    struct Node1633_2 *next;
};

int length1633_2(const struct Node1633_2 *head) {
    int count = 0;
    const struct Node1633_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
int gcd1633_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1633_3(int value) {
    return gcd1633_3(value, 66) * 12;
}
int gcd1633_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1633_4(int value) {
    return gcd1633_4(value, 16) * 19;
}
