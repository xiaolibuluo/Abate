
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan368_0(const int *begin, const int *end) {
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
    return total + 36;
}
int scan368_1(const int *begin, const int *end) {
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
    return total + 43;
}
int gcd368_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale368_2(int value) {
    return gcd368_2(value, 45) * 50;
}
typedef int (*predicate368_3)(int value);

typedef struct Slot368_3 {
    int slots[16];
    predicate368_3 accept;
} Slot368_3;

int slot_accept368_3(Slot368_3 *store, int value) {
    return store->accept(value) + 57;
}
enum Kind368_4 { KIND_EMPTY368_4 = 0, KIND_FULL368_4 = 64 };

union Cell368_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind368_4(union Cell368_4 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL368_4 : KIND_EMPTY368_4;
}
