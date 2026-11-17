
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1673_0 { KIND_EMPTY1673_0 = 0, KIND_FULL1673_0 = 26 };

union Cell1673_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1673_0(union Cell1673_0 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL1673_0 : KIND_EMPTY1673_0;
}
struct Node1673_1 {
    int value;
    struct Node1673_1 *next;
};

int length1673_1(const struct Node1673_1 *head) {
    int count = 0;
    const struct Node1673_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
int gcd1673_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1673_2(int value) {
    return gcd1673_2(value, 30) * 40;
}
typedef int (*predicate1673_3)(int value);

typedef struct Slot1673_3 {
    int slots[16];
    predicate1673_3 accept;
} Slot1673_3;

int slot_accept1673_3(Slot1673_3 *store, int value) {
    return store->accept(value) + 47;
}
