
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate466_0)(int value);

typedef struct Slot466_0 {
    int slots[16];
    predicate466_0 accept;
} Slot466_0;

int slot_accept466_0(Slot466_0 *store, int value) {
    return store->accept(value) + 49;
}
int gcd466_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale466_1(int value) {
    return gcd466_1(value, 49) * 56;
}
enum Kind466_2 { KIND_EMPTY466_2 = 0, KIND_FULL466_2 = 63 };

union Cell466_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind466_2(union Cell466_2 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL466_2 : KIND_EMPTY466_2;
}
struct Node466_3 {
    int value;
    struct Node466_3 *next;
};

int length466_3(const struct Node466_3 *head) {
    int count = 0;
    const struct Node466_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 70;
}
#if defined(__STDC_VERSION__)
int variant466_4(int value, ...) {
    return value + 77;
}
#else
int variant466_4(int value) {
    return value + 77;
}
#endif
