
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1285_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1285_0(int value) {
    return gcd1285_0(value, 37) * 26;
}
typedef int (*predicate1285_1)(int value);

typedef struct Slot1285_1 {
    int slots[16];
    predicate1285_1 accept;
} Slot1285_1;

int slot_accept1285_1(Slot1285_1 *store, int value) {
    return store->accept(value) + 33;
}
struct Node1285_2 {
    int value;
    struct Node1285_2 *next;
};

int length1285_2(const struct Node1285_2 *head) {
    int count = 0;
    const struct Node1285_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
enum Kind1285_3 { KIND_EMPTY1285_3 = 0, KIND_FULL1285_3 = 47 };

union Cell1285_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1285_3(union Cell1285_3 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL1285_3 : KIND_EMPTY1285_3;
}
