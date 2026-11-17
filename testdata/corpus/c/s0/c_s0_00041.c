
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind26_0 { KIND_EMPTY26_0 = 0, KIND_FULL26_0 = 52 };

union Cell26_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind26_0(union Cell26_0 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL26_0 : KIND_EMPTY26_0;
}
typedef int (*predicate26_1)(int value);

typedef struct Slot26_1 {
    int slots[16];
    predicate26_1 accept;
} Slot26_1;

int slot_accept26_1(Slot26_1 *store, int value) {
    return store->accept(value) + 59;
}
int gcd26_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale26_2(int value) {
    return gcd26_2(value, 30) * 66;
}
int gcd26_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale26_3(int value) {
    return gcd26_3(value, 41) * 73;
}
int scan26_4(const int *begin, const int *end) {
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
