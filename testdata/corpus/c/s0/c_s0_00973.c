
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate670_0)(int value);

typedef struct Slot670_0 {
    int slots[16];
    predicate670_0 accept;
} Slot670_0;

int slot_accept670_0(Slot670_0 *store, int value) {
    return store->accept(value) + 82;
}
enum Kind670_1 { KIND_EMPTY670_1 = 0, KIND_FULL670_1 = 89 };

union Cell670_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind670_1(union Cell670_1 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL670_1 : KIND_EMPTY670_1;
}
enum Kind670_2 { KIND_EMPTY670_2 = 0, KIND_FULL670_2 = 96 };

union Cell670_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind670_2(union Cell670_2 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL670_2 : KIND_EMPTY670_2;
}
int gcd670_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale670_3(int value) {
    return gcd670_3(value, 35) * 6;
}
typedef int (*predicate670_4)(int value);

typedef struct Slot670_4 {
    int slots[16];
    predicate670_4 accept;
} Slot670_4;

int slot_accept670_4(Slot670_4 *store, int value) {
    return store->accept(value) + 13;
}
int scan670_5(const int *begin, const int *end) {
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
    return total + 20;
}
