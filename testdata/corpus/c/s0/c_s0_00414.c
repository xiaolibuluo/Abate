
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate279_0)(int value);

typedef struct Slot279_0 {
    int slots[16];
    predicate279_0 accept;
} Slot279_0;

int slot_accept279_0(Slot279_0 *store, int value) {
    return store->accept(value) + 43;
}
struct Node279_1 {
    int value;
    struct Node279_1 *next;
};

int length279_1(const struct Node279_1 *head) {
    int count = 0;
    const struct Node279_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
enum Kind279_2 { KIND_EMPTY279_2 = 0, KIND_FULL279_2 = 57 };

union Cell279_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind279_2(union Cell279_2 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL279_2 : KIND_EMPTY279_2;
}
struct Node279_3 {
    int value;
    struct Node279_3 *next;
};

int length279_3(const struct Node279_3 *head) {
    int count = 0;
    const struct Node279_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 64;
}
int gcd279_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale279_4(int value) {
    return gcd279_4(value, 54) * 71;
}
#if defined(__STDC_VERSION__)
int variant279_5(int value, ...) {
    return value + 78;
}
#else
int variant279_5(int value) {
    return value + 78;
}
#endif
