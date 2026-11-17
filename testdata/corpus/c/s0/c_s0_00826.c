
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind569_0 { KIND_EMPTY569_0 = 0, KIND_FULL569_0 = 30 };

union Cell569_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind569_0(union Cell569_0 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL569_0 : KIND_EMPTY569_0;
}
int scan569_1(const int *begin, const int *end) {
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
    return total + 37;
}
int gcd569_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale569_2(int value) {
    return gcd569_2(value, 49) * 44;
}
enum Kind569_3 { KIND_EMPTY569_3 = 0, KIND_FULL569_3 = 51 };

union Cell569_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind569_3(union Cell569_3 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL569_3 : KIND_EMPTY569_3;
}
typedef int (*predicate569_4)(int value);

typedef struct Slot569_4 {
    int slots[16];
    predicate569_4 accept;
} Slot569_4;

int slot_accept569_4(Slot569_4 *store, int value) {
    return store->accept(value) + 58;
}
#if defined(__STDC_VERSION__)
int variant569_5(int value, ...) {
    return value + 65;
}
#else
int variant569_5(int value) {
    return value + 65;
}
#endif
