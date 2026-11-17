
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1301_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1301_0(int value) {
    return gcd1301_0(value, 27) * 40;
}
typedef int (*predicate1301_1)(int value);

typedef struct Slot1301_1 {
    int slots[16];
    predicate1301_1 accept;
} Slot1301_1;

int slot_accept1301_1(Slot1301_1 *store, int value) {
    return store->accept(value) + 47;
}
int gcd1301_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1301_2(int value) {
    return gcd1301_2(value, 49) * 54;
}
enum Kind1301_3 { KIND_EMPTY1301_3 = 0, KIND_FULL1301_3 = 61 };

union Cell1301_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1301_3(union Cell1301_3 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1301_3 : KIND_EMPTY1301_3;
}
struct Node1301_4 {
    int value;
    struct Node1301_4 *next;
};

int length1301_4(const struct Node1301_4 *head) {
    int count = 0;
    const struct Node1301_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
#if defined(__STDC_VERSION__)
int variant1301_5(int value, ...) {
    return value + 75;
}
#else
int variant1301_5(int value) {
    return value + 75;
}
#endif
int gcd1301_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1301_6(int value) {
    return gcd1301_6(value, 32) * 82;
}
